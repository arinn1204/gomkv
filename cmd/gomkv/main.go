package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/arinn1204/gomkv/internal/filesystem"
	"github.com/arinn1204/gomkv/pkg/ebml"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "path", "", "the path of the matroska file to parse")
	flag.Parse()
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	wrapper := filesystem.File{
		File: file,
	}

	doc, err := ebml.Read(&wrapper, "../../data/matroska_ebml.xml")
	jsonDoc, _ := json.MarshalIndent(doc, "", "  ")
	fmt.Println(string(jsonDoc))

	fmt.Println(err)
}
