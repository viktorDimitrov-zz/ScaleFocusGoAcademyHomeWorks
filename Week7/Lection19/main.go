package main

import (
	"io"
	"os"
)

type ReverseStringReader struct {
	i int // current reading index
	s []byte
}

func (r *ReverseStringReader) Read(b []byte) (n int, err error) {

	if r.i < 0 {
		return 0, io.EOF
	}
	n = copy(b, []byte{r.s[r.i]})
	r.i--
	return n, nil
}

func NewReverseStringReader(input string) *ReverseStringReader {
	rsr := ReverseStringReader{i: len(input) - 1, s: []byte(input)}
	return &rsr
}

func main() {
	ourReverseStringReader := NewReverseStringReader("Reverse this String")
	io.Copy(os.Stdout, ourReverseStringReader)

}
