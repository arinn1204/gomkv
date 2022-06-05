package mapper

import "github.com/arinn1204/gomkv/internal/ebml"

type Mapper[T any] interface {
	Map(size int64, ebml ebml.Ebml) (*T, error)
}

var GetID func(ebml *ebml.Ebml, maxCount int) (uint32, error)
var read func(ebml *ebml.Ebml, data []byte) (int, error)
var getSize func(ebml *ebml.Ebml) (int64, error)

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
}
