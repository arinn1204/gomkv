package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/pkg/types"
)

type SeekHead struct{}

func (SeekHead) Map(size int64, ebml ebml.Ebml) ([]types.SeekHead, error) {
	return nil, nil
}
