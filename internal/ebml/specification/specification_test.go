package specification

import (
	"encoding/xml"
	"io/fs"
	"strconv"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSerializeTopLevelEbmlSpec(t *testing.T) {
	specificationFile := "testdata/basicEbml.xml"
	spec, _ := GetSpecification(specificationFile)

	ebml := EbmlData{
		Name:              "EBMLMaxIDLength",
		Type:              "uinteger",
		Range:             "4",
		Default:           "4",
		MinimumOccurances: 1,
		MaximumOccurances: 1,
	}

	id, _ := strconv.ParseInt("42F2", 16, 16)

	data := make(map[uint32]EbmlData)
	data[uint32(id)] = ebml

	ebmlStructure := Ebml{
		Data:         data,
		Version:      4,
		DocumentType: "matroska",
	}

	assert.Equal(t, ebmlStructure, *spec)
}

func TestWillPanicIfFileNotFound(t *testing.T) {
	specificationFile := "testdata/notFound.xml"

	_, err := GetSpecification(specificationFile)

	expected := &fs.PathError{
		Op:   "open",
		Path: specificationFile,
		Err:  error(syscall.Errno(syscall.ENOENT)),
	}

	assert.Equal(
		t,
		expected,
		err,
	)
}

func TestWillPanicIfBadlyFormattedXml(t *testing.T) {
	specificationFile := "testdata/badXml.xml"

	_, err := GetSpecification(specificationFile)
	expected := &xml.SyntaxError{
		Msg:  "unexpected EOF",
		Line: 4,
	}

	assert.Equal(
		t,
		expected,
		err,
	)
}
