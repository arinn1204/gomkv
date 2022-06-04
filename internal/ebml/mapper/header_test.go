package mapper

import (
	"io"
	"testing"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/filesystem/mocks"
	"github.com/arinn1204/gomkv/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getHeaderTestData() []byte {
	return []byte{
		66,
		134,
		129,
		1,
		66,
		247,
		129,
		1,
		66,
		242,
		129,
		4,
		66,
		243,
		129,
		8,
		66,
		130,
		136,
		109,
		97,
		116,
		114,
		111,
		115,
		107,
		97,
		66,
		135,
		129,
		4,
		66,
		133,
		129,
		2,
	}
}

func TestCanProperlySerializeHeader(t *testing.T) {
	alreadyRead := 0
	mockReader := &mocks.Reader{}

	call := mockReader.On("Read", mock.AnythingOfType("int64"), mock.Anything)

	startPos := int64(0)
	call.Run(func(args mock.Arguments) {
		retArr := args.Get(1).([]byte)
		count := len(retArr)

		startArr := alreadyRead
		if startPos == args.Get(0).(int64) && alreadyRead > 0 {
			startArr--
			count--
		}
		copy(retArr, getHeaderTestData()[startArr:alreadyRead+int(count)])
		alreadyRead += int(count)
		call.Return(len(retArr), nil)
		startPos = args.Get(0).(int64)
	})

	spec, _ := specification.GetSpecification("../testdata/header_ebml.xml")
	reader := ebml.Ebml{
		File:          mockReader,
		CurrPos:       0,
		Specification: spec,
	}

	doc, err := Header{}.Map(int64(len(getHeaderTestData())), reader)

	assert.Nil(t, err)

	expectedHeader := types.Header{
		EBMLVersion:        1,
		EBMLReadVersion:    1,
		EBMLMaxIDLength:    4,
		EBMLMaxSizeLength:  8,
		DocType:            "matroska",
		DocTypeVersion:     4,
		DocTypeReadVersion: 2,
	}

	assert.Equal(t, expectedHeader, *doc)
}

func TestReturnsOutWhenEndOfFile(t *testing.T) {
	alreadyRead := 0
	mockReader := &mocks.Reader{}

	call := mockReader.On("Read", mock.AnythingOfType("int64"), mock.Anything)

	testData := getHeaderTestData()[:4]
	call.Run(func(args mock.Arguments) {
		retArr := args.Get(1).([]byte)

		copy(retArr, testData)
		if alreadyRead > 0 {
			call.Return(
				0,
				io.EOF,
			)
		} else {
			call.Return(alreadyRead, nil)
		}
		alreadyRead++
	})

	spec, _ := specification.GetSpecification("../testdata/header_ebml.xml")
	reader := ebml.Ebml{
		File:          mockReader,
		CurrPos:       0,
		Specification: spec,
	}

	doc, err := Header{}.Map(int64(len(getHeaderTestData())), reader)

	assert.Nil(t, doc)
	assert.Equal(t, io.EOF, err)
}
