package specification

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

var spec EbmlSepecification

func init() {
	spec = EbmlSpec{}
}

func TestCanSerializeTopLevelEbmlSpec(t *testing.T) {
	specificationFile = "testdata/basicEbml.xml"
	spec := spec.GetSpecification()

	ebml := Ebml{
		XMLName: xml.Name{
			Local: "EBMLSchema",
			Space: "urn:ietf:rfc:8794",
		},
		DocumentType: "matroska",
		Version:      4,
	}

	assert.Equal(t, ebml.XMLName, spec.XMLName)
}

func TestCanSerializeEbmlElements(t *testing.T) {
	specificationFile = "testdata/basicEbml.xml"

	spec := spec.GetSpecification()

	ebml := Ebml{
		XMLName: xml.Name{
			Local: "EBMLSchema",
		},
		Elements: []Element{
			{
				XMLName: xml.Name{
					Local: "element",
					Space: "urn:ietf:rfc:8794",
				},
				Name:              "EBMLMaxIDLength",
				Path:              "\\EBML\\EBMLMaxIDLength",
				ID:                "0x42F2",
				Type:              "uinteger",
				Range:             "4",
				Default:           4,
				MinimumOccurances: 1,
				MaximumOccurances: 1,
			},
		},
	}
	assert.Equal(t, len(ebml.Elements), len(spec.Elements))

	for index, element := range ebml.Elements {
		received := spec.Elements[index]
		assert.Equal(t, element, received)
	}
}

func TestWillPanicIfFileNotFound(t *testing.T) {
	specificationFile = "testdata/notFound.xml"
}
