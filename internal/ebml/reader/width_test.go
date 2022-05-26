package reader

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWidth(t *testing.T) {
	firstByte := 255
	expectedMap := map[int]int{
		255: 1,
		127: 2,
		63:  3,
		31:  4,
		15:  5,
		7:   6,
		3:   7,
		1:   8,
	}

	for firstByte > 0 {
		testName := fmt.Sprintf("TestGetWidth(%v_firstByte,%v_expected)", firstByte, expectedMap[firstByte])
		t.Run(testName, func(t *testing.T) {

			width := getWidth(byte(firstByte))

			assert.Equal(t, expectedMap[firstByte], int(width))

			firstByte = firstByte >> 1
		})
	}

}
