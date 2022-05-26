package reader

import (
	"github.com/arinn1204/gomkv/internal/io"
)

//EbmlReader will contain the IoReader as well as the current position of this members stream
type EbmlReader struct {
	Reader  io.IoReader
	CurrPos uint
}

var ebmlReader EbmlReader

//GetSize will return the size of the next element in the stream (in bytes)
func (reader EbmlReader) GetSize(width uint) uint64 {
	return uint64(0)
}

//GetWidth will return the width of the next element in the stream
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
