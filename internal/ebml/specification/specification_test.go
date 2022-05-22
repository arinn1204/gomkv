package specification

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"
)

func TestCanSerializeTopLevelEbmlSpec(t *testing.T) {
	spec := getSpecification()

	ebml := Ebml{
		XMLName: xml.Name{
			Local: "EBMLSchema",
		},
		DocumentType: "matroska",
		Version:      4,
	}

	if reflect.DeepEqual(ebml.XMLName, spec.XMLName) {
		t.Errorf("Expected '%v' but received '%v' instead.", fmt.Sprintf("%v", ebml.XMLName), fmt.Sprintf("%v", ebml.XMLName))
	}
}

func TestCanSerializeEbmlElements(t *testing.T) {
	specificationFile = "testdata/basicEbml.xml"
	spec := getSpecification()

	ebml := Ebml{
		XMLName: xml.Name{
			Local: "EBMLSchema",
		},
		Elements: []Element{
			{
				XMLName: xml.Name{
					Local: "element",
				},
				Name:              "EBMLMaxIDLength",
				Path:              "\\EBML\\EBMLMaxIDLength",
				ID:                17138,
				Type:              "uinteger",
				Range:             "4",
				Default:           4,
				MinimumOccurances: 1,
				MaximumOccurances: 1,
			},
		},
	}

	if len(ebml.Elements) != len(spec.Elements) {
		t.Errorf("Expected an array of %v elements but found %v instead.", len(ebml.Elements), len(spec.Elements))
		t.FailNow()
	}

	for index, element := range ebml.Elements {
		received := spec.Elements[index]
		if !reflect.DeepEqual(element, received) {
			t.Errorf("Expected element to be %+v but was %+v", element, received)
			t.Fail()
		}
	}
}
