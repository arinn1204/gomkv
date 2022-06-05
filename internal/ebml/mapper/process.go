package mapper

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"

	"github.com/arinn1204/gomkv/internal/array"
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
)

func process[T any](item *T, id uint32, ebml *ebml.Ebml) error {
	elemSize, err := getSize(ebml)

	if err != nil {
		ebml.CurrPos += elemSize
		return err
	}

	buf := make([]byte, elemSize)
	n, err := read(ebml, buf)

	if err != nil {
		return err
	}

	ebml.CurrPos += int64(n)
	element := ebml.Specification.Data[id]

	elems := reflect.ValueOf(item).Elem()
	field := elems.FieldByName(element.Name)
	return setElementData(buf, element, &field)
}

func setElementData(buf []byte, element *specification.EbmlData, field *reflect.Value) error {
	var err error
	switch element.Type {
	case "date":
		fallthrough
	case "binary":
		data := getData(buf)
		field.Set(reflect.ValueOf(data))
	case "uinteger":
		data := getData(buf)
		field.Set(reflect.ValueOf(uint(data)))
	case "float":
		data := getData(buf)
		value := math.Float32frombits(uint32(data))
		field.Set(reflect.ValueOf(value))
	case "utf-8":
		fallthrough
	case "string":
		field.Set(reflect.ValueOf(string(buf)))
	default:
		err = fmt.Errorf("failed to get data for %v", element.Type)
	}

	return err
}

func getData(buf []byte) uint64 {

	paddedBuf := make([]byte, 8)
	array.Pad(buf, paddedBuf)
	data := binary.BigEndian.Uint64(paddedBuf)
	return data
}
