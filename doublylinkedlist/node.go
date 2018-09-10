package doublylinkedlist

// Node represents doubly linked list node element
type Node struct {
	value int
	next  *Node
	prev  *Node
}

// NewNode is function to create a new node for doubly linked list
func NewNode(value int) *Node {
	return &Node{value, nil, nil}
}
