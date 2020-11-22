package fileutil

import (
	"io"
	"log"
	"os"

	"github.com/ghtalpo/egb/common/mathutil"
)

// Reader ...
type Reader interface {
	ReadByteSafe() byte
	ReadWordSafe() int
	ReadDWordSafe() int
	ReadVecSafe(length int) []byte
	ReadStringSafe(length int) string
	SkipBytesSafe(length int)
}

// Writer ...
type Writer interface {
	WriteByteSafe(b byte)
	WriteWordSafe(n int)
	WriteDWordSafe(n int)
	WriteStringSafe(s string, length int)
	WriteVecSafe(v []byte, startp int, length int)
}

// IFS ...
type IFS struct {
	f *os.File
}

// NewIFS ... a constructor
func NewIFS(f *os.File) *IFS {
	return &IFS{f: f}
}

// Seek ...
func (i *IFS) Seek(n int) error {
	if _, err := i.f.Seek(int64(n), 0); err != nil {
		panic(err)
	}
	return nil
}

// GetLen ...
func (i *IFS) GetLen() int {
	fileInfo, err := i.f.Stat()
	if err != nil {
		panic(err)
	}
	return int(fileInfo.Size())
}

// GetBuffer ...
func (i *IFS) GetBuffer(n int) []byte {
	b := make([]byte, n)
	readByte, err := i.f.Read(b)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if readByte != n {
		panic("not enough data")
	}
	return b
}

// Close ...
func (i *IFS) Close() {
	i.f.Close()
	i.f = nil
}

// OFS ...
type OFS struct {
	f *os.File
}

// NewOFS ... a constructor
func NewOFS() *OFS {
	return &OFS{}
}

// Store8 ...
func (o *OFS) Store8(b byte) error {
	bytesWritten, err := o.f.Write([]byte{b})
	if err != nil {
		panic(err)
	}
	if bytesWritten != 1 {
		panic("write error")
	}
	return nil
}

// Close ...
func (o *OFS) Close() {
	fileInfo, err := o.f.Stat()
	if err != nil {
		panic(err)
	}
	log.Printf("written size? %d", fileInfo.Size())
	o.f.Close()
	o.f = nil
}

// Serializer ...
type Serializer struct {
	data []byte
	pos  int
	ofs  *OFS
}

// NewSerializer ...
func NewSerializer(rawData []byte, rawPos int) *Serializer {
	s := Serializer{data: rawData, pos: rawPos}
	return &s
}

// ReadByteRaw ...
func (s *Serializer) ReadByteRaw() (byte, error) {
	v := s.data[s.pos]
	s.pos++
	return v, nil
}

// ReadByteSafe ...
func (s *Serializer) ReadByteSafe() byte {
	v, err := s.ReadByteRaw()
	if err != nil {
		panic(err)
	}
	return v
}

// ReadWordRaw ...
func (s *Serializer) ReadWordRaw() (int, error) {
	v := int(s.data[s.pos]) | (int(s.data[s.pos+1]) << 8)
	s.pos += 2
	return v, nil
}

// ReadWordSafe ...
func (s *Serializer) ReadWordSafe() int {
	v, err := s.ReadWordRaw()
	if err != nil {
		panic(err)
	}
	return v
}

// ReadDWordRaw ...
func (s *Serializer) ReadDWordRaw() (int, error) {
	v := mathutil.MakeDWordFromLittleEndianBytes(s.data[s.pos], s.data[s.pos+1], s.data[s.pos+2], s.data[s.pos+3])
	s.pos += 4
	return v, nil
}

// ReadDWordSafe ...
func (s *Serializer) ReadDWordSafe() int {
	v, err := s.ReadDWordRaw()
	if err != nil {
		panic(err)
	}
	return v
}

// ReadStringRaw ...
func (s *Serializer) ReadStringRaw(length int) (string, error) {
	v, err := s.ReadVecRaw(length)
	return string(v), err
}

// ReadStringSafe ...
func (s *Serializer) ReadStringSafe(length int) string {
	v, err := s.ReadStringRaw(length)
	if err != nil {
		panic(err)
	}
	return v
}

// SkipBytesRaw ...
func (s *Serializer) SkipBytesRaw(length int) error {
	s.pos += length
	return nil
}

// SkipBytesSafe ...
func (s *Serializer) SkipBytesSafe(length int) {
	if err := s.SkipBytesRaw(length); err != nil {
		panic(err)
	}
}

// ReadVecRaw ...
func (s *Serializer) ReadVecRaw(length int) ([]byte, error) {
	v := make([]byte, length)
	for p := range v {
		v[p] = s.data[s.pos+p]
	}

	s.pos += length
	return v, nil
}

// ReadVecSafe ...
func (s *Serializer) ReadVecSafe(length int) []byte {
	v, err := s.ReadVecRaw(length)
	if err != nil {
		panic(err)
	}
	return v
}

// SetOFS ...
func (s *Serializer) SetOFS(ofs *OFS) {
	s.ofs = ofs
}

// WriteZeroes ...
func (s *Serializer) WriteZeroes(startp int, endp int) error {
	if s.ofs == nil {
		log.Fatal()
	}
	for p := startp; p < endp; p++ {
		s.ofs.Store8(0)
	}
	return nil
}

// WriteByteSafe ...
func (s *Serializer) WriteByteSafe(v byte) {
	_ = s.WriteByte(v)
}

// WriteWordSafe ...
func (s *Serializer) WriteWordSafe(n int) {
	_ = s.WriteWord(n)
}

// WriteDWordSafe ...
func (s *Serializer) WriteDWordSafe(n int) {
	_ = s.WriteDWord(n)
}

// WriteStringSafe ...
func (s *Serializer) WriteStringSafe(str string, length int) {
	_ = s.WriteString(str, length)
}

// WriteVecSafe ...
func (s *Serializer) WriteVecSafe(v []byte, startp int, length int) {
	_ = s.WriteVec(v, startp, length)
}

// WriteByte ...
func (s *Serializer) WriteByte(v byte) error {
	if s.ofs == nil {
		log.Fatal()
	}
	s.ofs.Store8(v)
	return nil
}

// WriteWord ...
func (s *Serializer) WriteWord(v int) error {
	if s.ofs == nil {
		log.Fatal()
	}
	b0, b1 := mathutil.GetLittleEndianBytesFromWord(v)
	s.ofs.Store8(b0)
	s.ofs.Store8(b1)
	return nil
}

// WriteDWord ...
func (s *Serializer) WriteDWord(v int) error {
	if s.ofs == nil {
		log.Fatal()
	}
	b0, b1, b2, b3 := mathutil.GetLittleEndianBytesFromDWord(v)
	s.ofs.Store8(b0)
	s.ofs.Store8(b1)
	s.ofs.Store8(b2)
	s.ofs.Store8(b3)
	return nil
}

// WriteString ...
func (s *Serializer) WriteString(str string, length int) error {
	if s.ofs == nil {
		log.Fatal()
	}
	for i := 0; i < length; i++ {
		if i < len(str) {
			s.ofs.Store8(str[i])
		} else {
			s.ofs.Store8(0)
		}
	}
	return nil
}

// WriteVec ...
func (s *Serializer) WriteVec(v []byte, startp int, length int) error {
	if s.ofs == nil {
		log.Fatal()
	}
	for i := 0; i < length; i++ {
		s.ofs.Store8(v[startp+i])
	}
	return nil
}
