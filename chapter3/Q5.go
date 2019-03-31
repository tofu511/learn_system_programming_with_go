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
	CopyN(os.Stdout, file, 64)
}

func CopyN(dest io.Writer, src io.Reader, length int) (int64, error) {
	writeSize, err := io.Copy(dest, io.LimitReader(src, int64(length)))
	return writeSize, err
}
