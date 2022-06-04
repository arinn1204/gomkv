package mapper

import (
	"fmt"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

type Segment struct{}

func (Segment) Map(ebml ebml.Ebml, spec *specification.Ebml) ([]types.Segment, error) {
	//skip the header as that is being processed elsewhere
	if err := skipNextElement(&ebml); err != nil {
		return nil, fmt.Errorf("Map failed to create a segment. %v", err.Error())
	}

	id, err := getID(&ebml, 4)

	if err != nil {
		return nil, err
	}

	if spec.Data[uint32(id)].Name != "Segment" {
		return nil, fmt.Errorf("expected 'Segment' but found %v instead", spec.Data[uint32(id)].Name)
	}

	segSize, err := ebml.GetSize()

	if err != nil {
		return nil, err
	}

	endPos := ebml.CurrPos + segSize

	for ebml.CurrPos < endPos {
		id, err := getID(&ebml, 4)

		if err != nil {
			break
		}

		elem := spec.Data[id]

		_ = elem
	}

	segments := make([]types.Segment, 0)

	return segments, nil
}

func skipNextElement(ebml *ebml.Ebml) error {
	size, err := ebml.GetSize()

	if err != nil {
		return fmt.Errorf("skipNextElement failed - %v", err.Error())
	}

	ebml.CurrPos += size

	return nil
}
