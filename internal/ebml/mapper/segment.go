package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

type Segment struct{}

func (Segment) Map(ebml ebml.Ebml, spec *specification.Ebml) ([]types.Segment, error) {
	segments := make([]types.Segment, 0)

	return segments, nil
}
