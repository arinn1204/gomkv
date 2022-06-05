package mapper

import (
	"encoding/binary"
	"testing"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestWillProperlySerializeASeekHead(t *testing.T) {
	//seekHeadData := []byte{77, 187, 140, 83, 171, 132, 21, 73, 169, 102, 83, 172, 130, 16, 3, 77, 187, 140, 83, 171, 132, 22, 84, 174, 107, 83, 172, 130, 16, 131, 77, 187, 142, 83, 171, 132, 28, 83, 187, 107, 83, 172, 132, 46, 16, 255, 240, 77, 187, 142, 83, 171, 132, 18, 84, 195, 103, 83, 172, 132, 46, 17, 65, 154, 236, 79}
	idCount := 0
	GetID = func(ebml *ebml.Ebml, maxCount int) (uint32, error) {
		var id uint32
		switch idCount {
		case 0:
			id = 0x4DBB
		case 1:
			id = 0x53AC //seek ID
		case 2:
			id = 0x53AB //seek Position
		}
		idCount++
		return id, nil
	}

	read = func(ebml *ebml.Ebml, data []byte) (int, error) {
		switch idCount {
		case 2:
			binary.BigEndian.PutUint32(data, 0x1003)
		case 3:
			binary.BigEndian.PutUint64(data, 0x1549a966)
		}
		return len(data), nil
	}

	getSize = func(ebml *ebml.Ebml) (int64, error) {
		var size int64
		switch idCount {
		case 1:
			size = 12
		case 2:
			size = 4
		case 3:
			size = 8
		}

		return size, nil
	}

	seekHeadRes, err := seekHead{}.Map(int64(12), *testEbmlObj)

	assert.Nil(t, err)

	expected := types.SeekHead{
		Seeks: []types.Seek{
			{
				SeekPosition: 0x1003,
				SeekID:       0x1549a966,
			},
		},
	}
	assert.Equal(t, expected, *seekHeadRes)
}
