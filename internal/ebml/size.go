package ebml

import "encoding/binary"

//GetSize will return the size of the proceeding EBML element
func (ebml *Ebml) GetSize() int64 {
	buf := make([]byte, 1)
	n, _ := ebml.File.Read(ebml.CurrPos, buf)
	ebml.CurrPos += uint(n)

	seed := buf[0]
	width := getWidth(seed)

	size := int64(0)

	switch width {
	case 8:
		size = read(7, ebml, 0)
	case 7:
		size = read(6, ebml, seed)
	case 6:
		size = read(5, ebml, seed&3)
	case 5:
		size = read(4, ebml, seed&7)
	case 4:
		size = read(3, ebml, seed&15)
	case 3:
		size = read(2, ebml, seed&31)
	case 2:
		size = read(1, ebml, seed&63)
	case 1:
		size = int64(seed) & 127
	default:
		size = 0
	}

	return size
}

func read(count uint, ebml *Ebml, seed byte) int64 {
	readBuf := make([]byte, count)
	n, _ := ebml.File.Read(ebml.CurrPos, readBuf)
	ebml.CurrPos += uint(n)

	buf := make([]byte, 8)
	copy(buf[8-count:], readBuf)
	buf[8-count-1] = seed
	return int64(binary.BigEndian.Uint64(buf))
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
