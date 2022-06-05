package mapper

import (
	"errors"
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
		seErr := getSubElement(ebml, elementSize, element, segment)
		ebml.CurrPos += elementSize

		id, _ = GetID(ebml, 4)
		elementSize, err = ebml.GetSize()
		if err != nil || seErr != nil {
			break
		}

		element = ebml.Specification.Data[id]
	}
}

func seekElements(ebml *ebml.Ebml, segmentStart int64, segmentSize int64, segment *types.Segment, seekHeadId uint32, seekHeadSize int64) error {
	err := getSubElement(ebml, seekHeadSize, ebml.Specification.Data[seekHeadId], segment)

	if err != nil {
		return err
	}

	for _, seekHead := range segment.SeekHeads {
		errorChan := make(chan error, len(seekHead.Seeks))
		for _, seek := range seekHead.Seeks {
			go seekElement(*ebml, int64(seek.SeekPosition)+segmentStart, segment, errorChan)
		}

		for i := 0; i < len(seekHead.Seeks); i++ {
			c := <-errorChan
			if c == nil {
				continue
			}
			if err == nil {
				err = c
			} else {
				err = errors.New(err.Error() + c.Error())
			}
		}

		close(errorChan)
	}

	return err
}

func seekElement(ebml ebml.Ebml, elementPosition int64, segment *types.Segment, errors chan<- error) {
	ebml.CurrPos = int64(elementPosition)
	id, _ := GetID(&ebml, 4)
	size, err := ebml.GetSize()

	if err != nil {
		errors <- err
		return
	}

	errors <- getSubElement(&ebml, size, ebml.Specification.Data[id], segment)
}

func getSubElement(ebml *ebml.Ebml, size int64, element *specification.EbmlData, segment *types.Segment) error {
	var err error
	switch element.Name {
	case "Info":
		segment.Infos = make([]types.Info, 1)
	case "Tracks":
		segment.Tracks = make([]types.Track, 1)
	case "Tags":
		segment.Tags = make([]types.Tag, 1)
	case "Cues":
		segment.Points = make([]types.Point, 1)
	case "SeekHead":
		seekHead, seekHeadErr := SeekHead{}.Map(size, *ebml)
		segment.SeekHeads = append(segment.SeekHeads, *seekHead)
		if seekHeadErr != nil {
			err = errors.New("seekHead creation failed - " + seekHeadErr.Error())
		}
	case "Void":
		break
	}

	return err
}
