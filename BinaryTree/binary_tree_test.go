package binarytree

import (
	"testing"
)

func getNodes(values ...int) []*Node {
	nodes := make([]*Node, len(values))
	for _, value := range values {
		nodes = append(nodes, NewNode(value, nil, nil))
	}
	return nodes
}

func TestNewBinaryTreeWithoutArguments(t *testing.T) {
	tree := NewBinaryTree()
	if tree.Root != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, tree.Root)
	}
}

func TestNewBinaryTreeWithMultipleChild(t *testing.T) {
	tree := NewBinaryTree(10, 5, 15)

	if tree.Root.Value != 10 {
		t.Errorf("Expected '%+v' actual '%+v'", 10, tree.Root)
	}
	if tree.Root.Left.Value != 5 {
		t.Errorf("Expected '%+v' actual '%+v'", 5, tree.Root.Left.Value)
	}
	if tree.Root.Right.Value != 15 {
		t.Errorf("Expected '%+v' actual '%+v'", 15, tree.Root.Right.Value)
	}
}

func TestInsertRoot(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(10)

	if tree.Root.Value != 10 {
		t.Errorf("Expected '%+v' actual '%+v'", 10, tree.Root.Value)
	}
}

func TestInsertMultipleChildren(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	node := tree.Root
	for _, value := range [3]int{10, 5, 15} {
		if node == nil {
			t.Errorf("Expected '%+v' to not be nil", node)
		}
		if node.Value != value {
			t.Errorf("Expected '%+v' actual '%+v'", value, node.Value)
		}
		if node.Value == 10 {
			node = tree.Root.Left
			continue
		}
		node = tree.Root.Right
	}
}

func TestContainsOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if contains := tree.Contains(2); contains != false {
		t.Errorf("Expected '%+v' actual '%+v'", contains, false)
	}
}

func TestContainsOnExistingChild(t *testing.T) {
	tree := NewBinaryTree(4, 8, 2)
	for _, val := range [3]int{2, 4, 8} {
		if contains := tree.Contains(val); contains != true {
			t.Errorf("Expected '%+v' actual '%+v'", contains, true)
		}
	}
}

func TestContainsOnMissingChild(t *testing.T) {
	tree := NewBinaryTree(4, 8, 2)
	if contains := tree.Contains(5); contains != false {
		t.Errorf("Expected '%+v' actual '%+v'", contains, false)
	}
}

func TestFindOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if node := tree.Find(5); node != nil {
		t.Errorf("Expected '%+v' actual '%+v'", node, nil)
	}
}

func TestFindOnMissingChild(t *testing.T) {
	tree := NewBinaryTree(4, 8, 2)
	if node := tree.Find(5); node != nil {
		t.Errorf("Expected '%+v' actual '%+v'", node, nil)
	}
}

func TestFindOnExistingChild(t *testing.T) {
	tree := NewBinaryTree(4, 8, 2)
	for _, value := range [3]int{2, 4, 8} {
		if node := tree.Find(value); node == nil {
			t.Errorf("Expected '%+v' to not be nil", node)
		}
	}
}

func TestMinOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if min := tree.Min(); min != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, min)
	}
}

func TestMinOnRootOnly(t *testing.T) {
	tree := NewBinaryTree(2)
	if min := tree.Min(); min.Value != 2 {
		t.Errorf("Expected '%+v' actual '%+v'", 2, min.Value)
	}
}

func TestMinOnMultipleChild(t *testing.T) {
	tree := NewBinaryTree(20, 6, 8, 43, 12, 53)
	if min := tree.Min(); min.Value != 6 {
		t.Errorf("Expected '%+v' actual '%+v'", min.Value, 6)
	}
}

func TestMaxOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if max := tree.Max(); max != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, max)
	}
}

func TestMaxOnRootOnly(t *testing.T) {
	tree := NewBinaryTree(2)
	if max := tree.Max(); max.Value != 2 {
		t.Errorf("Expected '%+v' actual '%+v'", 2, max.Value)
	}
}

func TestMaxOnMultipleChild(t *testing.T) {
	tree := NewBinaryTree(20, 6, 8, 43, 12, 53)
	if max := tree.Max(); max.Value != 53 {
		t.Errorf("Expected '%+v' actual '%+v'", max.Value, 53)
	}
}

func TestFindParentOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if parent := tree.FindParent(5); parent != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, parent)
	}
}

func TestFindParentOnRoot(t *testing.T) {
	tree := NewBinaryTree(5)
	if parent := tree.FindParent(5); parent != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, parent)
	}
}

// 32
// 1 62
//   54 71
//  34    83
func TestFindParentOnMultipleChild(t *testing.T) {
	tree := NewBinaryTree(32, 62, 1, 54, 71, 34, 83)
	tests := map[int]int{1: 32, 62: 32, 54: 62, 71: 62, 34: 54, 83: 71}
	for value, expected := range tests {
		if par := tree.FindParent(value); par.Value != expected {
			t.Errorf("Expected '%+v' actual '%+v'", expected, par.Value)
		}
	}
}

func TestDeleteOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if actual := tree.Delete(0); actual != false {
		t.Errorf("Expected '%+v' actual '%+v'", false, actual)
	}
}

func TestDeleteOnExistingRoot(t *testing.T) {
	tree := NewBinaryTree(10)
	if actual := tree.Delete(10); actual != true {
		t.Errorf("Expected '%+v' actual '%+v'", true, actual)
	}
	if tree.Root != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, tree.Root)
	}
}

func TestDeleteOnMissingValue(t *testing.T) {
	tree := NewBinaryTree(10, 15, 25)
	if actual := tree.Delete(8); actual != false {
		t.Errorf("Expected '%+v' actual '%+v'", false, actual)
	}
}

func TestDeleteOnMultipleChild(t *testing.T) {
	values := []int{10, 54, 14, 61, 45, 93, 76}
	tree := NewBinaryTree(values...)

	for i := len(values) - 1; i >= 0; i -= 1 {
		value := values[i]
		if actual := tree.Delete(value); actual != true {
			t.Errorf("Delete value %d expected to be '%+v' actual '%+v'", value, true, actual)
		}
	}
	if tree.Root != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, tree.Root)
	}
}
