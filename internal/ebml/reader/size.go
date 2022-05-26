package reader

//GetSize will return the size of the proceeding EBML element
func (ebmlReader EbmlReader) GetSize(width int) int64 {
	return int64(0)
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
