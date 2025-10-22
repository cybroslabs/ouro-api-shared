package acquisition

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"iter"
	"math"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/utils/ptr"
)

var (
	_version = byte(1)
)

/*

1B - version (maybe magic number here?)
4B - period in seconds
any bytes - unit string (4B length + string)
<for each block>
	8B - start timestamp (unix)
	4B - bytes count
	1B - items count, 0 means 1 and so on, no empty block possible then
	1B - value type (1 double 8B, 2 int64 8B, 3 string xB, 4 timestamp 8B, 5 timestamp with timezone xB, 6 boolean 1B)
	<for each item>
		1B - item header, bitfield, status/nstatus present bit, exponent present bit, peak present bit, not used bit, higher 4 bites, value subtype, bit 4 usually marks previous value
		optional 8B - status
		optional 8B - nstatus
		optional 4B - exponent
		optional 8B - peak ts (unix)
		any bytes - value, depending on type, numbers, bool and timestamps have direct representation, string has present flag and length

value subtypes:
	double - no subtype (no previous value possible)
	integer - 0 - 0, 1 - prev value, 2 - single byte integer, 3 - double byte integer, 4 - four byte integer, 5 - eight byte integer
	string - 0 - empty string, 1 - previous, 2 - single byte length, 3 - double byte length, 4 - four byte length, 5-15 - length in that value - 4
	timestamp - no subtype (no previous value possible)
	timestamp with timezone - same as string
	boolean - 0 - false, 1 - true (no previous value possible, really not needed)
*/

const (
	typeUnspecified     = byte(0)
	typeDouble          = byte(1)
	typeInteger         = byte(2)
	typeString          = byte(3)
	typeTimestamp       = byte(4)
	typeTimestampWithTz = byte(5)
	typeBoolean         = byte(6)
)

type TimeWithTimeZone struct {
	Timestamp time.Time
}

type maxvalueHeader [29]byte

type ProfileValuesEncoder struct { // statefull, so no thread safe
	nextstring    string
	nextinteger   int64
	nexttype      byte
	nextstatus    int64
	nextnstatus   uint64
	nextexponent  int32
	nextstamp     time.Time
	unit          string
	periodseconds int32

	lastblockoffset int
	items           int
	buffer          bytes.Buffer
}

type ProfileValuesDecoder struct { // statefull, so no thread safe
	empty         bool
	periodseconds int32
	unit          string

	buffer       []byte
	headerlength int
}

type ProfileValueItem struct {
	Timestamp time.Time
	Value     *MeasuredValue
	Err       error
}

func NewProfileValuesEncoder(periodseconds int32, unit string) *ProfileValuesEncoder { // no value at all means no bytes, is it ok?
	ret := &ProfileValuesEncoder{
		unit:          unit,
		periodseconds: periodseconds,
	}
	ret.writeheader()
	return ret
}

func (pe *ProfileValuesEncoder) Reset(periodseconds int32, unit string) {
	pe.nextstring = ""
	pe.nextinteger = 0
	pe.nexttype = typeUnspecified
	pe.nextstatus = 0
	pe.nextnstatus = 0
	pe.nextexponent = 0
	pe.nextstamp = time.Time{}
	pe.unit = unit
	pe.periodseconds = periodseconds
	pe.items = 0
	pe.buffer.Reset()

	pe.writeheader()
}

func (pe *ProfileValuesEncoder) writeheader() {
	var tmp [4]byte

	_ = pe.buffer.WriteByte(_version)
	binary.BigEndian.PutUint32(tmp[:], uint32(pe.periodseconds))
	_, _ = pe.buffer.Write(tmp[:])
	// no bits, so just write four bytes length and unit itself, not ideal, but it is only single unit
	ln := len(pe.unit)
	binary.BigEndian.PutUint32(tmp[:], uint32(ln))
	_, _ = pe.buffer.Write(tmp[:])
	_, _ = pe.buffer.WriteString(pe.unit)
}

func (pe *ProfileValuesEncoder) Bytes() []byte { // should be called at the end, closing also current block, append can continue but it wont be so effective (new block, wasting bytes)
	pe.closeblock()
	return pe.buffer.Bytes()
}

func (pe *ProfileValuesEncoder) closeblock() {
	if pe.nexttype == typeUnspecified { // first or closed block, do nothing
		return
	}
	b := pe.buffer.Bytes() // this is a bit hardcore to update that buffer, but whatever ;)
	to := pe.lastblockoffset
	binary.BigEndian.PutUint32(b[to+8:], uint32(pe.buffer.Len()-to))
	b[to+12] = byte(pe.items - 1)
}

func (pe *ProfileValuesEncoder) startblock(ts time.Time, valueType byte) {
	var tmp [14]byte // header
	binary.BigEndian.PutUint64(tmp[:], uint64(ts.Unix()))
	tmp[13] = valueType
	pe.lastblockoffset = pe.buffer.Len()
	pe.items = 0
	pe.nexttype = valueType
	pe.nextstatus = 0
	pe.nextnstatus = 0
	pe.nextexponent = 0
	pe.nextstring = ""
	pe.nextinteger = 0
	_, _ = pe.buffer.Write(tmp[:])
}

func (pe *ProfileValuesEncoder) codevalueheader(dst *maxvalueHeader, status int64, nstatus uint64, exponent int32, pts *time.Time) int {
	off := 1
	dst[0] = 0

	if status != pe.nextstatus || nstatus != pe.nextnstatus {
		binary.BigEndian.PutUint64(dst[off:], uint64(status))
		binary.BigEndian.PutUint64(dst[off+8:], uint64(nstatus))
		pe.nextstatus = status
		pe.nextnstatus = nstatus
		dst[0] |= 1
		off += 16
	}
	if exponent != pe.nextexponent {
		binary.BigEndian.PutUint32(dst[off:], uint32(exponent))
		pe.nextexponent = exponent
		dst[0] |= 2
		off += 4
	}
	if pts != nil {
		binary.BigEndian.PutUint64(dst[off:], uint64(pts.Unix()))
		dst[0] |= 4
		off += 8
	}

	return off
}

func (pe *ProfileValuesEncoder) nextblock(ts time.Time, t byte) {
	if pe.items == 256 || pe.nexttype != t || !pe.nextstamp.Equal(ts) { // create a new block
		pe.closeblock()
		pe.startblock(ts, t)
	}
}

func (pe *ProfileValuesEncoder) AppendValue(ts time.Time, status int64, nstatus uint64, exponent int32, value any, pts *time.Time) error {
	switch v := value.(type) { // maybe more types here?
	case float64:
		pe.AppendDouble(ts, status, nstatus, exponent, v, pts)
	case int64:
		pe.AppendInteger(ts, status, nstatus, exponent, v, pts)
	case string:
		pe.AppendString(ts, status, nstatus, exponent, v, pts)
	case time.Time:
		pe.AppendTimestamp(ts, status, nstatus, exponent, v, pts)
	case TimeWithTimeZone:
		pe.AppendTimestampWithTz(ts, status, nstatus, exponent, v.Timestamp, pts)
	case bool:
		pe.AppendBoolean(ts, status, nstatus, exponent, v, pts)
	default:
		return fmt.Errorf("unknown type to add %T", v)
	}
	return nil
}

func (pe *ProfileValuesEncoder) AppendInteger(ts time.Time, status int64, nstatus uint64, exponent int32, value int64, pts *time.Time) {
	var tmp maxvalueHeader
	pe.nextblock(ts, typeInteger)

	pe.items++
	pe.nextstamp = ts.Add(time.Duration(pe.periodseconds) * time.Second)
	off := pe.codevalueheader(&tmp, status, nstatus, exponent, pts)
	if pe.nextinteger == value {
		tmp[0] |= 0x10
		_, _ = pe.buffer.Write(tmp[:off])
		return
	}

	pe.nextinteger = value
	if value == 0 { // special case
		_, _ = pe.buffer.Write(tmp[:off])
		return
	}
	if value == int64(int8(value&0xff)) {
		tmp[0] |= 0x20
		_, _ = pe.buffer.Write(tmp[:off])
		_ = pe.buffer.WriteByte(byte(value))
		return
	}
	if value == int64(int16(value&0xffff)) {
		tmp[0] |= 0x30
		_, _ = pe.buffer.Write(tmp[:off])
		binary.BigEndian.PutUint16(tmp[:], uint16(value))
		_, _ = pe.buffer.Write(tmp[:2])
		return
	}
	if value == int64(int32(value&0xffffffff)) {
		tmp[0] |= 0x40
		_, _ = pe.buffer.Write(tmp[:off])
		binary.BigEndian.PutUint32(tmp[:], uint32(value))
		_, _ = pe.buffer.Write(tmp[:4])
		return
	}

	tmp[0] |= 0x50
	_, _ = pe.buffer.Write(tmp[:off])
	binary.BigEndian.PutUint64(tmp[:], uint64(value))
	_, _ = pe.buffer.Write(tmp[:8])
}

func (pe *ProfileValuesEncoder) AppendDouble(ts time.Time, status int64, nstatus uint64, exponent int32, value float64, pts *time.Time) {
	var tmp maxvalueHeader
	pe.nextblock(ts, typeDouble)

	pe.items++
	pe.nextstamp = ts.Add(time.Duration(pe.periodseconds) * time.Second)
	off := pe.codevalueheader(&tmp, status, nstatus, exponent, pts)
	_, _ = pe.buffer.Write(tmp[:off])
	binary.BigEndian.PutUint64(tmp[:], math.Float64bits(value))
	_, _ = pe.buffer.Write(tmp[:8]) // write that everytime, no subtype
}

func (pe *ProfileValuesEncoder) AppendString(ts time.Time, status int64, nstatus uint64, exponent int32, value string, pts *time.Time) {
	pe.nextblock(ts, typeString)
	pe.appendString(ts, status, nstatus, exponent, value, pts)
}

func (pe *ProfileValuesEncoder) appendString(ts time.Time, status int64, nstatus uint64, exponent int32, value string, pts *time.Time) {
	var tmp maxvalueHeader

	pe.items++
	pe.nextstamp = ts.Add(time.Duration(pe.periodseconds) * time.Second)
	off := pe.codevalueheader(&tmp, status, nstatus, exponent, pts)
	if pe.nextstring == value {
		tmp[0] |= 0x10
		_, _ = pe.buffer.Write(tmp[:off])
		return
	}

	pe.nextstring = value
	ln := uint32(len(value))
	if ln == 0 {
		_, _ = pe.buffer.Write(tmp[:off])
		return
	}
	if ln < 12 {
		tmp[0] |= (byte(ln) + 4) << 4
		_, _ = pe.buffer.Write(tmp[:off])
		_, _ = pe.buffer.WriteString(value)
		return
	}
	if ln < 256 {
		tmp[0] |= 0x20
		_, _ = pe.buffer.Write(tmp[:off])
		_ = pe.buffer.WriteByte(byte(ln))
		_, _ = pe.buffer.WriteString(value)
		return
	}
	if ln < 65536 {
		tmp[0] |= 0x30
		_, _ = pe.buffer.Write(tmp[:off])
		binary.BigEndian.PutUint16(tmp[:], uint16(ln))
		_, _ = pe.buffer.Write(tmp[:2])
		_, _ = pe.buffer.WriteString(value)
		return
	}

	tmp[0] |= 0x40
	_, _ = pe.buffer.Write(tmp[:off])
	binary.BigEndian.PutUint32(tmp[:], uint32(ln))
	_, _ = pe.buffer.Write(tmp[:4])
	_, _ = pe.buffer.WriteString(value)
}

func (pe *ProfileValuesEncoder) AppendTimestamp(ts time.Time, status int64, nstatus uint64, exponent int32, value time.Time, pts *time.Time) {
	var tmp maxvalueHeader
	pe.nextblock(ts, typeTimestamp)

	pe.items++
	pe.nextstamp = ts.Add(time.Duration(pe.periodseconds) * time.Second)
	off := pe.codevalueheader(&tmp, status, nstatus, exponent, pts)
	_, _ = pe.buffer.Write(tmp[:off])
	binary.BigEndian.PutUint64(tmp[:], uint64(value.Unix()))
	_, _ = pe.buffer.Write(tmp[:8]) // no subtype
}

func (pe *ProfileValuesEncoder) AppendTimestampWithTz(ts time.Time, status int64, nstatus uint64, exponent int32, value time.Time, pts *time.Time) {
	pe.nextblock(ts, typeTimestampWithTz)
	str := value.Format(time.RFC3339)
	pe.appendString(ts, status, nstatus, exponent, str, pts)
}

func (pe *ProfileValuesEncoder) AppendBoolean(ts time.Time, status int64, nstatus uint64, exponent int32, value bool, pts *time.Time) {
	var tmp maxvalueHeader
	pe.nextblock(ts, typeBoolean)

	pe.items++
	pe.nextstamp = ts.Add(time.Duration(pe.periodseconds) * time.Second)
	off := pe.codevalueheader(&tmp, status, nstatus, exponent, pts)
	if value {
		tmp[0] |= 0x10 // use subtype
	}
	_, _ = pe.buffer.Write(tmp[:off])
}

func MergeProfileBlobs(dst []byte, src1, src2 []byte) ([]byte, error) {
	dst = dst[:0]
	pd1, err := NewProfileValuesDecoder(src1)
	if err != nil {
		return dst, err
	}
	if pd1.IsNil() {
		dst = append(dst, src2...)
		return dst, nil
	}
	pd2, err := NewProfileValuesDecoder(src2)
	if err != nil {
		return dst, err
	}
	if pd2.IsNil() {
		dst = append(dst, src1...)
		return dst, nil
	}
	if pd1.GetPeriod() != pd2.GetPeriod() {
		return dst, fmt.Errorf("cannot merge blobs with different period")
	}
	if pd1.GetUnit() != pd2.GetUnit() {
		return dst, fmt.Errorf("cannot merge blobs with different unit")
	}
	dst = append(dst, src1...)
	dst = append(dst, src2[pd2.headerlength:]...) // skip header of second blob
	return dst, nil
}

func NewProfileValuesDecoder(src []byte) (*ProfileValuesDecoder, error) { // creates no copy from src, so watch out
	var tmp [9]byte // version, period, unit string length
	ret := &ProfileValuesDecoder{}
	if len(src) == 0 {
		ret.empty = true
		return ret, nil
	}
	ret.buffer = src
	b := bytes.NewReader(src)
	_, err := io.ReadFull(b, tmp[:])
	if err != nil {
		return nil, err
	}
	if tmp[0] != _version {
		return nil, fmt.Errorf("invalid data version: %d", tmp[0])
	}
	ret.periodseconds = int32(binary.BigEndian.Uint32(tmp[1:5]))
	ln := int(binary.BigEndian.Uint32(tmp[5:9]))
	if ln > 0 {
		unitb := make([]byte, ln)
		_, err = io.ReadFull(b, unitb)
		if err != nil {
			return nil, err
		}
		ret.unit = string(unitb)
	}
	ret.headerlength = 9 + ln
	return ret, nil
}

func (pd *ProfileValuesDecoder) IsNil() bool {
	return pd.empty
}

func (pd *ProfileValuesDecoder) GetPeriod() (d time.Duration) {
	if pd.empty {
		return
	}
	d = time.Duration(pd.periodseconds) * time.Second
	return
}

func (pd *ProfileValuesDecoder) GetPeriodSeconds() (d int32) {
	if pd.empty {
		return
	}
	d = pd.periodseconds
	return
}

func (pd *ProfileValuesDecoder) GetUnit() (u string) {
	if pd.empty {
		return
	}
	u = pd.unit
	return
}

func (pd *ProfileValuesDecoder) GetLastTimeStamp() (ltt time.Time, err error) {
	var tmp [14]byte
	var lt time.Time

	if pd.empty {
		return
	}

	bf := bytes.NewReader(pd.buffer)
	_, err = bf.Seek(int64(pd.headerlength), io.SeekStart)
	if err != nil {
		return
	}

	for bf.Len() != 0 {
		_, err = io.ReadFull(bf, tmp[:])
		if err != nil {
			return
		}

		b := binary.BigEndian.Uint32(tmp[8:])
		if b < 15 {
			err = fmt.Errorf("data error, invalid block size in bytes")
			return
		}
		ts := time.Unix(int64(binary.BigEndian.Uint64(tmp[:]))+int64(pd.periodseconds)*int64(tmp[12]), 0)
		if lt.Before(ts) {
			lt = ts
		}
		_, err = bf.Seek(int64(b-14), io.SeekCurrent)
		if err != nil {
			return
		}
	}
	ltt = lt
	return
}

type decodeContext struct {
	prevstring   string
	previnteger  int64
	prevstatus   int64
	prevnstatus  uint64
	prevexponent int32
	haspeak      bool // no alloc peak
	peakts       time.Time
}

func (ctx *decodeContext) decodeItemHeader(bf io.Reader) (itemid byte, err error) {
	var tmp [16]byte

	_, err = io.ReadFull(bf, tmp[:1])
	if err != nil {
		return
	}
	itemid = tmp[0]
	if itemid&1 != 0 { // there are statuses
		_, err = io.ReadFull(bf, tmp[:])
		if err != nil {
			return
		}
		ctx.prevstatus = int64(binary.BigEndian.Uint64(tmp[:]))
		ctx.prevnstatus = binary.BigEndian.Uint64(tmp[8:])
	}
	if itemid&2 != 0 { // exponent is present
		_, err = io.ReadFull(bf, tmp[:4])
		if err != nil {
			return
		}
		ctx.prevexponent = int32(binary.BigEndian.Uint32(tmp[:]))
	}
	ctx.haspeak = false
	if itemid&4 != 0 { // peak timestamp
		_, err = io.ReadFull(bf, tmp[:8])
		if err != nil {
			return
		}
		ctx.peakts = time.Unix(int64(binary.BigEndian.Uint64(tmp[:])), 0)
		ctx.haspeak = true
	}
	return
}

func (ctx *decodeContext) clear() {
	ctx.prevstring = ""
	ctx.previnteger = 0
	ctx.prevstatus = 0
	ctx.prevnstatus = 0
	ctx.prevexponent = 0
	ctx.haspeak = false
	ctx.peakts = time.Time{}
}

func (pd *ProfileValuesDecoder) Values() iter.Seq[ProfileValueItem] {
	return func(yield func(ProfileValueItem) bool) {
		if pd.empty {
			return
		}

		bff := bytes.NewReader(pd.buffer)
		_, err := bff.Seek(int64(pd.headerlength), io.SeekStart)
		if err != nil {
			yield(ProfileValueItem{Err: err})
			return
		}

		var tmp [14]byte // at least block header
		var ctx decodeContext
		var itemid byte

		for bff.Len() != 0 {
			ctx.clear()
			_, err = io.ReadFull(bff, tmp[:])
			if err != nil {
				yield(ProfileValueItem{Err: err})
				return
			}

			cnt := int(tmp[12]) + 1
			ts := time.Unix(int64(binary.BigEndian.Uint64(tmp[:])), 0)
			b := binary.BigEndian.Uint32(tmp[8:])
			if b < 15 {
				err = fmt.Errorf("data error, invalid block size in bytes")
				yield(ProfileValueItem{Err: err})
				return
			}

			bf := io.LimitReader(bff, int64(b-14))
			switch tmp[13] {
			case typeUnspecified:
				err = fmt.Errorf("data error, unspecified type in block")
				yield(ProfileValueItem{Err: err})
				return
			case typeDouble:
				for range cnt {
					itemid, err = ctx.decodeItemHeader(bf)
					if err != nil {
						yield(ProfileValueItem{Err: err})
						return
					}
					if itemid&0xf8 != 0 { // really handle every fucking case?
						err = fmt.Errorf("data error, invalid item id byte %x", itemid)
						yield(ProfileValueItem{Err: err})
						return
					}

					val := ProfileValueItem{
						Timestamp: ts,
						Value: MeasuredValue_builder{
							Status:   ptr.To(ctx.prevstatus), // a bit weird, so value is valid only inside yield? that should be checked
							Nstatus:  ptr.To(ctx.prevnstatus),
							Exponent: ptr.To(ctx.prevexponent),
						}.Build(),
					}
					if ctx.haspeak {
						val.Value.SetPeakTs(timestamppb.New(ctx.peakts))
					}
					_, err = io.ReadFull(bf, tmp[:8])
					if err != nil {
						yield(ProfileValueItem{Err: err})
						return
					}
					val.Value.SetDoubleValue(math.Float64frombits(binary.BigEndian.Uint64(tmp[:])))

					if !yield(val) {
						return
					}
					ts = ts.Add(time.Duration(pd.periodseconds) * time.Second)
				}
			case typeInteger:
				for range cnt {
					itemid, err = ctx.decodeItemHeader(bf)
					if err != nil {
						yield(ProfileValueItem{Err: err})
						return
					}

					val := ProfileValueItem{
						Timestamp: ts,
						Value: MeasuredValue_builder{
							Status:   ptr.To(ctx.prevstatus), // a bit weird, so value is valid only inside yield? that should be checked
							Nstatus:  ptr.To(ctx.prevnstatus),
							Exponent: ptr.To(ctx.prevexponent),
						}.Build(),
					}
					if ctx.haspeak {
						val.Value.SetPeakTs(timestamppb.New(ctx.peakts))
					}

					switch itemid >> 4 {
					case 0:
						ctx.previnteger = 0
					case 1:
					case 2:
						_, err = io.ReadFull(bf, tmp[:1])
						ctx.previnteger = int64(int8(tmp[0]))
					case 3:
						_, err = io.ReadFull(bf, tmp[:2])
						ctx.previnteger = int64(int16(binary.BigEndian.Uint16(tmp[:2])))
					case 4:
						_, err = io.ReadFull(bf, tmp[:4])
						ctx.previnteger = int64(int32(binary.BigEndian.Uint32(tmp[:4])))
					case 5:
						_, err = io.ReadFull(bf, tmp[:8])
						ctx.previnteger = int64(binary.BigEndian.Uint64(tmp[:8]))
					default:
						err = fmt.Errorf("data error, invalid item id byte %x", itemid)
					}
					if err != nil {
						yield(ProfileValueItem{Err: err})
						return
					}
					val.Value.SetIntegerValue(ctx.previnteger)

					if !yield(val) {
						return
					}
					ts = ts.Add(time.Duration(pd.periodseconds) * time.Second)
				}
			case typeString:
				for range cnt {
					itemid, err = ctx.decodeItemHeader(bf)
					if err != nil {
						yield(ProfileValueItem{Err: err})
						return
					}

					val := ProfileValueItem{
						Timestamp: ts,
						Value: MeasuredValue_builder{
							Status:   ptr.To(ctx.prevstatus), // a bit weird, so value is valid only inside yield? that should be checked
							Nstatus:  ptr.To(ctx.prevnstatus),
							Exponent: ptr.To(ctx.prevexponent),
						}.Build(),
					}
					if ctx.haspeak {
						val.Value.SetPeakTs(timestamppb.New(ctx.peakts))
					}

					err = decodeString(itemid, &ctx, bf)
					if err != nil {
						yield(ProfileValueItem{Err: err})
						return
					}
					val.Value.SetStringValue(ctx.prevstring)

					if !yield(val) {
						return
					}
					ts = ts.Add(time.Duration(pd.periodseconds) * time.Second)
				}
			case typeTimestamp:
				for range cnt {
					_, err = ctx.decodeItemHeader(bf)
					if err != nil {
						yield(ProfileValueItem{Err: err})
						return
					}

					val := ProfileValueItem{
						Timestamp: ts,
						Value: MeasuredValue_builder{
							Status:   ptr.To(ctx.prevstatus), // a bit weird, so value is valid only inside yield? that should be checked
							Nstatus:  ptr.To(ctx.prevnstatus),
							Exponent: ptr.To(ctx.prevexponent),
						}.Build(),
					}
					if ctx.haspeak {
						val.Value.SetPeakTs(timestamppb.New(ctx.peakts))
					}
					_, err = io.ReadFull(bf, tmp[:8])
					if err != nil {
						yield(ProfileValueItem{Err: err})
						return
					}
					val.Value.SetTimestampValue(timestamppb.New(time.Unix(int64(binary.BigEndian.Uint64(tmp[:])), 0)))

					if !yield(val) {
						return
					}
					ts = ts.Add(time.Duration(pd.periodseconds) * time.Second)
				}
			case typeTimestampWithTz:
				for range cnt {
					itemid, err = ctx.decodeItemHeader(bf)
					if err != nil {
						yield(ProfileValueItem{Err: err})
						return
					}

					val := ProfileValueItem{
						Timestamp: ts,
						Value: MeasuredValue_builder{
							Status:   ptr.To(ctx.prevstatus), // a bit weird, so value is valid only inside yield? that should be checked
							Nstatus:  ptr.To(ctx.prevnstatus),
							Exponent: ptr.To(ctx.prevexponent),
						}.Build(),
					}
					if ctx.haspeak {
						val.Value.SetPeakTs(timestamppb.New(ctx.peakts))
					}

					err = decodeString(itemid, &ctx, bf)
					if err != nil {
						yield(ProfileValueItem{Err: err})
						return
					}
					val.Value.SetTimestampTzValue(ctx.prevstring)

					if !yield(val) {
						return
					}
					ts = ts.Add(time.Duration(pd.periodseconds) * time.Second)
				}
			case typeBoolean:
				for range cnt {
					itemid, err = ctx.decodeItemHeader(bf)
					if err != nil {
						yield(ProfileValueItem{Err: err})
						return
					}

					val := ProfileValueItem{
						Timestamp: ts,
						Value: MeasuredValue_builder{
							Status:   ptr.To(ctx.prevstatus), // a bit weird, so value is valid only inside yield? that should be checked
							Nstatus:  ptr.To(ctx.prevnstatus),
							Exponent: ptr.To(ctx.prevexponent),
						}.Build(),
					}
					if ctx.haspeak {
						val.Value.SetPeakTs(timestamppb.New(ctx.peakts))
					}
					if itemid&0x10 != 0 {
						val.Value.SetBoolValue(true)
					} else {
						val.Value.SetBoolValue(false)
					}

					if !yield(val) {
						return
					}
					ts = ts.Add(time.Duration(pd.periodseconds) * time.Second)
				}
			default:
				err = fmt.Errorf("data error, unknown type %d in block", tmp[13])
				yield(ProfileValueItem{Err: err})
				return
			}
		}
	}
}

func decodeString(itemid byte, ctx *decodeContext, bf io.Reader) (err error) {
	var tmp [4]byte
	id := itemid >> 4
	switch id {
	case 0:
		ctx.prevstring = ""
	case 1:
	case 2:
		_, err = io.ReadFull(bf, tmp[:1])
		if err != nil {
			return
		}
		str := make([]byte, tmp[0])
		_, err = io.ReadFull(bf, str)
		if err != nil {
			return
		}
		ctx.prevstring = string(str)
	case 3:
		_, err = io.ReadFull(bf, tmp[:2])
		if err != nil {
			return
		}
		str := make([]byte, binary.BigEndian.Uint16(tmp[:]))
		_, err = io.ReadFull(bf, str)
		if err != nil {
			return
		}
		ctx.prevstring = string(str)
	case 4:
		_, err = io.ReadFull(bf, tmp[:4])
		if err != nil {
			return
		}
		str := make([]byte, binary.BigEndian.Uint32(tmp[:]))
		_, err = io.ReadFull(bf, str)
		if err != nil {
			return
		}
		ctx.prevstring = string(str)
	default:
		id -= 4
		str := make([]byte, id)
		_, err = io.ReadFull(bf, str)
		if err != nil {
			return
		}
		ctx.prevstring = string(str)
	}
	return
}
