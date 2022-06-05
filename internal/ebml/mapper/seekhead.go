package mapper

import (
	"fmt"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

type seekHead struct{}

func (seekHead) Map(size int64, ebml ebml.Ebml) (*types.SeekHead, error) {
	seekHead := new(types.SeekHead)
	err := readUntil(
		&ebml,
		ebml.CurrPos+size,
		func(id uint32, endPos int64, element *specification.EbmlData) error {
			if element.Name != "Seek" {
				ebml.CurrPos = endPos
				return fmt.Errorf("expecting element Seek but got '%v'", element.Name)
			}
			seek, err := createSeek(&ebml, endPos)
			seekHead.Seeks = append(seekHead.Seeks, seek)
			return err
		},
	)

	return seekHead, err
}

func createSeek(ebml *ebml.Ebml, endPosition int64) (*types.Seek, error) {
	seek := new(types.Seek)
	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPos int64, element *specification.EbmlData) error {
			return process(seek, id, endPos-ebml.CurrPos, ebml)
		},
	)
	return seek, err
}
