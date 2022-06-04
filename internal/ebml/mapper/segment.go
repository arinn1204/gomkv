package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

type Segment struct{}

func (Segment) Map(ebml ebml.Ebml, spec *specification.Ebml) ([]types.Segment, error) {
	err := skipHeader(&ebml)

	if err != nil {
		return nil, err
	}

	return make([]types.Segment, 0), nil
}

func skipHeader(ebml *ebml.Ebml) error {
	size, err := ebml.GetSize()

	if err != nil {
		return err
	}

	ebml.CurrPos += size

	return nil
}
