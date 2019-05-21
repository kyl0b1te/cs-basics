package binarytree

type BinaryTree struct {
	Root *Node
}

func NewBinaryTree() BinaryTree {
	return BinaryTree{nil}
}

func (t *BinaryTree) Insert(node *Node) {
	if t.Root == nil {
		t.Root = node
		return
	}
	t.insertNode(t.Root, node)
}

func (t BinaryTree) Contains(value int) bool {
	if t.Root == nil {
		return false
	}
	return t.containsNode(t.Root, value)
}

func (t BinaryTree) containsNode(current *Node, value int) bool {
	if value > current.Value && current.Right != nil {
		return t.containsNode(current.Right, value)
	}
	if value < current.Value && current.Left != nil {
		return t.containsNode(current.Left, value)
	}
	if value == current.Value {
		return true
	}
	return false
}

func (t BinaryTree) insertNode(current *Node, node *Node) {
	if node.Value < current.Value {
		if current.Left == nil {
			current.Left = node
		} else {
			t.insertNode(current.Left, node)
		}
	} else if node.Value > current.Value {
		if current.Right == nil {
			current.Right = node
		} else {
			t.insertNode(current.Right, node)
		}
	}
}
