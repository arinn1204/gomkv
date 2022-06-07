package mapper

import (
	"fmt"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/utils"
)

type Mapper[T any] interface {
	Map(size int64, ebml ebml.Ebml) (*T, error)
}

var GetID func(ebml *ebml.Ebml, maxCount int) (uint32, error)
var read func(ebml *ebml.Ebml, data []byte) (int, error)
var getSize func(ebml *ebml.Ebml) (int64, error)
var readUntil func(ebml *ebml.Ebml, end int64, process func(id uint32, endPosition int64, element *specification.EbmlData) error) error

func init() {
	GetID = func(ebml *ebml.Ebml, maxCount int) (uint32, error) {
		return getID(ebml, maxCount)
	}

	read = func(ebml *ebml.Ebml, data []byte) (int, error) {
		return ebml.File.Read(ebml.CurrPos, data)
	}

	getSize = func(ebml *ebml.Ebml) (int64, error) {
		return ebml.GetSize()
	}

	readUntil = readUntilElementFound
}

func readUntilElementFound(
	ebml *ebml.Ebml,
	end int64,
	process func(id uint32, endPosition int64, element *specification.EbmlData) error) error {
	var err error
	for ebml.CurrPos < end {
		id, idErr := GetID(ebml, 4)

		if idErr != nil {
			err = utils.ConcatErr(err, idErr)
			break
		}

		size, sizeErr := getSize(ebml)

		if sizeErr != nil {
			err = utils.ConcatErr(err, sizeErr)
			break
		}
		err = utils.ConcatErr(err, utils.ConcatErr(idErr, sizeErr))

		element := ebml.Specification.Data[id]

		if element == nil {
			err = utils.ConcatErr(err, fmt.Errorf("unknown element of id 0x%X", id))
			ebml.CurrPos += size
			continue
		}

		err = utils.ConcatErr(err, process(id, ebml.CurrPos+size, element))
	}

	return err
}
