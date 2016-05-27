package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "log"
)

const (
	commandRegex = `^(?:(\d+)?(?:,?(\d+|\${1})?))?((?i)[a-z]{1})$`
)

var (
	fileHandle = NewFile()
)

type lRange struct {
	min, max int
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	commandLoop(input)
}

func loopback(line []byte) bool {
	if len(line) < 1 {
		return true
	}
	return line[0] == '.' && len(line) == 1
}

func commandLoop(input *bufio.Scanner) {
	fmt.Printf(":")
	for input.Scan() {
		cmd := input.Bytes()
		subloop, e := parseCommand(cmd)
		switch {
		case e != nil:
			fmt.Println(e)
		case subloop != nil:
			subloop(cmd, input)
		default:
		}
		fmt.Printf(":")
	}
}

func parseCommand(line []byte) (Loop, error) {
	switch {
	case loopback(line):
		return nil, nil
	case line[0] == 'a':
		return appnd, nil
	case line[0] == 'p':
		return echo, nil
	case line[0] == 'w':
		return save, nil
	case line[0] == 'q':
		return quit, nil
	case line[0] == 'e':
		return nil, errors.New("you intentionally errored")
	default:
		return nil, errors.New("? unrecognized command")
	}
}

func parseRange(cmd []byte, fh *file) (*lRange, error) {
	var e error
	lRange := &lRange{}

	regex := regexp.MustCompile(commandRegex)
	matches := regex.FindAllStringSubmatch(string(cmd), -1)
	if matches == nil {
		return nil, errors.New("invalid command syntax")
	}

	if matches[0][1] == "" {
		lRange.min = 0
	} else {
		lRange.min, e = strconv.Atoi(matches[0][1])
		if e != nil {
			return nil, e
		}
	}

	if matches[0][2] == "" || matches[0][2] == "$" {
		lRange.max = fileHandle.Len()
	} else {
		lRange.max, e = strconv.Atoi(matches[0][2])
		if e != nil {
			return nil, e
		}
	}

	if lRange.max > fh.Len() {
		lRange.max = fh.Len()
	}

	if lRange.min > lRange.max {
		return nil, errors.New("range error: min was greater than max")
	}

	return lRange, nil

}
