package io

import "os"

//File is the structure that will be used to mock the file pointer
//Eventually this will be used to have a mutex to control multiple threads
type File struct {
	File *os.File
}

var ebmlFile File

//Read is a wrapper around os.File.Read
func (ebmlFile *File) Read(startPos uint, count uint) []byte {
	file := ebmlFile.File

	file.Seek(int64(startPos), 0)
	buf := make([]byte, count)

	file.Read(buf)

	return buf
}
