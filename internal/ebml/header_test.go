package ebml

import (
	"errors"
	"testing"

	"github.com/arinn1204/gomkv/internal/filesystem/mocks"
	"github.com/arinn1204/gomkv/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getHeaderTestData() []byte {
	return []byte{
		26,
		69,
		223,
		163,
		163,
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
	ebml := &mocks.Reader{}
	reader := Ebml{
		File:              ebml,
		CurrPos:           0,
		SpecificationPath: "testdata/header_ebml.xml",
	}

	call := ebml.On("Read", mock.AnythingOfType("int64"), mock.Anything)

	call.Run(func(args mock.Arguments) {
		retArr := args.Get(1).([]byte)
		count := len(retArr)

		copy(retArr, getHeaderTestData()[alreadyRead:alreadyRead+int(count)])
		alreadyRead += int(count)
		call.Return(len(retArr), nil)
	})

	doc, err := reader.Read()

	assert.Nil(t, err)
	assert.Nil(t, doc.Segments)

	expectedHeader := types.Header{
		EBMLVersion:        1,
		EBMLReadVersion:    1,
		EBMLMaxIDLength:    4,
		EBMLMaxSizeLength:  8,
		DocType:            "matroska",
		DocTypeVersion:     4,
		DocTypeReadVersion: 2,
	}

	assert.Equal(t, expectedHeader, doc.Header)
}

func TestFailsWhenNotEbmlDocument(t *testing.T) {
	alreadyRead := 0
	ebml := &mocks.Reader{}
	reader := Ebml{
		File:              ebml,
		CurrPos:           0,
		SpecificationPath: "testdata/header_ebml.xml",
	}

	call := ebml.On("Read", mock.AnythingOfType("int64"), mock.Anything)

	testData := getHeaderTestData()[:3]
	call.Run(func(args mock.Arguments) {
		retArr := args.Get(1).([]byte)
		count := len(retArr)

		copy(retArr, testData)
		alreadyRead += int(count)
		call.Return(len(retArr), nil)
	})

	doc, err := reader.Read()
	assert.Equal(t, types.EbmlDocument{}, doc)
	assert.Equal(t, errors.New("incorrect type of file expected magic number found 1a45df00"), err)
}
