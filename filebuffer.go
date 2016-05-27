package main

import (
	"bytes"
	"io/ioutil"
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

// func (f *file) String() string {
// 	return string(bytes.Join(f.lines, []byte{'\n'}))
// }

// func (f *file) Bytes() [][]byte {
// 	return f.lines
// }

func (f *file) Lines() [][]byte {
	return f.lines
}

func (f *file) Len() int {
	return len(f.lines)
}

func (f *file) Save() error {
	return ioutil.WriteFile("tmp2.txt", bytes.Join(f.lines, []byte{'\n'}), 0644)
}
