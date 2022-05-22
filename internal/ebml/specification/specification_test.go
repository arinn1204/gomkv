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
	}

	if reflect.DeepEqual(ebml.XMLName, spec.XMLName) {
		t.Errorf("Expected '%v' but received '%v' instead.", fmt.Sprintf("%v", ebml.XMLName), fmt.Sprintf("%v", ebml.XMLName))
	}
}
