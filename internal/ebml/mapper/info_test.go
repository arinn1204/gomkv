package mapper

import (
	"encoding/binary"
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
	//infoData := []byte{42, 215, 177, 131, 15, 66, 64, 77, 128, 163, 108, 105, 98, 101, 98, 109, 108, 32, 118, 49, 46, 52, 46, 50, 32, 43, 32, 108, 105, 98, 109, 97, 116, 114, 111, 115, 107, 97, 32, 118, 49, 46, 54, 46, 52, 87, 65, 166, 109, 107, 118, 109, 101, 114, 103, 101, 32, 118, 54, 51, 46, 48, 46, 48, 32, 40, 39, 69, 118, 101, 114, 121, 116, 104, 105, 110, 103, 39, 41, 32, 54, 52, 45, 98, 105, 116, 68, 137, 132, 73, 17, 26, 0, 68, 97, 136, 9, 87, 150, 253, 18, 182, 46, 0, 115, 164, 144, 238, 6, 5, 118, 217, 31, 105, 95, 33, 117, 255, 128, 185, 32, 38, 185}
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
		}

		return uuid.Nil
	}

	var expected interface{}
	for i := 0; i < 8; i++ {
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
