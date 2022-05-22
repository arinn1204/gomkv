package specification

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type Ebml struct {
	XMLName      xml.Name  `xml:"EBMLSchema"`
	DocumentType string    `xml:"docType,attr"`
	Version      int       `xml:"version,attr"`
	Elements     []Element `xml:"element"`
}

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

type EbmlSepecification interface {
	GetSpecification() Ebml
}

type EbmlSpec struct{}

var ebmlSpec EbmlSepecification
var specificationFile string

func init() {
	specificationFile = "data/matroska_ebml.xml"
	ebmlSpec = EbmlSpec{}
}

func (e EbmlSpec) GetSpecification() Ebml {
	xmlFile, err := os.Open(specificationFile)

	if err != nil {
		log.Fatal("Failed to open specification.")
	}

	defer xmlFile.Close()

	rawValue, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		log.Fatal("Failed to parse the specification xml.")
	}

	var ebml Ebml

	xml.Unmarshal(rawValue, &ebml)

	return ebml
}
