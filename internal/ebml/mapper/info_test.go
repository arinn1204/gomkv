package mapper

import (
	"testing"

	"github.com/arinn1204/gomkv/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	infoData := []byte{42, 215, 177, 131, 15, 66, 64, 77, 128, 163, 108, 105, 98, 101, 98, 109, 108, 32, 118, 49, 46, 52, 46, 50, 32, 43, 32, 108, 105, 98, 109, 97, 116, 114, 111, 115, 107, 97, 32, 118, 49, 46, 54, 46, 52, 87, 65, 166, 109, 107, 118, 109, 101, 114, 103, 101, 32, 118, 54, 51, 46, 48, 46, 48, 32, 40, 39, 69, 118, 101, 114, 121, 116, 104, 105, 110, 103, 39, 41, 32, 54, 52, 45, 98, 105, 116, 68, 137, 132, 73, 17, 26, 0, 68, 97, 136, 9, 87, 150, 253, 18, 182, 46, 0, 115, 164, 144, 238, 6, 5, 118, 217, 31, 105, 95, 33, 117, 255, 128, 185, 32, 38, 185}
	reader := getMockData(infoData)

	info, err := info{}.Map(0, *reader)
	assert.Nil(t, err)

	expected := []types.Info{
		{},
	}
	assert.Equal(t, expected, info)
}
