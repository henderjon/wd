package main

import (
	"bufio"
	"fmt"
	"os"
)

var fh = NewFile()

func main() {
	r := bufio.NewScanner(os.Stdin)
	for r.Scan() {
		parseCommand(r.Bytes())
		fmt.Println("```")
		fmt.Println(fh.String())
		fmt.Println("```")
	}
}

func parseCommand(line []byte) {
	switch line[0] {
	case 'a':
		fmt.Println("a")
		fh.append(line)
		// case 'b' :
		// 	fmt.Println("b")
		// case 'c' :
		// 	fmt.Println("c")
	}
}
