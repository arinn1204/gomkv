package reader

//GetSize will return the size of the proceeding EBML element
func (ebmlReader *EbmlReader) GetSize() int64 {
	buf := make([]byte, 1)
	ret := ebmlReader.Reader.Read(ebmlReader.CurrPos, buf)
	ebmlReader.CurrPos += uint(ret)

	seed := buf[0]
	width := getWidth(seed)

	size := int64(0)

	switch width {
	case 8:
		size = read(7, ebmlReader, 0)
	case 7:
		size = read(6, ebmlReader, seed)
	case 6:
		size = read(5, ebmlReader, seed&3)
	case 5:
		size = read(4, ebmlReader, seed&7)
	case 4:
		size = read(3, ebmlReader, seed&15)
	case 3:
		size = read(2, ebmlReader, seed&31)
	case 2:
		size = read(1, ebmlReader, seed&63)
	case 1:
		size = read(0, ebmlReader, seed&127)
	default:
		size = 0
	}

	return size
}

func read(count int, reader *EbmlReader, seed byte) int64 {
	result := int64(seed) << (count * 8)

	if count == 0 {
		return result
	}

	file := reader.Reader

	buf := make([]byte, count)

	file.Read(reader.CurrPos, buf)
	reader.CurrPos += uint(len(buf))

	for i := count - 1; i > 0; i-- {
		result |= int64(buf[i]) << ((i - 1) * 8)
	}

	return result
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
