package main

import (
	"os"
)

func main() {
	filePath := "../../data/Better.Call.Saul.S06.SPECIAL.American.Greed.James.McGill.1080p.AMZN.WEB-DL.DDP2.0.H.264-SKiZOiD.mkv"
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}
	defer file.Close()
}
