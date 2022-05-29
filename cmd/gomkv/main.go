package main

import (
	"os"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/filesystem"
)

func main() {
	filePath := "../../data/Better.Call.Saul.S06.SPECIAL.American.Greed.James.McGill.1080p.AMZN.WEB-DL.DDP2.0.H.264-SKiZOiD.mkv"
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	wrapper := filesystem.File{
		File: file,
	}

	ebml := ebml.Ebml{
		File:    wrapper,
		CurrPos: 0,
	}

	doc := ebml.Read()
	_ = doc
}
