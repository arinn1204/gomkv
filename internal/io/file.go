package io

import "os"

//File is the structure that will be used to mock the file pointer
//Eventually this will be used to have a mutex to control multiple threads
type File struct {
	File *os.File
}

type EbmlReader interface {
	Read(f *File, startPos uint, buf []byte) int
}

var ebmlFile File

func init() {
	ebmlFile = File{}
}

//Read is a wrapper around os.File.Read
func (ebmlFile File) Read(startPos uint, buf []byte) int {
	file := ebmlFile.File

	file.Seek(int64(startPos), 0)

	ret, _ := file.Read(buf)

	return ret
}
