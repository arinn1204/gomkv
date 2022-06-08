package mapper

import (
	"testing"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
	"github.com/stretchr/testify/assert"
)

var expectedTags []*types.Tag

func init() {
	expectedTags = []*types.Tag{
		{
			Targets: &types.Target{
				TagAttachmentUID: 1234,
				TargetTypeValue:  1,
				TargetType:       "type",
				TagTrackUID:      1111,
				TagEditionUID:    2222,
			},
		},
		{
			Targets: &types.Target{
				TagAttachmentUID: 1234,
				TargetTypeValue:  1,
				TargetType:       "type",
				TagTrackUID:      1111,
				TagEditionUID:    2222,
			},
			Tags: []*types.SimpleTag{{
				TagName:         "Taggity Tag",
				TagLanguage:     "en",
				TagLanguageIETF: "enover",
				TagDefault:      0,
				TagDefaultBogus: "bogusdata",
				TagString:       "stringity strings",
				Child: &types.SimpleTag{
					TagName:         "Taggity Tag2",
					TagLanguage:     "fr",
					TagLanguageIETF: "fr",
					TagDefault:      1,
					TagDefaultBogus: "bogusdata",
					TagString:       "bonjour como sa va",
				},
			}, {
				TagName:   "Taggity Tag2",
				TagString: "stringity strings",
			}},
		},
	}
}

func TestTagsFormatting(t *testing.T) {
	var elementid uint32
	var endPosition int64
	var expectedData []byte

	read = func(ebml *ebml.Ebml, data []byte) (int, error) {
		copy(data, expectedData)
		return len(expectedData), nil
	}

	readUntil = func(
		ebml *ebml.Ebml,
		end int64,
		process func(id uint32, endPosition int64, element *specification.EbmlData) error) error {
		return process(elementid, endPosition, ebml.Specification.Data[elementid])
	}

	_, err := tags{}.Map(64, *testEbmlObj)
	assert.Nil(t, err)
}
