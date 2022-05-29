package ebml

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"

	"github.com/arinn1204/gomkv/internal/array"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/filesystem"
	"github.com/arinn1204/gomkv/pkg/types"
)

//Ebml will contain the IoReader as well as the current position of this members stream
type Ebml struct {
	File              filesystem.Reader
	CurrPos           int64
	SpecificationPath string
}

var ebmlIdHex string

func init() {
	ebmlIdHex = "1A45DFA3"
}

func (ebml Ebml) Read() (types.EbmlDocument, error) {
	err := validateMagicNum(&ebml)
	if err != nil {
		return types.EbmlDocument{}, err
	}

	header, _ := createHeader(ebml)

	return types.EbmlDocument{
			Header:   header,
			Segments: make([]types.Segment, 1),
		},
		nil
}

func validateMagicNum(ebml *Ebml) error {
	idBuf := make([]byte, 4)
	n, err := ebml.File.Read(ebml.CurrPos, idBuf)

	if err != nil {
		return err
	}

	ebml.CurrPos += int64(n)

	id := binary.BigEndian.Uint32(idBuf)

	decEbmlId, _ := strconv.ParseUint(ebmlIdHex, 16, 32)

	if decEbmlId != uint64(id) {
		return fmt.Errorf("incorrect type of file expected magic number of %x but found %x", ebmlIdHex, id)
	}

	return nil
}

func setElementData(buf []byte, element specification.EbmlData, field *reflect.Value) error {
	switch element.Type {
	case "uinteger":
		paddedBuf := make([]byte, 8)
		array.Pad(buf, paddedBuf)
		data := binary.BigEndian.Uint64(paddedBuf)
		field.Set(reflect.ValueOf(uint(data)))
	case "utf-8":
	case "string":
		field.Set(reflect.ValueOf(string(buf)))
	case "binary":
		field.Set(reflect.ValueOf(buf))
	case "date":
		paddedBuf := make([]byte, 8)
		array.Pad(buf, paddedBuf)
		data := binary.BigEndian.Uint64(paddedBuf)
		field.Set(reflect.ValueOf(data))
	}

	return fmt.Errorf("failed to get data for %v", element.Type)
}
