package mapper

import (
	"fmt"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/pkg/types"
)

type seekHead struct{}

func (seekHead) Map(size int64, ebml ebml.Ebml) (*types.SeekHead, error) {
	var err error
	id, err := GetID(&ebml, 2)

	if err != nil {
		return nil, err
	}

	element := ebml.Specification.Data[id]

	if element == nil || element.Name != "Seek" {
		return nil, fmt.Errorf("unknown element id of %x", id)
	}

	seeks, err := createSeeks(ebml, size)

	return &types.SeekHead{
		Seeks: seeks,
	}, err
}

func createSeeks(ebml ebml.Ebml, size int64) ([]types.Seek, error) {
	seeks := make([]types.Seek, 0)

	var err error
	id := uint32(0x4DBB)
	endPos := ebml.CurrPos + size

	for ebml.CurrPos < endPos && err == nil && id == 0x4DBB {
		seekSize, seekErr := ebml.GetSize()

		if seekErr != nil {
			err = fmt.Errorf("SeekHead error - %v", seekErr.Error())
			break
		}

		seek := types.Seek{}
		seekEnd := ebml.CurrPos + seekSize
		for ebml.CurrPos < seekEnd {
			elemId, elemErr := GetID(&ebml, 2)
			if elemErr != nil {
				err = elemErr
				break
			}

			process(&seek, uint16(elemId), &ebml)
		}
		if err != nil {
			break
		}
		seeks = append(seeks, seek)
		id, err = GetID(&ebml, 2)
	}

	return seeks, err
}
