package main

import (
	"io"
	"os"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	for i, _ := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	var a MyReader
	reader.Validate(a)
	io.Copy(os.Stdout, a)
}
