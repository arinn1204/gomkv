package ebml

import (
	"encoding/binary"

	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/filesystem"
)

//Ebml will contain the IoReader as well as the current position of this members stream
type Ebml struct {
	File          filesystem.Reader
	CurrPos       int64
	Specification *specification.Ebml
}

func (ebml *Ebml) Copy() *Ebml {
	return &Ebml{
		File:          ebml.File,
		CurrPos:       ebml.CurrPos,
		Specification: ebml.Specification,
	}
}

//GetSize will return the size of the proceeding EBML element
func (ebml *Ebml) GetSize() (int64, error) {
	buf := make([]byte, 1)
	n, err := ebml.File.Read(ebml.CurrPos, buf)

	if err != nil {
		return 0, err
	}

	ebml.CurrPos += int64(n)

	seed := buf[0]
	width := getWidth(seed)

	size := int64(0)

	switch width {
	case 8:
		size, err = read(7, ebml, 0)
	case 7:
		size, err = read(6, ebml, seed)
	case 6:
		size, err = read(5, ebml, seed&3)
	case 5:
		size, err = read(4, ebml, seed&7)
	case 4:
		size, err = read(3, ebml, seed&15)
	case 3:
		size, err = read(2, ebml, seed&31)
	case 2:
		size, err = read(1, ebml, seed&63)
	case 1:
		size = int64(seed) & 127
	default:
		size = 0
	}

	return size, err
}

func read(count uint, ebml *Ebml, seed byte) (int64, error) {
	readBuf := make([]byte, count)
	n, err := ebml.File.Read(ebml.CurrPos, readBuf)

	if err != nil {
		return 0, err
	}

	ebml.CurrPos += int64(n)

	buf := make([]byte, 8)
	copy(buf[8-count:], readBuf)
	buf[8-count-1] = seed
	return int64(binary.BigEndian.Uint64(buf)), err
}

func getWidth(firstByte byte) uint {

	result := uint(0)
	first := byte(255)

	for first > 0 {
		if (firstByte | first) == first {
			result++
		}

		first >>= 1
	}

	return result
}
