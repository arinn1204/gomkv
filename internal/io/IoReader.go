package io

type IoReader interface {
	Read(startPos uint, buf []byte) int
}
