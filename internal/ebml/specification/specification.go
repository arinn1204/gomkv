package specification

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Ebml struct {
	Data         map[int64]EbmlData
	DocumentType string
	Version      int
}

type EbmlData struct {
	Name              string
	Type              string
	Range             string
	Default           int
	MinimumOccurances int
	MaximumOccurances int
}

type ebmlStructure struct {
	XMLName      xml.Name     `xml:"EBMLSchema"`
	DocumentType string       `xml:"docType,attr"`
	Version      int          `xml:"version,attr"`
	Elements     []elementXml `xml:"element"`
}

type elementXml struct {
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

var specificationFile string

func init() {
	specificationFile = "data/matroska_ebml.xml"
}

//GetSpecification is a method used to read the matroska specification and return a mapped form of it that is easier to parse.
//This form will have the structure of map[elementID]=EBMLInformation
//The information will contain all the necessary element info: Name, Type, Range, Default, MinOccur, MaxOccur
func GetSpecification() Ebml {
	structure := readSpecification()

	data := make(map[int64]EbmlData)
	ebml := Ebml{
		Version:      structure.Version,
		DocumentType: structure.DocumentType,
		Data:         data,
	}

	for _, element := range structure.Elements {
		bitSize := (len(element.ID) - 2) * 4
		id, _ := strconv.ParseInt(element.ID, 16, bitSize)
		data[id] = EbmlData{
			Name:              element.Name,
			Type:              element.Type,
			Range:             element.Range,
			Default:           element.Default,
			MinimumOccurances: element.MinimumOccurances,
			MaximumOccurances: element.MaximumOccurances,
		}
	}

	return ebml

}

func readSpecification() ebmlStructure {
	xmlFile, err := os.Open(specificationFile)

	if err != nil {
		log.Panicf("Failed to open specification. -- '%+v'", err)
	}

	defer xmlFile.Close()

	rawValue, _ := ioutil.ReadAll(xmlFile)

	var ebml ebmlStructure

	err = xml.Unmarshal(rawValue, &ebml)

	if err != nil {
		log.Panicf("Failed to parse the specification xml. -- '%+v'", err)
	}

	return ebml
}
