package ebml

//GetSize will return the size of the proceeding EBML element
func (ebml *Ebml) GetSize() int64 {
	buf := ebml.File.Read(ebml.CurrPos, 1)
	ebml.CurrPos += uint(len(buf))

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
		size = read(0, ebml, seed&127)
	default:
		size = 0
	}

	return size
}

func read(count uint, ebml *Ebml, seed byte) int64 {
	result := int64(seed) << (count * 8)

	if count == 0 {
		return result
	}

	buf := ebml.File.Read(ebml.CurrPos, count)
	ebml.CurrPos += uint(len(buf))

	for i := uint(0); i < count; i++ {
		result += int64(buf[i]) << ((count - 1 - i) * 8)
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
