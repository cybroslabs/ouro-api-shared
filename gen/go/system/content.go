package system

import (
	embed "embed"
	"io"
	"io/fs"
)

//go:embed content/license.pem
var content embed.FS

var subfolder fs.FS

func init() {
	subfolder, _ = fs.Sub(content, "content")
}

// ReadFile reads and returns the content of the named file.
func ReadFile(path string) ([]byte, error) {
	f, err := subfolder.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	// Read the file content
	return io.ReadAll(f)
}
