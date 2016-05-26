package main

import (
	"bufio"
	"fmt"
)

func echo(input *bufio.Scanner) {
	fmt.Println(fileHandle.String())
}

func app(input *bufio.Scanner) {
	for input.Scan() {

		line := input.Bytes()

		if shouldCommand(line) {
			return
		}

		fileHandle.append(line)
	}
}

type Loop func(*bufio.Scanner)
