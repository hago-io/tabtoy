package tabtoy

import "encoding/binary"

type binaryReader struct {
	buf []byte
}

func (reader *binaryReader) ReadBool(x *bool) {
	n := reader.buf[0]
	reader.buf = reader.buf[1:]
	*x = n != 0
}

func (reader *binaryReader) ReadUInt16(x *uint16) {
	*x = binary.LittleEndian.Uint16(reader.buf[0:2])
	reader.buf = reader.buf[2:]
}

func (reader *binaryReader) ReadUInt32(x *uint32) {
	*x = binary.LittleEndian.Uint32(reader.buf[0:4])
	reader.buf = reader.buf[4:]
}
func (reader *binaryReader) ReadUInt64(x *uint64) {
	*x = binary.LittleEndian.Uint64(reader.buf[0:8])
	reader.buf = reader.buf[8:]
}

func (reader *binaryReader) ReadBytes(x *[]byte) {
	var l uint16
	reader.ReadUInt16(&l)
	*x = make([]byte, l)

	copy(*x, reader.buf[0:l])
	reader.buf = reader.buf[l:]
}

func (reader *binaryReader) ReadInt16(x *int16) {

	var v uint16
	reader.ReadUInt16(&v)

	*x = int16(v)
}

func (reader *binaryReader) ReadInt32(x *int32) {
	var v uint32
	reader.ReadUInt32(&v)

	*x = int32(v)
}

func (reader *binaryReader) ReadInt64(x *int64) {
	var v uint64
	reader.ReadUInt64(&v)

	*x = int64(v)
}
