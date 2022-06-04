package mapper

import (
	"fmt"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

type Segment struct{}

func (Segment) Map(size int64, ebml ebml.Ebml) (*types.Segment, error) {
	var id uint32
	segmentStart := ebml.CurrPos

	var segment types.Segment
	id, _ = GetID(&ebml, 4)

	elementSize, err := ebml.GetSize()
	if err != nil {
		return nil, err
	}

	element := ebml.Specification.Data[id]

	if element == nil {
		return nil, fmt.Errorf("unknown element ID of %x", id)
	}

	if element.Name == "SeekHead" {
		seekElements(&ebml, segmentStart, size, &segment, id, elementSize)
		ebml.CurrPos = segmentStart + size
	} else {
		crawl(&ebml, size, &segment, id, elementSize)
	}

	return &segment, err
}

func crawl(ebml *ebml.Ebml, segmentSize int64, segment *types.Segment, initialElementId uint32, initialElementSize int64) {
	var err error

	endPos := ebml.CurrPos + segmentSize
	id := initialElementId
	elementSize := initialElementSize
	element := ebml.Specification.Data[initialElementId]
	for ebml.CurrPos < endPos {
		getSubElement(ebml, elementSize, element, segment)
		ebml.CurrPos += elementSize

		id, _ = GetID(ebml, 4)
		elementSize, err = ebml.GetSize()
		if err != nil {
			break
		}

		element = ebml.Specification.Data[id]
	}
}

func seekElements(ebml *ebml.Ebml, segmentStart int64, segmentSize int64, segment *types.Segment, seekHeadId uint32, seekHeadSize int64) error {
	getSubElement(ebml, seekHeadSize, ebml.Specification.Data[seekHeadId], segment)
	errors := make(chan error)

	for _, seekHead := range segment.SeekHeads {
		for _, seek := range seekHead.Seeks {
			seekElement(*ebml, int64(seek.SeekPosition)+segmentStart, segment, errors)
		}
	}

	return nil
}

func seekElement(ebml ebml.Ebml, elementPosition int64, segment *types.Segment, errors chan<- error) {
	ebml.CurrPos = int64(elementPosition)
	id, _ := GetID(&ebml, 4)
	size, err := ebml.GetSize()

	if err != nil {
		errors <- err
		return
	}

	getSubElement(&ebml, size, ebml.Specification.Data[id], segment)
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
