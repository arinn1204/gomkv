package specification

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSerializeTopLevelEbmlSpec(t *testing.T) {
	specificationFile := "testdata/basicEbml.xml"
	spec := GetSpecification(specificationFile)

	ebml := EbmlData{
		Name:              "EBMLMaxIDLength",
		Type:              "uinteger",
		Range:             "4",
		Default:           4,
		MinimumOccurances: 1,
		MaximumOccurances: 1,
	}

	id, _ := strconv.ParseInt("0x42F2", 16, 16)

	data := make(map[uint32]EbmlData)
	data[uint32(id)] = ebml

	ebmlStructure := Ebml{
		Data:         data,
		Version:      4,
		DocumentType: "matroska",
	}

	assert.Equal(t, ebmlStructure, spec)
}

func TestWillPanicIfFileNotFound(t *testing.T) {
	specificationFile := "testdata/notFound.xml"

	defer func() {
		err := recover().(string)
		if err == "" {
			assert.Fail(t, "Expected panic when %v was not found.", specificationFile)
		}

		pattern := "Failed to open specification. -- .*"
		match, _ := regexp.Match(pattern, []byte(err))

		assert.True(t, match, "Failed to match the pattern -- '%v'. Found '%v", pattern, err)
	}()

	GetSpecification(specificationFile)
}

func TestWillPanicIfBadlyFormattedXml(t *testing.T) {
	specificationFile := "testdata/badXml.xml"

	defer func() {
		err := recover().(string)
		if err == "" {
			assert.Fail(t, "Expected panic but not found.", specificationFile)
		}

		pattern := "Failed to parse the specification xml. -- .*"
		match, _ := regexp.Match(pattern, []byte(err))

		assert.True(t, match, "Failed to match the pattern -- '%v'. But found %v", pattern, err)

	}()

	GetSpecification(specificationFile)
}
