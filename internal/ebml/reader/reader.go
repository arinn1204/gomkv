package reader

import (
	"github.com/arinn1204/gomkv/internal/io"
)

type EbmlReader struct {
	Reader  io.IoReader
	CurrPos uint
}

type IEbmlReader interface {
	GetSize(width uint) uint64
	GetWidth() uint
}

var ebmlReader EbmlReader

func (reader EbmlReader) GetSize(width uint) uint64 {
	return uint64(0)
}

func (ebmlReader EbmlReader) GetWidth() uint {
	firstByte := make([]byte, 1)

	ret := ebmlReader.Reader.Read(0, firstByte)

	if ret == 0 || len(firstByte) == 0 {
		return 0
	}

	result := uint(0)
	first := byte(255)

	for first > 0 {
		if (firstByte[0] | first) == first {
			result++
		}

		first >>= 1
	}

	return result
}
