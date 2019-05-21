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

func (t BinaryTree) Contains(node *Node) bool {
	if t.Root == nil {
		return false
	}
	return t.containsNode(t.Root, node)
}

func (t BinaryTree) containsNode(current *Node, node *Node) bool {
	if node.Value > current.Value && current.Right != nil {
		return t.containsNode(current.Right, node)
	}
	if node.Value < current.Value && current.Left != nil {
		return t.containsNode(current.Left, node)
	}
	if node.Value == current.Value {
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
