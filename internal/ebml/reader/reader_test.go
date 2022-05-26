package reader

import (
	"fmt"
	"testing"

	"github.com/arinn1204/gomkv/internal/io/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetWidth(t *testing.T) {
	ebml := &mocks.Reader{}

	reader := EbmlReader{
		Reader:  ebml,
		CurrPos: 0,
	}

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
			ebml.On("Read", mock.AnythingOfType("uint"), mock.Anything).
				Return(1).
				Run(func(args mock.Arguments) {
					inputArr := args.Get(1).([]byte)
					inputArr[0] = byte(firstByte)
				})

			width := reader.GetWidth()

			assert.Equal(t, expectedMap[firstByte], int(width))

			firstByte = firstByte >> 1
		})
	}

}
