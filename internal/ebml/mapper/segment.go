package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

type Segment struct{}

func (Segment) Map(size int64, ebml ebml.Ebml) (*types.Segment, error) {
	endPos := ebml.CurrPos + size
	var err error
	var elementSize int64
	var id uint32

	var segment types.Segment

	for ebml.CurrPos < endPos {
		id, _ = GetID(&ebml, 4)

		elementSize, err = ebml.GetSize()
		if err != nil {
			break
		}

		element := ebml.Specification.Data[id]
		getSubElement(&ebml, elementSize, element, &segment)

		ebml.CurrPos += elementSize
	}

	return &segment, err
}

func getSubElement(ebml *ebml.Ebml, size int64, element *specification.EbmlData, segment *types.Segment) chan<- error {
	errorChan := make(chan error)
	switch element.Name {
	case "SeekHead":
		func() {
			seekHead, err := SeekHead{}.Map(size, *ebml)
			segment.SeekHeads = append(segment.SeekHeads, *seekHead)
			if err != nil {
				errorChan <- err
			}
		}()
	case "Void":
		break
	}

	return errorChan
}
