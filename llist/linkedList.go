package llist

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

// CallLinkedList ... execute linked list
func CallLinkedList() {

	var s string
	list := CreateList()

	for {
		fmt.Printf("'add(a)' or 'remove(r)' or 'reverse' or 'quit(q)' > ")
		if sc.Scan() {
			s = sc.Text()
		}
		if s == "quit" || s == "q" {
			break
		}

		if s == "add" || s == "a" {
			index := InputIndex()
			value := InputValue()
			list = InsertNode(value, index, list)
		}

		if s == "remove" || s == "r" {
			index := InputIndex()
			list = RemoveNode(index, list)
		}

		if s == "reverse" || s == "rv" {
			list = ReverseList(list)
		}
		PrintList(list)
	}
}
