package main

import (
	"bufio"
	"fmt"
	"os"
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

func quit(input *bufio.Scanner) {
	os.Exit(0)
}

type Loop func(*bufio.Scanner)
