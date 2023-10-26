package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
}

// Initialize a new list. The number of elements size.
func New(size int) *LinkedList {

	if size <= 0 {
		fmt.Println("ERROR (in New): list size is set incorrectly")
		return nil
	}

	head := node{}
	current := &head
	for i := 1; i < size; i++ {
		current.next = &node{}
		current = current.next
	}
	return &LinkedList{head: &head}
}

// List output
func (l *LinkedList) PrintList() {
	current := l.head

	for current.next != nil {
		fmt.Printf("%d -> ", current.val)
		current = current.next
	}
	fmt.Printf("%d\n", current.val)
}

// Return the length of the list
func (l *LinkedList) Size() int {
	size := 1
	current := l.head
	for current.next != nil {
		size++
		current = current.next
	}
	return size
}

// Getting an item from the pos position.
func (l *LinkedList) At(pos int) int {

	if pos < 0 || pos >= l.Size() {
		fmt.Println("ERROR (in At): list index out of range")
		return 0
	}

	current := l.head
	for i := 0; i != pos && current.next != nil; i++ {
		current = current.next
	}

	return current.val
}

// Add an item to the end of the list.
func (l *LinkedList) Add(val int) {
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &node{val: val, next: nil}
}

// Make the value at the pos position equal to val.
func (l *LinkedList) UpdateAt(pos int, val int) {

	if pos < 0 || pos >= l.Size() {
		fmt.Println("ERROR (in UpdateAt): list index out of range")
		return
	}

	current := l.head
	for i := 0; i != pos && current.next != nil; i++ {
		current = current.next
	}

	current.val = val
}

// Remove an element from the end.
func (l *LinkedList) Pop() {

	if l.head.next == nil {
		fmt.Println("ERROR (in Pop): Unable to delete element, list size is already 1")
		return
	}

	current := l.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

// Remove an item from the pos position.
func (l *LinkedList) DeleteAt(pos int) {

	if pos < 0 || pos >= l.Size() {
		fmt.Println("ERROR (in DeleteAt): list index out of range")
		return
	}

	if l.head.next == nil {
		fmt.Println("ERROR (in DeleteAt): Unable to delete element, list size is already 1")
		return
	}

	if pos == 0 {
		l.head = l.head.next
		return
	}

	current := l.head
	for i := 0; i != pos-1 && current.next != nil; i++ {
		current = current.next
	}

	current.next = current.next.next
}

// Creating a linked list from a slice.
func NewFromSlise(s []int) *LinkedList {

	if len(s) == 0 {
		fmt.Println("ERROR (in NewFromSlise): Unable convert an empty slice to a list")
	}

	head := node{val: s[0]}
	current := &head
	for i := 0; i < len(s)-1; i++ {
		current.next = &node{val: s[i+1]}
		current = current.next
	}
	return &LinkedList{head: &head}
}
