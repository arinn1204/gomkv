package mapper

import (
	"encoding/binary"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

type Mapper interface {
	Map(ebml ebml.Ebml, spec *specification.Ebml) (types.Header, error)
}

type Header struct{}

func (Header) Map(ebml ebml.Ebml, spec *specification.Ebml) (types.Header, error) {
	startPos := ebml.CurrPos
	headerSize, err := ebml.GetSize()

	if err != nil {
		return types.Header{}, err
	}

	header := types.Header{}

	for ebml.CurrPos < startPos+headerSize {
		id, err := getID(&ebml)

		if err != nil {
			return header, err
		}

		element := spec.Data[uint32(id)]
		err = process(&header, id, &ebml, element)

		if err != nil {
			return header, err
		}
	}

	return header, nil
}

func getID(ebml *ebml.Ebml) (uint16, error) {
	buf := make([]byte, 2)
	n, err := ebml.File.Read(ebml.CurrPos, buf)
	if err != nil {
		return 0, err
	}

	ebml.CurrPos += int64(n)

	return binary.BigEndian.Uint16(buf), nil
}
