package binarytree

type BinaryTree struct {
	Root *Node
}

func NewBinaryTree(nodes ...*Node) BinaryTree {
	if len(nodes) == 0 {
		return BinaryTree{nil}
	}

	tree := BinaryTree{nil}
	for _, node := range nodes {
		tree.Insert(node)
	}
	return tree
}

func (t *BinaryTree) Insert(node *Node) {
	if t.Root == nil {
		t.Root = node
		return
	}
	t.insertNode(t.Root, node)
}

func (t *BinaryTree) Contains(value int) bool {
	if t.Root == nil {
		return false
	}
	return t.search(t.Root, value) != nil
}

func (t *BinaryTree) Find(value int) *Node {
	if t.Root == nil {
		return nil
	}
	return t.search(t.Root, value)
}

func (t *BinaryTree) FindParent(value int) *Node {
	if t.Root == nil || t.Root.Value == value {
		return nil
	}
	return t.iterate(t.Root, func (current *Node) (*Node, bool) {
		if value < current.Value && current.Left != nil {
			if current.Left.Value == value {
				return current, false
			}
			return current.Left, true
		}
		if value > current.Value && current.Right != nil {
			if current.Right.Value == value {
				return current, false
			}
			return current.Right, true
		}
		return nil, false
	})
}

func (t *BinaryTree) Min() *Node {
	if t.Root == nil {
		return nil
	}
	return t.iterate(t.Root, func (current *Node) (*Node, bool) {
		if current.Left == nil {
			return current, false
		}
		return current.Left, true
	})
}

func (t *BinaryTree) Max() *Node {
	if t.Root == nil {
		return nil
	}
	return t.iterate(t.Root, func (current *Node) (*Node, bool) {
		if current.Right == nil {
			return current, false
		}
		return current.Right, true
	})
}

func (t BinaryTree) iterate(current *Node, fn func (node *Node) (*Node, bool)) *Node {
	next, keep := fn(current)
	if next != nil && keep == true {
		return t.iterate(next, fn)
	}
	return next
}

func (t BinaryTree) search(current *Node, value int) *Node {
	if value > current.Value && current.Right != nil {
		return t.search(current.Right, value)
	}
	if value < current.Value && current.Left != nil {
		return t.search(current.Left, value)
	}
	if value == current.Value {
		return current
	}
	return nil
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
