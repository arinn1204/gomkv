package ebml

import (
	"errors"
	"fmt"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/mapper"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/filesystem"
	"github.com/arinn1204/gomkv/pkg/types"
)

type ebmlObj[T any] struct {
	data *T
	err  error
}

func Read(file *filesystem.File, specPath string) ([]types.EbmlDocument, error) {
	spec, err := specification.GetSpecification(specPath)
	if err != nil {
		return nil, err
	}
	ebml := ebml.Ebml{
		File:          file,
		CurrPos:       0,
		Specification: spec,
	}

	return buildDoc(&ebml)
}

func buildDoc(ebml *ebml.Ebml) ([]types.EbmlDocument, error) {
	var err error
	var id uint32
	channel := make(chan ebmlObj[types.EbmlDocument])

	for {
		id, err = mapper.GetID(ebml, 4)

		if err != nil {
			break
		}

		if id == 0x1A45DFA3 { //EBML Header
			go readDocument(channel, *ebml)
			//Incrememt to the next document if within a stream
			if err = incDoc(ebml); err != nil {
				break
			}
		} else { //Unknown top level element
			err = fmt.Errorf("unknown top level element (%x)", id)
			break
		}
	}

	docs := make([]types.EbmlDocument, 0)
	for c := range channel {
		docs = append(docs, types.EbmlDocument{
			Header:   c.data.Header,
			Segments: c.data.Segments,
		})

		if err == nil {
			err = c.err
		} else {
			err = errors.New(err.Error() + c.err.Error())
		}
	}

	return docs, err
}

func incDoc(ebml *ebml.Ebml) error {
	size, err := ebml.GetSize()
	if err != nil {
		return err
	}
	ebml.CurrPos += size

	id, err := mapper.GetID(ebml, 4)

	//EBML documents allow for multiple segments
	//so we will skip all the segments to allow the documents to be read in parallel
	for id == 0x18538067 && err == nil {
		size, _ = ebml.GetSize()
		ebml.CurrPos += size
		id, err = mapper.GetID(ebml, 4)
	}
	return err
}

func readDocument(channel chan<- ebmlObj[types.EbmlDocument], ebml ebml.Ebml) {
	size, err := ebml.GetSize()

	if err != nil {
		channel <- ebmlObj[types.EbmlDocument]{
			data: nil,
			err:  err,
		}
		close(channel)
		return
	}

	header, err := mapper.Header{}.Map(size, ebml)
	ebml.CurrPos += size

	if err != nil {
		channel <- ebmlObj[types.EbmlDocument]{
			data: &types.EbmlDocument{
				Header:   header,
				Segments: nil,
			},
			err: err,
		}
		close(channel)
		return
	}

	id, _ := mapper.GetID(&ebml, 4)
	size, err = ebml.GetSize()
	seg := make(chan ebmlObj[types.Segment], 5)
	numSegments := 0
	for id == 0x18538067 && err == nil {
		go readSegment(ebml, seg, size)
		ebml.CurrPos += size
		id, _ = mapper.GetID(&ebml, 4)
		size, err = ebml.GetSize()
		numSegments++
	}

	segments := make([]types.Segment, 0)
	for numSegments != 0 {
		segmentData := <-seg
		segments = append(segments, *segmentData.data)
		numSegments--

		if segmentData.err != nil {
			if err == nil {
				err = segmentData.err
			} else {
				err = errors.New(err.Error() + segmentData.err.Error())
			}
		}
	}

	close(seg)

	channel <- ebmlObj[types.EbmlDocument]{
		data: &types.EbmlDocument{
			Header:   header,
			Segments: segments,
		},
		err: err,
	}

	if err != nil {
		close(channel)
	}
}

func readSegment(ebml ebml.Ebml, seg chan<- ebmlObj[types.Segment], size int64) {
	segment, err := mapper.Segment{}.Map(size, ebml)
	seg <- ebmlObj[types.Segment]{
		data: segment,
		err:  err,
	}
	ebml.CurrPos += size

	if err != nil {
		close(seg)
	}
}
