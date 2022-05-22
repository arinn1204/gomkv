package main

import (
	"flag"
	"fmt"
)

func main() {
	pathPtr := flag.String("path", " ", "The path of the file to read")
	flag.Parse()

	fmt.Printf("Path to read: '%v'\n", *pathPtr)
}
