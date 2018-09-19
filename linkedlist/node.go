package linkedlist

// Node represents linked list node element
type Node struct {
	Value interface{}
	next  *Node
}

// NewNode is function to create a new node for linked list
func NewNode(value interface{}) *Node {
	return &Node{value, nil}
}
