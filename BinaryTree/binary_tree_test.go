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
	tree := NewBinaryTree()
	nodes := getNodes(10, 5, 15)

	for _, node := range nodes {
		tree.Insert(node)
	}

	if tree.Root.Value != 10 {
		t.Errorf("Expected '%+v' actual '%+v'", 10, tree.Root.Value)
	}
	if tree.Root.Left == nil {
		t.Errorf("Expected '%+v' to not be nil", tree.Root.Left)
	}
	if tree.Root.Left.Value != 5 {
		t.Errorf("Expected '%+v' actual '%+v'", 5, tree.Root.Left.Value)
	}
	if tree.Root.Right == nil {
		t.Errorf("Expected '%+v' to not be nil", tree.Root.Right)
	}
	if tree.Root.Right.Value != 15 {
		t.Errorf("Expected '%+v' actual '%+v'", 15, tree.Root.Right.Value)
	}
}

func TestContainsOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if contains := tree.Contains(NewNode(2, nil, nil)); contains != false {
		t.Errorf("Expected '%+v' actual '%+v'", contains, false)
	}
}

func TestContainsOnExistingChild(t *testing.T) {
	tree := NewBinaryTree()
	for _, node := range getNodes(4, 8, 2) {
		tree.Insert(node)
	}

	if contains := tree.Contains(NewNode(2, nil, nil)); contains != true {
		t.Errorf("Expected '%+v' actual '%+v'", contains, true)
	}
	if contains := tree.Contains(NewNode(4, nil, nil)); contains != true {
		t.Errorf("Expected '%+v' actual '%+v'", contains, true)
	}
	if contains := tree.Contains(NewNode(8, nil, nil)); contains != true {
		t.Errorf("Expected '%+v' actual '%+v'", contains, true)
	}
}

func TestContainsOnMissingChild(t *testing.T) {
	tree := NewBinaryTree()
	for _, node := range getNodes(4, 8, 2) {
		tree.Insert(node)
	}

	if contains := tree.Contains(NewNode(5, nil, nil)); contains != false {
		t.Errorf("Expected '%+v' actual '%+v'", contains, false)
	}
}
