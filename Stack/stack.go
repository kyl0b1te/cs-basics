package stack

import "github.com/zhikiri/data-structs/LinkedList"

// Stack represents abstract data structure
type Stack struct {
	list *linkedlist.LinkedList
}

// NewStack is a function that creates a new Stack
func NewStack() *Stack {
	list := linkedlist.NewList(linkedlist.DefaultCompare)
	return &Stack{list: &list}
}

// IsEmpty is a function that check is stack currently empty
func (Stack Stack) IsEmpty() bool {
	return Stack.list.Head == nil
}

// Push is a function that add a new element into the stack
func (Stack *Stack) Push(value int) *linkedlist.Node {
	return Stack.list.Prepend(value)
}

// Pop is a function that retrieve next element from stack
func (Stack Stack) Pop() *linkedlist.Node {
	return Stack.list.DeleteHead()
}
