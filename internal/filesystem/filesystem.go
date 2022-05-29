package filesystem

import (
	"io"
	"os"
	"sync"
)

//File is the structure that will be used to mock the file pointer
//Eventually this will be used to have a mutex to control multiple threads
type File struct {
	File *os.File
}

//Reader is a wrapper around os.Read()
//Reader has the benefit of having the File structure associated with it
type Reader interface {
	Read(startPos uint, buf []byte) (int, error)
}

var mutex sync.Mutex

func init() {
	mutex = sync.Mutex{}
}

//Read is a wrapper around os.File.Read
func (ebmlFile File) Read(startPos uint, buf []byte) (int, error) {
	file := ebmlFile.File

	mutex.Lock()
	defer mutex.Unlock()
	_, err := file.Seek(int64(startPos), 0)

	if err != nil {
		return handleErr(err)
	}

	n, err := file.Read(buf)

	if err != nil {
		return handleErr(err)
	}

	return n, nil
}

func handleErr(err error) (int, error) {
	if err == io.EOF {
		return 0, err
	}

	panic(err)

}
