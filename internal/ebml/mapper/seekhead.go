package mapper

import (
	"encoding/binary"
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
			var set func(*types.Seek, any)
			var err error
			switch element.Name {
			case "SeekID":
				set = func(v *types.Seek, a any) {
					val := binary.BigEndian.Uint32(a.([]byte))
					seek.SeekID = uint64(val)
				}
			case "SeekPosition":
				set = func(v *types.Seek, a any) {
					v.SeekPosition = a.(uint)
				}
			default:
				ebml.CurrPos = endPos
			}

			if set != nil {
				var data any
				data, err = getFieldData(id, endPos-ebml.CurrPos, ebml)
				set(seek, data)
			}
			return err
		},
	)
	return seek, err
}
