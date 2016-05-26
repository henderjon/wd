package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var (
	fileHandle = NewFile()
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	commandLoop(input)
}

func shouldCommand(line []byte) bool {
	return line[0] == '.' && len(line) == 1
}

func commandLoop(input *bufio.Scanner) {
	fmt.Printf(":")
	for input.Scan() {
		subloop, e := parseCommand(input.Bytes())
		switch {
		case e != nil:
			fmt.Println(e)
		case subloop != nil:
			subloop(input)
		default:
		}
		fmt.Printf(":")
	}
}

func parseCommand(line []byte) (Loop, error) {
	switch {
	case shouldCommand(line):
		fallthrough
	default:
		// return to the command loop
		return nil, nil
	case line[0] == 'a':
		return app, nil
	case line[0] == 'p':
		return echo, nil
	case line[0] == 'q':
		return quit, nil
	case line[0] == 'e':
		return nil, errors.New("you intentionally errored")
	}
}
