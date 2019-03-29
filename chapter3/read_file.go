package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
