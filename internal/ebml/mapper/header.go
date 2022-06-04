package mapper

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/arinn1204/gomkv/internal/array"
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/pkg/types"
)

type Mapper[T any] interface {
	Map(size int64, ebml ebml.Ebml) (*T, error)
}

type Header struct{}

func (Header) Map(size int64, ebml ebml.Ebml) (*types.Header, error) {
	startPos := ebml.CurrPos

	header := types.Header{}

	for ebml.CurrPos < startPos+size {
		id, err := GetID(&ebml, 2)

		if err != nil {
			return nil, err
		}

		element := ebml.Specification.Data[id]
		if err = process(&header, uint16(id), &ebml, element); err != nil {
			return nil, err
		}
	}

	return &header, nil
}

//GetID is a function that will return the ID of the following EBML element
func GetID(ebml *ebml.Ebml, count int) (uint32, error) {
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
