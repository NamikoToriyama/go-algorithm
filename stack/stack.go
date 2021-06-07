package stack

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

// CallStack ... execute stack
func CallStack() {
	var s string
	list := CreateStack()

	for {
		fmt.Printf("'push' or 'pop' or 'quit(q)'> ")
		if sc.Scan() {
			s = sc.Text()
		}

		if s == "quit" || s == "q" {
			break
		}
		if s == "push" {
			value := InputValue()
			list = Push(value, list)
		}
		if s == "pop" {
			list = Pop(list)
		}
		PrintStack(list)
	}
}
