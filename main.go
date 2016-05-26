package main

import (
	"bufio"
	"fmt"
	"os"
)

var fh = NewFile()

func main() {
	r := bufio.NewScanner(os.Stdin)
	prompt()
	for r.Scan() {
		parseCommand(r.Bytes())
		fmt.Println("```")
		fmt.Println(fh.String())
		fmt.Println("```")
		prompt()
	}
}

func prompt(){
	fmt.Printf(":")
}

func parseCommand(line []byte) {
	if len(line) == 1 && line[0] == '.' {
		prompt()
		return
	}
	switch line[0] {
	case 'a':
		fmt.Println("a")
		fh.append(line[1:])
	case 'r':
		fmt.Println("r")
		fh.replace(2, line[1:])
		// case 'b' :
		// 	fmt.Println("b")
		// case 'c' :
		// 	fmt.Println("c")
	}
}

func readInput(r bufio.Scanner){
	for r.Scan() {
		return
	}
}
