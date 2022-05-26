package io

//IoReader is a wrapper around os.Read()
//IoReader has the benefit of having the File structure associated with it
type IoReader interface {
	Read(startPos uint, buf []byte) int
}
