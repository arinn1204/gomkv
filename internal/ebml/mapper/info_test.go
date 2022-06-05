package mapper

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"testing"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/pkg/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestCanProperlyParseData(t *testing.T) {
	var elementId int
	var elementSize int
	var testData []byte

	GetID = func(ebml *ebml.Ebml, maxCount int) (uint32, error) {
		return uint32(elementId), nil
	}

	read = func(ebml *ebml.Ebml, data []byte) (int, error) {
		testData := getTestData()
		if len(data) > len(testData) {
			copy(data, testData)
		} else {
			copy(data, testData[:len(data)])
		}
		return len(data), nil
	}

	getSize = func(ebml *ebml.Ebml) (int64, error) {
		return int64(elementSize), nil
	}

	getTestData = func() []byte {
		return testData
	}

	getExpectedInfoValue := func(info *types.Info, index int) interface{} {
		switch index {
		case 0:
			return info.SegmentUID
		case 1:
			return info.PrevUID
		case 2:
			return info.NextUID
		case 3:
			return info.SegmentFamily
		case 4:
			return info.DateUTC
		case 5:
			return info.Duration
		case 6:
			return info.MuxingApp
		case 7:
			return info.WritingApp
		case 8:
			return info.TimestampScale
		}

		return uuid.Nil
	}

	var expected interface{}
	for i := 0; i < 9; i++ {
		//The test itself
		var testName string
		switch i {
		case 0:
			elementId = 0x73A4
			testName = "SegmentUID"
			testData = []byte{238, 6, 5, 118, 217, 31, 105, 95, 33, 117, 255, 128, 185, 32, 38, 185}
			expected, _ = uuid.FromBytes(testData)
		case 1:
			elementId = 0x3CB923
			testName = "PreviousUID"
			testData = []byte{238, 6, 5, 118, 217, 31, 105, 95, 33, 117, 255, 128, 185, 32, 38, 185}
			expected, _ = uuid.FromBytes(testData)
		case 2:
			elementId = 0x3EB923
			testName = "NextUID"
			testData = []byte{238, 6, 5, 118, 217, 31, 105, 95, 33, 117, 255, 128, 185, 32, 38, 185}
			expected, _ = uuid.FromBytes(testData)
		case 3:
			elementId = 0x4444
			testName = "SegmentFamily"
			testData = []byte{238, 6, 5, 118, 217, 31, 105, 95, 33, 117, 255, 128, 185, 32, 38, 185}
			expected, _ = uuid.FromBytes(testData)
		case 4:
			elementId = 0x4461
			testName = "DateUTC"
			testData = make([]byte, 8)
			binary.BigEndian.PutUint64(testData, 4312894731298)
			expected = uint64(4312894731298)
		case 5:
			elementId = 0x4489
			testName = "Duration"
			testData = make([]byte, 4)
			binary.BigEndian.PutUint32(testData, math.Float32bits(12346.123))
			expected = float32(12346.123)
		case 6:
			elementId = 0x4D80
			testName = "Muxing App"
			expected = "Muxing App"
			testData = []byte(expected.(string))
		case 7:
			elementId = 0x5741
			testName = "Writing App"
			expected = "Writing App"
			testData = []byte(expected.(string))
		case 8:
			elementId = 0x2AD7B1
			testName = "TimestampScale"
			testData = make([]byte, 8)
			binary.BigEndian.PutUint64(testData, 1_000_000)
			expected = uint(1_000_000)
		}

		t.Run(
			testName,
			func(t *testing.T) {
				elementSize = len(getTestData())
				info, err := info{}.Map(int64(elementSize), *testEbmlObj)
				assert.Nil(t, err)
				result := getExpectedInfoValue(info, i)
				assert.Equal(t, expected, result)
			},
		)
	}
}

func TestWillAddToErrors(t *testing.T) {
	elementId := 0x4D80
	expected := "Muxing App"
	testData := []byte(expected)
	elementSize := len(testData)

	for i := 0; i < 5; i++ {
		GetID = func(ebml *ebml.Ebml, maxCount int) (uint32, error) {
			return uint32(elementId), nil
		}

		read = func(ebml *ebml.Ebml, data []byte) (int, error) {
			testData := getTestData()
			if len(data) > len(testData) {
				copy(data, testData)
			} else {
				copy(data, testData[:len(data)])
			}
			return len(data), nil
		}

		getSize = func(ebml *ebml.Ebml) (int64, error) {
			return int64(elementSize), nil
		}

		getTestData = func() []byte {
			return testData
		}

		var expected error
		switch i {
		case 0:
			expected = errors.New("foobar")
			GetID = func(ebml *ebml.Ebml, maxCount int) (uint32, error) {
				return 0, errors.New("foobar")
			}
		case 1:
			expected = errors.New("unrecognized id of 0xFFFFF")
			GetID = func(ebml *ebml.Ebml, maxCount int) (uint32, error) {
				return 0xFFFFF, nil
			}
		case 2:
			expected = io.EOF
			getSize = func(ebml *ebml.Ebml) (int64, error) {
				return 1, io.EOF
			}
		case 3:
			elementId = 0x4444
			getSize = func(ebml *ebml.Ebml) (int64, error) {
				return 1, io.EOF
			}
			testData = []byte{238, 6, 5, 118, 217, 31, 105, 95, 33, 117, 255, 128, 185, 32, 38, 185}
			expected = io.EOF
		case 4:
			elementId = 0x4444
			testData = []byte{238, 7, 6, 5, 118, 217, 31, 105, 95, 33, 117, 255, 128, 185, 32, 38, 185}
			expected = errors.New("invalid UUID (got 17 bytes)")
		}

		t.Run(
			fmt.Sprintf("TestWillAddToErrors %v", i),
			func(t *testing.T) {
				elementSize = len(getTestData())
				_, err := info{}.Map(int64(elementSize), *testEbmlObj)
				assert.NotNil(t, err)
				assert.Equal(t, expected, err)
			},
		)
	}
}
func TestWillSkipUndefinedElements(t *testing.T) {
	elementId := 0x42F7

	GetID = func(ebml *ebml.Ebml, maxCount int) (uint32, error) {
		return uint32(elementId), nil
	}

	getSize = func(ebml *ebml.Ebml) (int64, error) {
		return 1, nil
	}

	info, err := info{}.Map(int64(1), *testEbmlObj)
	assert.Nil(t, err)
	assert.Equal(t, types.Info{}, *info)

}
