package linkedlist

// Node represents linked list node element
type Node struct {
	value int
	next  *Node
}

// NewNode is function to create a new node for linked list
func NewNode(value int) *Node {
	return &Node{value, nil}
}
