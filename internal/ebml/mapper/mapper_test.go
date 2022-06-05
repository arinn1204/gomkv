package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/filesystem/mocks"
	"github.com/stretchr/testify/mock"
)

var getTestData func() []byte
var testEbmlObj *ebml.Ebml

func init() {
	spec, _ := specification.GetSpecification("testdata/matroska_ebml.xml")
	testEbmlObj = &ebml.Ebml{Specification: spec}
}

func getMockData(data []byte) *ebml.Ebml {
	mockReader := &mocks.Reader{}
	call := mockReader.On("Read", mock.AnythingOfType("int64"), mock.Anything)

	startPos := int64(0)
	alreadyRead := 0
	call.Run(func(args mock.Arguments) {
		retArr := args.Get(1).([]byte)
		count := len(retArr)

		startArr := alreadyRead
		if startPos == args.Get(0).(int64) && alreadyRead > 0 {
			startArr--
			count--
		}

		copy(retArr, data[startArr:alreadyRead+int(count)])
		alreadyRead += int(count)
		call.Return(len(retArr), nil)
		startPos = args.Get(0).(int64)
	})

	spec, _ := specification.GetSpecification("testdata/matroska_ebml.xml")
	reader := ebml.Ebml{
		File:          mockReader,
		CurrPos:       0,
		Specification: spec,
	}

	return &reader
}
