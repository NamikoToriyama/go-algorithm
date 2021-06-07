package llist

import (
	"fmt"
	"strconv"
)

// Node ... element of linked list
type Node struct {
	value int
	next  *Node
}

// InputIndex ... return the index entered by the user
func InputIndex() int {
	var index int
	fmt.Printf("index > ")
	if sc.Scan() {
		t := sc.Text()
		index, _ = strconv.Atoi(t)
	}
	return index
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

// CreateList ...create linked list
func CreateList() *Node {
	fmt.Println("create list")
	node := &Node{0, nil} // dummy
	return node
}

// PrintList ... print linked list
func PrintList(list *Node) {
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

// InsertNode ...insert node from an index parameter
func InsertNode(value int, n int, list *Node) *Node {
	tmp := list
	if n == 0 {
		return &Node{value, list}
	}
	if tmp.next == nil {
		fmt.Println("cannot insert list!")
		return list
	}
	for range make([]int, n-1) {
		tmp = tmp.next
		if tmp.next == nil {
			fmt.Println("cannot insert list!")
			return list
		}

	}
	nextPointer := tmp.next
	tmp.next = &Node{value, nextPointer}
	return list
}

// RemoveNode ... remove node from an index parameter
func RemoveNode(index int, list *Node) *Node {
	if list.next == nil {
		fmt.Println("cannot remove empty list!")
		return list
	}

	if index == 0 {
		return list.next
	}

	tmp := list
	for range make([]int, index-1) {
		if tmp.next == nil {
			fmt.Println("cannot remove list!")
			return list
		}
		tmp = tmp.next
	}
	exPointer := tmp // Previous pointer
	if tmp.next == nil {
		exPointer.next = nil
		fmt.Println("cannot remove list!")
		return list
	}
	tmp = tmp.next
	exPointer.next = tmp.next

	return list
}

// ReverseList ... reverse linked list
func ReverseList(list *Node) *Node {
	reverseList := &Node{list.value, nil}

	if list.next == nil {
		fmt.Println("list is empty!")
		return list
	}

	for list.next != nil {
		tmp := list.next
		// Change the direction of the pointer
		list.next = reverseList
		reverseList = list
		list = tmp
	}
	return reverseList
}
