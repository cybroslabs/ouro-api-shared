package zstd

import (
	"encoding/binary"
	"fmt"
	"io"
	"sync"

	"github.com/klauspost/compress/zstd"
	"google.golang.org/grpc/encoding"
)

// Name is the name registered for the zstd compressor.
const Name = "zstd"

var encoderOptions = []zstd.EOption{
	// The default zstd window size is 8MB, which is much larger than the
	// typical RPC message and wastes a bunch of memory.
	zstd.WithWindowSize(512 * 1024),
}

func init() {
	c := &compressor{}
	c.poolCompressor.New = func() any {
		w, err := zstd.NewWriter(io.Discard, encoderOptions...)
		if err != nil {
			panic(err)
		}
		return &writer{Encoder: w, pool: &c.poolCompressor}
	}
	encoding.RegisterCompressor(c)
}

type writer struct {
	*zstd.Encoder
	pool *sync.Pool
}

// SetLevel updates the registered zstd compressor to use the compression level specified (zstd.HuffmanOnly is not supported).
// NOTE: this function must only be called during initialization time (i.e. in an init() function),
// and is not thread-safe.
//
// The error returned will be nil if the specified level is valid.
func SetLevel(level zstd.EncoderLevel) error {
	if level < zstd.SpeedFastest || level > zstd.SpeedBestCompression {
		return fmt.Errorf("grpc: invalid zstd compression level: %d", level)
	}
	c := encoding.GetCompressor(Name).(*compressor)
	c.poolCompressor.New = func() any {
		w, err := zstd.NewWriter(io.Discard, append(encoderOptions, zstd.WithEncoderLevel(level))...)
		if err != nil {
			panic(err)
		}
		return &writer{Encoder: w, pool: &c.poolCompressor}
	}
	return nil
}

func (c *compressor) Compress(w io.Writer) (io.WriteCloser, error) {
	z := c.poolCompressor.Get().(*writer)
	z.Reset(w)
	return z, nil
}

func (z *writer) Close() error {
	defer z.pool.Put(z)
	return z.Encoder.Close()
}

type reader struct {
	*zstd.Decoder
	pool *sync.Pool
}

func (c *compressor) Decompress(r io.Reader) (io.Reader, error) {
	z, inPool := c.poolDecompressor.Get().(*reader)
	if !inPool {
		newZ, err := zstd.NewReader(r)
		if err != nil {
			return nil, err
		}
		return &reader{Decoder: newZ, pool: &c.poolDecompressor}, nil
	}
	if err := z.Reset(r); err != nil {
		c.poolDecompressor.Put(z)
		return nil, err
	}
	return z, nil
}

func (z *reader) Read(p []byte) (n int, err error) {
	n, err = z.Decoder.Read(p)
	if err == io.EOF {
		z.pool.Put(z)
	}
	return n, err
}

// RFC1952 specifies that the last four bytes "contains the size of
// the original (uncompressed) input data modulo 2^32."
// gRPC has a max message size of 2GB so we don't need to worry about wraparound.
func (c *compressor) DecompressedSize(buf []byte) int {
	last := len(buf)
	if last < 4 {
		return -1
	}
	return int(binary.LittleEndian.Uint32(buf[last-4 : last]))
}

func (c *compressor) Name() string {
	return Name
}

type compressor struct {
	poolCompressor   sync.Pool
	poolDecompressor sync.Pool
}
