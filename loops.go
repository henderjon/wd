package main

import (
	"bufio"
	"fmt"
	"os"
)

func echo(cmd []byte, input *bufio.Scanner) {

	// lr, e := parseRange(cmd, fileHandle)
	for k, v := range fileHandle.Lines() {
		fmt.Printf("%6d\t%s\n", k, v)
	}
}

func appnd(cmd []byte, input *bufio.Scanner) {
	for input.Scan() {

		line := input.Bytes()

		if loopback(line) {
			return
		}

		fileHandle.append(line)
	}
}

func quit(cmd []byte, input *bufio.Scanner) {
	os.Exit(0)
}

func save(cmd []byte, input *bufio.Scanner) {
	fileHandle.Save()
}

type Loop func([]byte, *bufio.Scanner)
