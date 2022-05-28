package io

import (
	"os"
	"sync"
)

//File is the structure that will be used to mock the file pointer
//Eventually this will be used to have a mutex to control multiple threads
type File struct {
	File *os.File
}

var ebmlFile File
var mutex sync.Mutex

func init() {
	mutex = sync.Mutex{}
}

//Read is a wrapper around os.File.Read
func (ebmlFile *File) Read(startPos uint, buf []byte) int {
	file := ebmlFile.File

	mutex.Lock()
	defer mutex.Unlock()
	file.Seek(int64(startPos), 0)
	n, _ := file.Read(buf)

	return n
}
