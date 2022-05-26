package io

type EbmlReader interface {
	Read(f *File, startPos uint, buf []byte) int
}
