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

	if err = validateMagicNum(&ebml, &spec); err != nil {
		return doc, err
	}

	return buildDoc(&ebml, &spec)
}

func buildDoc(ebml *ebml.Ebml, spec *specification.Ebml) (types.EbmlDocument, error) {
	headerChan := createItem[types.Header](mapper.Header{}, *ebml, spec)
	segmentsChan := createItem[[]types.Segment](mapper.Segment{}, *ebml, spec)

	header := <-headerChan
	segments := <-segmentsChan

	var err error

	if header.err != nil {
		err = header.err
	}

	if segments.err != nil {
		if err == nil {
			err = segments.err
		} else {
			err = errors.New(err.Error() + segments.err.Error())
		}
	}

	return types.EbmlDocument{
		Header:   header.data,
		Segments: segments.data,
	}, err
}

func createItem[T any](mapper mapper.Mapper[T], ebml ebml.Ebml, spec *specification.Ebml) <-chan ebmlObj[T] {
	channel := make(chan ebmlObj[T])

	go func() {
		data, err := mapper.Map(ebml, spec)
		obj := ebmlObj[T]{
			data: data,
			err:  err,
		}
		channel <- obj
	}()

	return channel
}

func validateMagicNum(ebml *ebml.Ebml, spec *specification.Ebml) error {
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
