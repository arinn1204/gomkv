package main

import (
	"os"

	"github.com/arinn1204/gomkv/internal/filesystem"
	"github.com/arinn1204/gomkv/pkg/ebml"
)

func main() {
	betterCallSaul := "../../data/Better.Call.Saul.S06.SPECIAL.American.Greed.James.McGill.1080p.AMZN.WEB-DL.DDP2.0.H.264-SKiZOiD.mkv"
	//smallMkv := "/Users/arinn/Projects/VideoCatalogue/source/test/Grains.Tests.Integration/TestData/CodecParser/small.mkv"
	file, err := os.Open(betterCallSaul)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	wrapper := filesystem.File{
		File: file,
	}

	doc, _ := ebml.Read(&wrapper, "../../data/matroska_ebml.xml")
	_ = doc
}
