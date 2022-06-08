package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
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
		return ebml.GetID(maxCount)
	}

	read = func(ebml *ebml.Ebml, data []byte) (int, error) {
		return ebml.File.Read(ebml.CurrPos, data)
	}

	getSize = func(ebml *ebml.Ebml) (int64, error) {
		return ebml.GetSize()
	}

	readUntil = func(ebml *ebml.Ebml, end int64, process func(id uint32, endPosition int64, element *specification.EbmlData) error) error {
		return ebml.ReadUntilElementFound(end, process)
	}
}
