package reader

import (
	"github.com/arinn1204/gomkv/internal/io"
)

type EbmlReader struct {
	File    *io.File
	currPos uint
}

type IEbmlReader interface {
	GetSize(width uint) uint64
	GetWidth() uint
}

var reader EbmlReader

func (reader EbmlReader) GetSize(width uint) uint64 {
	return uint64(0)
}

func (reader EbmlReader) GetWidth() uint {
	return uint(0)
}
