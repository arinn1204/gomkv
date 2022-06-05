package mapper

import (
	"testing"

	"github.com/arinn1204/gomkv/pkg/types"
	"github.com/stretchr/testify/assert"
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
	reader := getMockData(getHeaderTestData())

	doc, err := Header{}.Map(int64(len(getHeaderTestData())), *reader)

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
