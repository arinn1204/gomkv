package mapper

import (
	"testing"

	"github.com/arinn1204/gomkv/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestWillProperlySerializeASeekHead(t *testing.T) {
	seekHeadData := []byte{77, 187, 140, 83, 171, 132, 21, 73, 169, 102, 83, 172, 130, 16, 3, 77, 187, 140, 83, 171, 132, 22, 84, 174, 107, 83, 172, 130, 16, 131, 77, 187, 142, 83, 171, 132, 28, 83, 187, 107, 83, 172, 132, 46, 16, 255, 240, 77, 187, 142, 83, 171, 132, 18, 84, 195, 103, 83, 172, 132, 46, 17, 65, 154, 236, 79}

	reader := getMockData(seekHeadData)

	seekHeadRes, err := SeekHead{}.Map(int64(len(seekHeadData)), *reader)

	assert.Nil(t, err)

	expected := types.SeekHead{
		Seeks: []types.Seek{
			{
				SeekPosition: 0x1003,
				SeekID:       0x1549a966,
			},
			{
				SeekPosition: 0x1083,
				SeekID:       0x1654ae6b,
			},
			{
				SeekPosition: 0x2e10fff0,
				SeekID:       0x1c53bb6b,
			},
			{
				SeekPosition: 0x2e11419a,
				SeekID:       0x1254c367,
			},
		},
	}
	assert.Equal(t, expected, *seekHeadRes)
}
