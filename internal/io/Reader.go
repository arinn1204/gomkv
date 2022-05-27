package io

//Reader is a wrapper around os.Read()
//Reader has the benefit of having the File structure associated with it
type Reader interface {
	Read(startPos uint, size uint) []byte
}
