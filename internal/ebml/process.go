package ebml

import (
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/arinn1204/gomkv/internal/array"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
)

func process[T any](item *T, id uint16, ebml *Ebml, element specification.EbmlData) error {
	elemSize := ebml.GetSize()

	buf := make([]byte, elemSize)
	n, err := ebml.File.Read(ebml.CurrPos, buf)

	if err != nil {
		return err
	}

	ebml.CurrPos += int64(n)

	elems := reflect.ValueOf(item).Elem()
	field := elems.FieldByName(element.Name)
	return setElementData(buf, element, &field)
}

func setElementData(buf []byte, element specification.EbmlData, field *reflect.Value) error {
	switch element.Type {
	case "uinteger":
		paddedBuf := make([]byte, 8)
		array.Pad(buf, paddedBuf)
		data := binary.BigEndian.Uint64(paddedBuf)
		field.Set(reflect.ValueOf(uint(data)))
		return nil
	case "utf-8":
	case "string":
		field.Set(reflect.ValueOf(string(buf)))
		return nil
	case "binary":
		field.Set(reflect.ValueOf(buf))
		return nil
	case "date":
		paddedBuf := make([]byte, 8)
		array.Pad(buf, paddedBuf)
		data := binary.BigEndian.Uint64(paddedBuf)
		field.Set(reflect.ValueOf(data))
		return nil
	}

	return fmt.Errorf("failed to get data for %v", element.Type)
}
