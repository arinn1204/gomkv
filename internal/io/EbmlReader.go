package io

type EbmlReader interface {
	Read(startPos uint, buf []byte) int
}
