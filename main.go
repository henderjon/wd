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
	commandRegex = `^(?:(\d+)?(?:(,)?(\d+|\${1})?))?((?i)[a-z]{1})$`
)

var (
	fileHandle  = NewFile()
	currentLine = 0
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
		// return true
	}
	return len(line) == 1 && line[0] == '.'
}

func commandLoop(input *bufio.Scanner) {
	fmt.Printf(":")
	for input.Scan() {
		cmd := input.Bytes()
		subloop, e := parseFullCommand(cmd)
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

func parseFullCommand(line []byte) (Loop, error) {

	if loopback(line) {
		return nil, nil
	}

	c, ok := parseCommand(line)

	switch {
	case !ok:
		return nil, nil
	case c == `i`:
		return insert, nil
	case c == `a`:
		return add, nil
	case c == `p`:
		return echo, nil
	case c == `w`:
		return save, nil
	case c == `l`:
		return setLine, nil
	case c == `q`:
		return quit, nil
	case c == `e`:
		return nil, errors.New("? you intentionally errored")
	default:
		return nil, errors.New("? unrecognized command")
	}
}

func parseCommand(cmd []byte) (string, bool) {
	regex := regexp.MustCompile(commandRegex)
	matches := regex.FindAllStringSubmatch(string(cmd), -1)
	if matches == nil {
		return "", false
	}
	return matches[0][4], true
}

func parseRange(cmd []byte, fh *file) (*lRange, error) {
	var e error
	lRange := &lRange{}

	regex := regexp.MustCompile(commandRegex)
	matches := regex.FindAllStringSubmatch(string(cmd), -1)
	if matches == nil {
		return nil, errors.New("? invalid command syntax")
	}

	if matches[0][1] == "" {
		lRange.min = 0
	} else {
		lRange.min, e = strconv.Atoi(matches[0][1])
		if e != nil {
			return nil, e
		}
	}

	switch {
	case matches[0][2] == "": // was there a comma provided?
		lRange.max = lRange.min
	case matches[0][3] == "$":
		lRange.max = fh.Len()
	case matches[0][3] == "":
		lRange.max = lRange.min
	default:
		lRange.max, e = strconv.Atoi(matches[0][3])
		if e != nil {
			return nil, e
		}
	}

	if lRange.max > fh.Len() {
		lRange.max = fh.Len()
	}

	if lRange.min > lRange.max {
		return nil, errors.New("? range error: min was greater than max")
	}

	// currentLine = lRange.max
	return lRange, nil

}
