package io

import "os"

//File is the structure that will be used to mock the file pointer
//Eventually this will be used to have a mutex to control multiple threads
type File struct {
	file *os.File
}

func Read(f *File, startPos uint, buf []byte) int {
	file := f.file

	file.Seek(int64(startPos), 0)

	ret, _ := file.Read(buf)

	return ret
}
