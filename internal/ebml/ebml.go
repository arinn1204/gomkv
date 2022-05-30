package ebml

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/filesystem"
	"github.com/arinn1204/gomkv/pkg/types"
)

//Ebml will contain the IoReader as well as the current position of this members stream
type Ebml struct {
	File              filesystem.Reader
	CurrPos           int64
	SpecificationPath string
}

type ebmlHeader struct {
	header types.Header
	err    error
}

type ebmlSegment struct {
	segment types.Segment
	err     error
}

func (ebml Ebml) Read() (types.EbmlDocument, error) {
	doc := types.EbmlDocument{}
	spec, err := specification.GetSpecification(ebml.SpecificationPath)
	if err != nil {
		return doc, err
	}

	err = validateMagicNum(&ebml, spec)
	if err != nil {
		return doc, err
	}

	headerChan := make(chan (ebmlHeader))
	segmentChan := make(chan (ebmlSegment))

	go func() {
		h, err := createHeader(ebml, spec)
		header := ebmlHeader{
			header: h,
			err:    err,
		}

		headerChan <- header
	}()

	go func() {
		segmentChan <- ebmlSegment{}
	}()

	ebmlHeader := <-headerChan
	segment := <-segmentChan

	doc.Header = ebmlHeader.header
	doc.Segments = append(doc.Segments, segment.segment)

	if ebmlHeader.err != nil {
		err = ebmlHeader.err
	}

	if segment.err != nil {
		if err == nil {
			err = segment.err
		} else {
			err = errors.New(err.Error() + segment.err.Error())
		}
	}

	return doc, err
}

func validateMagicNum(ebml *Ebml, spec specification.Ebml) error {
	idBuf := make([]byte, 4)
	n, err := ebml.File.Read(ebml.CurrPos, idBuf)

	if err != nil {
		return err
	}

	ebml.CurrPos += int64(n)

	id := binary.BigEndian.Uint32(idBuf)
	elem := spec.Data[id]

	if elem.Name != "EBML" {
		return fmt.Errorf("incorrect type of file expected magic number found %x", id)
	}

	return nil
}
