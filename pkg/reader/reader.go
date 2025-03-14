package reader

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Reader struct {
	f             *os.File
	Data          []byte
	SegmentBase   int
	SegmentOffset int
	Curr          byte

	byteRecord []byte
}

func New(filepath string) (*Reader, error) {
	r := Reader{}
	f, err := r.Read(filepath)
	if err != nil {
		return nil, err
	}

	r.f = f
	data, err := r.ReadFull(r.f)
	if err != nil {
		r.Close()
		return nil, err
	}

	r.Data = data
	return &r, nil
}

func (r *Reader) Close() {
	r.f.Close()
}

func (r *Reader) BeginByteRecord() {
	r.byteRecord = []byte{}
}

func (r *Reader) Read(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)

	return file, err
}

func (r *Reader) ReadNextBytes(file *os.File, number int) ([]byte, error) {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		return bytes, err
	}
	return bytes, nil
}

func (r *Reader) ReadFull(file *os.File) ([]byte, error) {
	data := []byte{}

	for {
		d, err := r.GetNextBytes(file, 1)
		if err != nil {
			break
		}

		data = append(data, d...)
	}

	r.Data = data

	return data, nil
}

func (r *Reader) GetNextBytes(f *os.File, n int) ([]byte, error) {
	data, err := r.ReadNextBytes(f, n)
	if err != nil {
		return []byte{}, fmt.Errorf("no more bytes")
	}

	code := make([]byte, n)
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.LittleEndian, code)
	if err != nil {
		return []byte{}, err
	}
	return code, nil
}

func (r *Reader) AccessData(offset int) (byte, error) {
	idx := r.SegmentOffset + offset
	if idx >= len(r.Data) {
		return 0, fmt.Errorf("index out of range %d", idx)
	}
	return r.Data[idx], nil
}

func (r *Reader) ReadByte() (byte, error) {
	if r.SegmentOffset >= len(r.Data) {
		r.Curr = 0
		return 0, fmt.Errorf("end of data")
	} else {
		r.Curr = r.Data[r.SegmentOffset]
	}
	r.SegmentBase = r.SegmentOffset
	r.SegmentOffset++

	r.byteRecord = append(r.byteRecord, r.Curr)
	return r.Curr, nil
}

func (r *Reader) Rewind(offset int) {
	r.SegmentOffset = offset
	r.Curr = r.Data[r.SegmentOffset]
	r.SegmentBase = r.SegmentOffset
}

func (r *Reader) Dump() string {
	var binout bytes.Buffer
	var out bytes.Buffer

	for _, b := range r.Data {
		binout.WriteString(fmt.Sprintf("%08b ", b))
		out.WriteString(fmt.Sprintf("0x%2x,", b))
	}
	return fmt.Sprintf("%s\n%s", binout.String(), out.String())
}

func (r *Reader) EndByteRecord() []byte {
	rec := r.byteRecord
	r.byteRecord = []byte{}
	return rec
}
