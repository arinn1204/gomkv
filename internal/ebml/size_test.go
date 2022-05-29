package ebml

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/arinn1204/gomkv/internal/filesystem/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testData struct {
	size     int64
	numCalls int
}

func getData(width int) []byte {
	w, _ := widthMap.GetInverse(width)
	return []byte{
		byte(w.(int)),
		byte(19),
		byte(18),
		byte(17),
		byte(16),
		byte(15),
		byte(14),
		byte(13),
		byte(12),
	}
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
			numCalls: 2,
		},
	}

	for i, expected := range sizes {
		ebml := &mocks.Reader{}
		reader := Ebml{
			File:    ebml,
			CurrPos: 0,
		}
		data := getData(i + 1)

		alreadyRead := 0

		call := ebml.On("Read", mock.AnythingOfType("uint"), mock.Anything)

		call.Run(func(args mock.Arguments) {
			retArr := args.Get(1).([]byte)
			count := len(retArr)

			copy(retArr, data[alreadyRead:alreadyRead+int(count)])
			alreadyRead += int(count)
			call.Return(len(retArr), nil)
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

func TestEndianess(t *testing.T) {

	data := getData(6)
	ebml := &mocks.Reader{}
	reader := Ebml{
		File:    ebml,
		CurrPos: 0,
	}
	call := ebml.On("Read", mock.AnythingOfType("uint"), mock.Anything)

	alreadyRead := 0

	call.Run(func(args mock.Arguments) {
		retArr := args.Get(1).([]byte)
		count := len(retArr)

		copy(retArr, data[alreadyRead:alreadyRead+int(count)])
		alreadyRead += int(count)
		call.Return(len(retArr), nil)
	})

	result := reader.GetSize()

	be := make([]byte, 8)
	binary.BigEndian.PutUint64(be, uint64(result))

	assert.Equal(t, be, []byte{0, 0, 3, 19, 18, 17, 16, 15})
}

func TestReadReturnsZero(t *testing.T) {
	ebml := &mocks.Reader{}
	reader := Ebml{
		File:    ebml,
		CurrPos: 0,
	}
	ebml.On("Read", mock.AnythingOfType("uint"), mock.Anything).Return(0, nil)
	result := reader.GetSize()
	assert.Equal(t, int64(0), result)
}
