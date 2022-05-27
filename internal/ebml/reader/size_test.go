package reader

import (
	"fmt"
	"testing"

	"github.com/arinn1204/gomkv/internal/io/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testData struct {
	size     int64
	numCalls int
}

func TestGetSizeWithDifferentWidths(t *testing.T) {
	sizes := []testData{
		{
			size:     127,
			numCalls: 1,
		},
		{
			size:     16147,
			numCalls: 2,
		},
		{
			size:     2036498,
			numCalls: 2,
		},
		{
			size:     252908049,
			numCalls: 2,
		},
		{
			size:     30384722192,
			numCalls: 2,
		},
		{
			size:     3380442370063,
			numCalls: 2,
		},
		{
			size:     865393246736142,
			numCalls: 2,
		},
		{
			size:     5367889050668557,
			numCalls: 3,
		},
	}

	for i, expected := range sizes {
		ebml := &mocks.Reader{}
		reader := EbmlReader{
			Reader:  ebml,
			CurrPos: 0,
		}
		width, _ := widthMap.GetInverse(i + 1)

		data := []byte{
			byte(width.(int)),
			byte(19),
			byte(18),
			byte(17),
			byte(16),
			byte(15),
			byte(14),
			byte(13),
			byte(12),
		}

		alreadyRead := 0

		ebml.On("Read", mock.AnythingOfType("uint"), mock.Anything).
			Return(1).
			Run(func(args mock.Arguments) {
				count := args.Get(0).(int)
				arr := args.Get(1).([]byte)
				//mimic the reading of data and copy into the argument array
				copy(arr, data[alreadyRead:count])

				alreadyRead += count
			})

		testName := fmt.Sprintf("GetSize(width=%v)", i)

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
