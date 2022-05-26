package reader

import (
	"fmt"
	"testing"

	"github.com/arinn1204/gomkv/internal/io/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testData struct {
	size      int64
	readCount int
	numCalls  int
}

func TestGetSizeWithDifferentWidths(t *testing.T) {
	sizes := []testData{
		{
			size:      1,
			readCount: 1,
			numCalls:  2,
		},
	}

	data := []byte{
		byte(6),
		byte(2),
		byte(3),
		byte(4),
		byte(5),
		byte(6),
		byte(7),
	}

	for i, expected := range sizes {
		testName := fmt.Sprintf("GetSize(width=%v)", i)
		ebml := &mocks.Reader{}
		reader := EbmlReader{
			Reader:  ebml,
			CurrPos: 0,
		}

		ebml.On("Read", mock.MatchedBy(expected.readCount), mock.Anything).
			Return(1).
			Run(func(args mock.Arguments) {
				count := args.Get(0).(int)
				arr := args.Get(1).([]byte)
				//mimic the reading of data and copy into the argument array
				copy(arr, data[0:count])
			})

		width, _ := widthMap.GetInverse(i)
		data[0] = byte(width.(int))

		t.Run(
			testName,
			func(t *testing.T) {
				size := reader.GetSize()
				ebml.AssertNumberOfCalls(t, "Read", expected.numCalls)
				assert.Equal(t, expected.size, size)
			},
		)
	}
}
