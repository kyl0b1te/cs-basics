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

func TestNewBinaryTree(t *testing.T) {
	tree := NewBinaryTree()
	if tree.Root != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, tree.Root)
	}
}

func TestInsertRoot(t *testing.T) {
	tree := NewBinaryTree()
	node := NewNode(10, nil, nil)
	tree.Insert(node)

	if tree.Root != node {
		t.Errorf("Expected '%+v' actual '%+v'", node, tree.Root)
	}
}

func TestInsertRootChildren(t *testing.T) {
	tree := NewBinaryTree(getNodes(10, 5, 15)...)

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
	tree := NewBinaryTree(getNodes(4, 8, 2)...)
	for _, val := range [3]int{2, 4, 8} {
		if contains := tree.Contains(val); contains != true {
			t.Errorf("Expected '%+v' actual '%+v'", contains, true)
		}
	}
}

func TestContainsOnMissingChild(t *testing.T) {
	tree := NewBinaryTree(getNodes(4, 8, 2)...)
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
	tree := NewBinaryTree(getNodes(4, 8, 2)...)
	if node := tree.Find(5); node != nil {
		t.Errorf("Expected '%+v' actual '%+v'", node, nil)
	}
}
