package reader

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vishalkuo/bimap"
)

var widthMap *bimap.BiMap

func init() {
	widthMap = bimap.NewBiMap()
	widthMap.Insert(255, 1)
	widthMap.Insert(127, 2)
	widthMap.Insert(63, 3)
	widthMap.Insert(31, 4)
	widthMap.Insert(15, 5)
	widthMap.Insert(7, 6)
	widthMap.Insert(3, 7)
	widthMap.Insert(1, 8)
}

func TestGetWidth(t *testing.T) {
	firstByte := 255

	for firstByte > 0 {
		expected, _ := widthMap.Get(firstByte)
		testName := fmt.Sprintf("TestGetWidth(%v_firstByte,%v_expected)", firstByte, expected)
		t.Run(testName, func(t *testing.T) {

			width := getWidth(byte(firstByte))

			assert.Equal(t, expected, int(width))

			firstByte = firstByte >> 1
		})
	}

}
