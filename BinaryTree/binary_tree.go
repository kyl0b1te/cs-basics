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
