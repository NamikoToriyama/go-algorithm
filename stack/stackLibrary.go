package stack

import (
	"fmt"
	"strconv"
)

// Node ... element of linked list
type Node struct {
	value int
	next  *Node
}

// InputValue ... return the value entered by the user
func InputValue() int {
	var value int
	fmt.Printf("value > ")
	if sc.Scan() {
		t := sc.Text()
		value, _ = strconv.Atoi(t)
	}
	return value
}

// CreateStack ...create linked list
func CreateStack() *Node {
	fmt.Println("create list")
	node := &Node{0, nil} // dummy
	return node
}

// PrintStack ... print linked list
func PrintStack(list *Node) {
	if list.next == nil {
		fmt.Println("list is empty!")
		return
	}

	for list.next != nil {
		fmt.Printf("%d ", list.value)
		list = list.next
	}
	fmt.Println("")
}

// Push ... push to top of stack
func Push(value int, list *Node) *Node {
	return &Node{value, list}
}

// Pop ... pop the top of the stack
func Pop(list *Node) *Node {
	if list.next == nil {
		fmt.Println("cannot remove empty list!")
		return list
	}
	return list.next
}
