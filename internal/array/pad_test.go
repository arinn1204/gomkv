package array

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWillPadArray(t *testing.T) {
	dest := make([]byte, 8)
	src := []byte{
		8,
		1,
		3,
		5,
		3,
	}

	err := Pad(src, dest)

	expected := []byte{
		0,
		0,
		0,
		8,
		1,
		3,
		5,
		3,
	}

	assert.Equal(t, expected, dest)
	assert.Nil(t, err)
}

func TestWillReturnErrorIfSourceIsLargerThanDestination(t *testing.T) {
	dest := make([]byte, 5)
	src := make([]byte, 8)

	err := Pad(src, dest)

	assert.Equal(t, err, errors.New("destination buffer not large enough - destination buffer must be greater than 8"))
}
