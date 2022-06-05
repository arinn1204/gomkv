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
			elemErr := <-errorChan
			if elemErr == nil {
				continue
			}
			if err == nil {
				err = elemErr
			} else {
				err = errors.New(err.Error() + elemErr.Error())
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
		infos, infoErr := info{}.Map(size, *ebml)
		segment.Infos = infos
		if infoErr != nil {
			err = errors.New("info creation failed - " + infoErr.Error())
		}
	case "Tracks":
		tracks, trackErr := tracks{}.Map(size, *ebml)
		segment.Tracks = tracks
		if trackErr != nil {
			err = errors.New("tracks creation failed - " + trackErr.Error())
		}
	case "Tags":
		tags, tagErr := tags{}.Map(size, *ebml)
		segment.Tags = tags
		if tagErr != nil {
			err = errors.New("tags creation failed - " + tagErr.Error())
		}
	case "Cues":
		cues, cueErr := cues{}.Map(size, *ebml)
		segment.Points = cues
		if cueErr != nil {
			err = errors.New("cues creation failed - " + cueErr.Error())
		}
	case "SeekHead":
		seekHead, seekHeadErr := seekHead{}.Map(size, *ebml)
		segment.SeekHeads = append(segment.SeekHeads, *seekHead)
		if seekHeadErr != nil {
			err = errors.New("seekHead creation failed - " + seekHeadErr.Error())
		}
	case "Void":
		break
	}

	return err
}
