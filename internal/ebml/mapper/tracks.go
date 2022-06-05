package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/pkg/types"
)

type tracks struct{}

func (tracks) Map(size int64, ebml ebml.Ebml) ([]types.Track, error) {
	return nil, nil
}
