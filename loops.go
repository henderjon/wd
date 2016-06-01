package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func echo(cmd []byte, input *bufio.Scanner) {

	_, e := parseRange(cmd, fileHandle)
	if e != nil {
		fmt.Println(e)
		return
	}

	for k, v := range fileHandle.Lines() {
		// if k >= lr.min && k <= lr.max {
			if k == currentLine {
				fmt.Printf("%6d > %s\n", k, v)
			}else{
				fmt.Printf("%6d | %s\n", k, v)
			}
		// }
	}
	log.Println(currentLine)
}

func insert(cmd []byte, input *bufio.Scanner) {
	for input.Scan() {

		line := input.Bytes()

		if loopback(line) {
			return
		}

		fileHandle.append(currentLine, line)
		currentLine += 1
	}
}

func add(cmd []byte, input *bufio.Scanner) {
	for input.Scan() {

		line := input.Bytes()

		if loopback(line) {
			log.Println(fileHandle.lines)
			return
		}

		currentLine += 1
		fileHandle.append(currentLine, line)
	}
}

func quit(cmd []byte, input *bufio.Scanner) {
	os.Exit(0)
}

func save(cmd []byte, input *bufio.Scanner) {
	fileHandle.Save()
}

func setLine(cmd []byte, input *bufio.Scanner) {
	lr, e := parseRange(cmd, fileHandle)
	if e != nil {
		fmt.Println(e)
		return
	}

	currentLine = lr.min
}

// if lr, e := parseRange(cmd, fileHandle); e == nil {
// 	log.Printf("range: [%d, %d]\n", lr.min, lr.max)
// }else{
// 	log.Println(e)
// }

// if c, ok := parseCommand(cmd); ok {
// 	log.Println("command:", c)
// }else{
// 	log.Println("trouble parsing command")
// }

type Loop func([]byte, *bufio.Scanner)
