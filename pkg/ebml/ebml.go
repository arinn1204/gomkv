package ebml

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/mapper"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/filesystem"
	"github.com/arinn1204/gomkv/pkg/types"
)

type ebmlObj[T any] struct {
	data T
	err  error
}

func Read(file *filesystem.File, specPath string) (types.EbmlDocument, error) {

	ebml := ebml.Ebml{
		File:              file,
		CurrPos:           0,
		SpecificationPath: specPath,
	}

	doc := types.EbmlDocument{}
	spec, err := specification.GetSpecification(ebml.SpecificationPath)
	if err != nil {
		return doc, err
	}

	err = validateMagicNum(&ebml, spec)
	if err != nil {
		return doc, err
	}

	headerChan := make(chan (ebmlObj[types.Header]))
	segmentChan := make(chan (ebmlObj[[]types.Segment]))

	go func() {
		h, err := mapper.Header{}.Map(ebml, &spec)
		header := ebmlObj[types.Header]{
			data: h,
			err:  err,
		}

		headerChan <- header
	}()

	go func() {
		segments, err := mapper.Segment{}.Map(ebml, &spec)

		data := ebmlObj[[]types.Segment]{
			data: segments,
			err:  err,
		}

		segmentChan <- data
	}()

	ebmlHeader := <-headerChan
	segment := <-segmentChan

	doc.Header = ebmlHeader.data
	doc.Segments = segment.data

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

func validateMagicNum(ebml *ebml.Ebml, spec specification.Ebml) error {
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
