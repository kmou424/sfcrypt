package v2

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"github.com/kmou424/ero"
	"github.com/kmou424/sfcrypt/app/version"
	"io"
	"os"
	"reflect"
)

func init() {
	gob.Register(SFHeader{})
}

// SFHeader 128 bytes magic header
type SFHeader struct {
	Magic   [16]byte
	Version [16]byte
	// reserved for future
	Reserve [96]byte
}

var DefHeader = func() *SFHeader {
	header := &SFHeader{}
	copy(header.Magic[:], "sfcrypt")
	copy(header.Version[:], version.GetVersion())
	return header
}()

var MaxHeaderSize = func() int {
	header := &SFHeader{}
	bytes, err := header.Bytes()
	if err != nil {
		panic(ero.Wrap(err, "failed to calculate max header size").Error())
	}
	// empty gob size aka reserved for no field size
	baseSize := len(bytes)

	// only exported field can be encoded
	exportedNum := 0
	structRef := reflect.TypeOf(header).Elem()
	for i := 0; i < structRef.NumField(); i++ {
		field := structRef.Field(i)
		// filter unexported fields because of gob will never encode them
		if !field.IsExported() {
			continue
		}

		exportedNum += 1
		baseSize -= int(field.Type.Size())
		baseSize -= len(field.Name)
	}
	// remove sign for exported fields
	baseSize -= exportedNum * 2

	// add sign for all fields
	maxSize := baseSize + 2*structRef.NumField()
	for i := 0; i < structRef.NumField(); i++ {
		field := structRef.Field(i)
		maxSize += int(field.Type.Size())
		baseSize += len(field.Name)
	}
	// edge case: if byte value > 127, size should be twice
	maxSize *= 2

	return maxSize
}()

func (f *SFHeader) Bytes() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(f)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (f *SFHeader) Parse(buf []byte) error {
	dec := gob.NewDecoder(bytes.NewReader(buf))
	err := dec.Decode(f)
	if err != nil {
		return ero.Newf("parse header error: %s", err.Error())
	}
	return nil
}

func (f *SFHeader) ReadFromFile(file *os.File, fallbackOnFailed bool) (n int, err error) {
	if fallbackOnFailed {
		offset, innerErr := file.Seek(0, io.SeekCurrent)
		if innerErr != nil {
			return 0, ero.Wrap(innerErr, "seek failed")
		}
		defer func() {
			if err != nil {
				_, innerErr = file.Seek(offset, io.SeekStart)
				if innerErr != nil {
					err = errors.Join(err, ero.Wrap(innerErr, "seek failed"))
				}
			}
		}()
	}

	lengthBuf := make([]byte, 4)
	n, err = file.Read(lengthBuf)
	if n != 4 {
		return n, ero.New("can't read header length")
	}
	if err != nil {
		return 0, err
	}

	length := int(binary.BigEndian.Uint32(lengthBuf))

	// check max length to avoid too large header
	if length > MaxHeaderSize {
		return 0, ero.Newf("header size %d exceeds max allowed size %d", length, MaxHeaderSize)
	}

	buf := make([]byte, length)
	n, err = file.Read(buf)
	if n != length {
		return n, ero.Newf("requires length %d bytes, but actually read %d", length, n)
	}
	if err != nil {
		return 0, err
	}

	return len(lengthBuf) + length, f.Parse(buf)
}

func (f *SFHeader) WriteToFile(file *os.File) (int, error) {
	buf, err := f.Bytes()
	if err != nil {
		return 0, err
	}

	lengthBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBuf, uint32(len(buf)))

	writeBuf := append(lengthBuf, buf...)
	n, err := file.Write(writeBuf)
	if err != nil {
		return n, err
	}

	return n, nil
}
