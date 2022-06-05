package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/pkg/types"
)

type info struct{}

func (info) Map(size int64, ebml ebml.Ebml) ([]types.Info, error) {
	return nil, nil
}
