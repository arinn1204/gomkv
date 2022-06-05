package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/pkg/types"
)

type cues struct{}

func (cues) Map(size int64, ebml ebml.Ebml) ([]types.Point, error) {
	return nil, nil
}
