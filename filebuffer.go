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
	f.lines = append(f.lines, line)
}

func (f *file) replace(index int, line []byte) {
	switch {
	case len(f.lines) < index:
		return
	case len(f.lines) == index:
		f.append(line)
	case len(f.lines) > index:
		f.lines[index] = line
	}
}

func (f *file) String() string {
	return string(bytes.Join(f.lines, []byte{'\n'}))
}
