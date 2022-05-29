package ebml

import (
	"encoding/binary"
	"reflect"

	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

func createHeader(ebml Ebml) (types.Header, error) {
	startPos := ebml.CurrPos
	headerSize := ebml.GetSize()

	header := types.Header{}
	spec := specification.GetSpecification(ebml.SpecificationPath)

	for ebml.CurrPos < startPos+headerSize {
		id, err := getId(&ebml)

		if err != nil {
			return header, err
		}

		err = process(&header, id, &ebml, spec)

		if err != nil {
			return header, err
		}
	}

	return header, nil
}

func getId(ebml *Ebml) (uint16, error) {
	buf := make([]byte, 2)
	n, err := ebml.File.Read(ebml.CurrPos, buf)
	if err != nil {
		return 0, err
	}

	ebml.CurrPos += int64(n)

	return binary.BigEndian.Uint16(buf), nil
}

func process(header *types.Header, id uint16, ebml *Ebml, spec specification.Ebml) error {
	elemSize := ebml.GetSize()
	element := spec.Data[uint32(id)]

	buf := make([]byte, elemSize)
	n, err := ebml.File.Read(ebml.CurrPos, buf)

	if err != nil {
		return err
	}

	ebml.CurrPos += int64(n)

	elems := reflect.ValueOf(header).Elem()
	field := elems.FieldByName(element.Name)
	setElementData(buf, element, &field)

	return nil
}
