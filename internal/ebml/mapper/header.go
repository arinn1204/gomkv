package mapper

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/arinn1204/gomkv/internal/array"
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

type Mapper[T any] interface {
	Map(ebml ebml.Ebml, spec *specification.Ebml) (T, error)
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
		id, err := getID(&ebml, 2)

		if err != nil {
			return header, err
		}

		element := spec.Data[id]
		err = process(&header, uint16(id), &ebml, element)

		if err != nil {
			return header, err
		}
	}

	return header, nil
}

func getID(ebml *ebml.Ebml, count int) (uint32, error) {
	buf := make([]byte, count)
	n, err := ebml.File.Read(ebml.CurrPos, buf)
	if err != nil {
		if err == io.EOF {
			return 0, err
		}
		return 0, fmt.Errorf("getID failed to read: %v", err.Error())
	}

	ebml.CurrPos += int64(n)

	paddedBuf := make([]byte, 4)
	array.Pad(buf, paddedBuf)

	return binary.BigEndian.Uint32(paddedBuf), nil
}
