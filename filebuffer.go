package main

import (
	"bytes"
	"io/ioutil"
	// "log"
)

type file struct {
	lines [][]byte
}

func NewFile() *file {
	return &file{
		make([][]byte, 0),
	}
}

func (f *file) append(lineNumber int, line []byte) {

	var pre, post [][]byte

	lineNumber -= 1

	if  len(f.lines) >= lineNumber {
		pre  = f.lines[:lineNumber]
		post = f.lines[lineNumber:]
	}

	f.lines = append(pre, line)
	f.lines = append(f.lines, post...)
}

func (f *file) replace(index int, line []byte) {
	switch {
	case len(f.lines) < index:
		return
	case len(f.lines) == index:
		f.append(index, line)
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
