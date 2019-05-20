package binarytree

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func NewNode(value int, left *Node, right *Node) *Node {
	return &Node{value, left, right}
}
