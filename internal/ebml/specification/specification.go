package specification

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

//Ebml is the overarching xml element name
type Ebml struct {
	XMLName      xml.Name  `xml:"EBMLSchema"`
	DocumentType string    `xml:"docType,attr"`
	Version      int       `xml:"version,attr"`
	Elements     []Element `xml:"element"`
}

//Element is the xml element that each piece of hte specification is defined within
type Element struct {
	XMLName           xml.Name `xml:"element"`
	Name              string   `xml:"name,attr"`
	Path              string   `xml:"path,attr"`
	ID                string   `xml:"id,attr"`
	Type              string   `xml:"type,attr"`
	Range             string   `xml:"range,attr"`
	Default           int      `xml:"default,attr"`
	MinimumOccurances int      `xml:"minOccurs,attr"`
	MaximumOccurances int      `xml:"maxOccurs,attr"`
}

//EbmlSepecification is the interface that the specification module adheres to
type EbmlSepecification interface {
	GetSpecification() Ebml
}

//EbmlSpec is the struct implementing the modules interface
type EbmlSpec struct{}

var ebmlSpec EbmlSepecification
var specificationFile string

func init() {
	specificationFile = "data/matroska_ebml.xml"
	ebmlSpec = EbmlSpec{}
}

//GetSpecification will return a formatted form of the EBML specification
func (e EbmlSpec) GetSpecification() Ebml {
	xmlFile, err := os.Open(specificationFile)

	if err != nil {
		log.Panicf("Failed to open specification. -- '%+v'", err)
	}

	defer xmlFile.Close()

	rawValue, _ := ioutil.ReadAll(xmlFile)

	var ebml Ebml

	err = xml.Unmarshal(rawValue, &ebml)

	if err != nil {
		log.Panicf("Failed to parse the specification xml. -- '%+v'", err)
	}

	return ebml
}
