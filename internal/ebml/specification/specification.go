package specification

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
)

//Ebml is the mapping of the specification to an easier to use data structure
//The mapping will be from item id code to the element definition
//The doctype and version come directly from the specification used to parse the file
//not from the file itself
type Ebml struct {
	Data         map[uint32]EbmlData
	DocumentType string
	Version      int
}

//EbmlData is a structure that represents the element in the matroska specification
type EbmlData struct {
	Name              string
	Type              string
	Range             string
	Default           string
	MinimumOccurances int
	MaximumOccurances int
}

type ebmlStructure struct {
	XMLName      xml.Name     `xml:"EBMLSchema"`
	DocumentType string       `xml:"docType,attr"`
	Version      int          `xml:"version,attr"`
	Elements     []elementXML `xml:"element"`
}

type elementXML struct {
	XMLName           xml.Name `xml:"element"`
	Name              string   `xml:"name,attr"`
	Path              string   `xml:"path,attr"`
	ID                string   `xml:"id,attr"`
	Type              string   `xml:"type,attr"`
	Range             string   `xml:"range,attr"`
	Default           string   `xml:"default,attr"`
	MinimumOccurances int      `xml:"minOccurs,attr"`
	MaximumOccurances int      `xml:"maxOccurs,attr"`
}

//GetSpecification is a method used to read the matroska specification and return a mapped form of it that is easier to parse.
//This form will have the structure of map[elementID]=EBMLInformation
//The information will contain all the necessary element info: Name, Type, Range, Default, MinOccur, MaxOccur
func GetSpecification(path string) (*Ebml, error) {
	structure, err := readSpecification(path)

	if err != nil {
		return nil, err
	}

	data := make(map[uint32]EbmlData)
	ebml := Ebml{
		Version:      structure.Version,
		DocumentType: structure.DocumentType,
		Data:         data,
	}

	for _, element := range structure.Elements {
		bitSize := (len(element.ID) - 2) * 4
		id, err := strconv.ParseUint(element.ID[2:], 16, bitSize)

		if err != nil {
			continue
		}

		data[uint32(id)] = EbmlData{
			Name:              element.Name,
			Type:              element.Type,
			Range:             element.Range,
			Default:           element.Default,
			MinimumOccurances: element.MinimumOccurances,
			MaximumOccurances: element.MaximumOccurances,
		}
	}

	return &ebml, err

}

func readSpecification(path string) (ebmlStructure, error) {
	xmlFile, err := os.Open(path)

	if err != nil {
		return ebmlStructure{}, err
	}

	defer xmlFile.Close()

	rawValue, _ := ioutil.ReadAll(xmlFile)

	var ebml ebmlStructure

	err = xml.Unmarshal(rawValue, &ebml)

	if err != nil {
		return ebmlStructure{}, err
	}

	return ebml, nil
}
