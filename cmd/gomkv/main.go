package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/arinn1204/gomkv/internal/filesystem"
	"github.com/arinn1204/gomkv/pkg/ebml"
)

func main() {
	var filePath string
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	wrapper := filesystem.File{
		File: file,
	}

	doc, _ := ebml.Read(&wrapper, "../../data/matroska_ebml.xml")
	jsonDoc, _ := json.MarshalIndent(doc, "", "  ")
	fmt.Println(string(jsonDoc))
}
