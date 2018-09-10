package doublylinkedlist

import (
	"testing"
)

func TestNewNodeDefaultValues(t *testing.T) {
	node := NewNode(1)
	if node.value != 1 {
		t.Error("Expected `1`, actual", node.value)
	}
	if node.next != nil {
		t.Error("Expected `nil`, actual", node.next)
	}
	if node.prev != nil {
		t.Error("Expected `nil`, actual", node.prev)
	}
}

func TestNewListDefaultValues(t *testing.T) {
	list := NewList()
	if list.head != nil {
		t.Error("Expected `nil`, actual", list.head)
	}
	if list.tail != nil {
		t.Error("Expected `nil`, actual", list.tail)
	}
}
