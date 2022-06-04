package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/pkg/types"
)

type Segment struct{}

func (Segment) Map(size uint32, ebml ebml.Ebml) (*types.Segment, error) {
	endPos := ebml.CurrPos + int64(size)

	for ebml.CurrPos < endPos {
		id, err := GetID(&ebml, 4)

		if err != nil {
			break
		}

		elem := ebml.Specification.Data[id]

		_ = elem
	}

	return &types.Segment{}, nil
}
