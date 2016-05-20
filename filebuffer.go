package main

import (
	"bytes"
)

type file struct {
	lines [][]byte
}

func NewFile() *file {
	return &file{
		make([][]byte, 0),
	}
}

func (f *file) append(line []byte) {
	// for _, l := range lines {
	f.lines = append(f.lines, line)
	// }
}

func (f *file) String() string {
	return string(bytes.Join(f.lines, []byte{'\n'}))
}
