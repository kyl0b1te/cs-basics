package doublylinkedlist

// DoublyLinkedList represents doubly linked list data structure itself
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// NewList is a function for create new "instance" of doubly linked list
func NewList() DoublyLinkedList {
	return DoublyLinkedList{nil, nil}
}
