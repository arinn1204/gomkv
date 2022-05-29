package array

import (
	"fmt"
)

//Pad will take two input arrays, the source and destination
//it will take the source array and place it into the destination array, padding the beginning with zeros
func Pad(src []byte, dest []byte) error {
	if len(src) > len(dest) {
		return fmt.Errorf("destination buffer not large enough - destination buffer must be greater than %v", len(src))
	}

	copy(dest[len(dest)-len(src):], src)
	return nil
}
