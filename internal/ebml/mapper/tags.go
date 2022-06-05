package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/pkg/types"
)

type tags struct{}

func (tags) Map(size int64, ebml ebml.Ebml) ([]types.Tag, error) {
	return nil, nil
}
