package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/HomeWork/Week2/NamikoToriyama/llist"
	"github.com/HomeWork/Week2/NamikoToriyama/stack"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var s string
	fmt.Printf("'LinkedList(ll)' or 'Stack(s)'> ")
	if sc.Scan() {
		s = sc.Text()
	}

	if s == "ll" {
		llist.CallLinkedList()
	}

	if s == "s" {
		stack.CallStack()
	}
}
