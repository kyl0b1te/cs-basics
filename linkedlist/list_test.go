package linkedlist

import (
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	list := NewList()

	for i := 0; i < b.N; i++ {
		list.Append(NewNode(i))
	}
}

func BenchmarkIndex(b *testing.B) {
	list := NewList()
	list.Append(NewNode(43))
	list.Append(NewNode(37))
	list.Append(NewNode(83))

	for i := 0; i < b.N; i++ {
		list.Index(37)
	}
}

func TestNewList(t *testing.T) {
	list := NewList()
	if list.head != nil {
		t.Error("Expected list head is `nil` actual", list.head)
	}
}

func TestNewNode(t *testing.T) {
	node := NewNode(99)
	if node.value != 99 {
		t.Error("Expected node value is `99` actual", node.value)
	}
	if node.next != nil {
		t.Errorf("Expected node next is `nil` actual `%+v`", node.value)
	}
}

func TestAppendOneNode(t *testing.T) {
	list := NewList()
	node := NewNode(99)

	list.Append(node)
	if list.head != node {
		t.Errorf("Expected list head is equals to `%x` for `%x`", node, list.head)
	}
}

func TestAppendMultipleNodes(t *testing.T) {
	list := NewList()
	nodes := [...]*Node{NewNode(10), NewNode(20), NewNode(30), NewNode(40)}

	for _, node := range nodes {
		list.Append(node)
	}

	listNode := list.head
	for _, node := range nodes {
		if listNode != node {
			t.Fatalf("Expected list value is `%+v` actual `%+v`", node, listNode)
		}
		listNode = listNode.next
	}
}

func TestIndexOnEmptyList(t *testing.T) {
	list := NewList()
	if index := list.Index(99); index != -1 {
		t.Errorf("Expected `%+v` actual `%+v`", -1, index)
	}
}

func TestIndexOnFilledList(t *testing.T) {
	list := NewList()
	nodes := [...]*Node{NewNode(12), NewNode(83), NewNode(43), NewNode(64)}

	for _, node := range nodes {
		list.Append(node)
	}

	var index int
	for i, node := range nodes {
		if index = list.Index(node.value); i != index {
			t.Errorf("Expected value `%+v` has index `%+v` actual `%+v`", node.value, i, index)
		}
	}
}

func TestIndexOfMissingNode(t *testing.T) {
	list := NewList()
	list.Append(NewNode(10))
	if index := list.Index(99); index != -1 {
		t.Errorf("Expected `%+v` actual `%+v`", -1, index)
	}
}

func TestDeleteFirstNode(t *testing.T) {
	list := NewList()
	t10 := NewNode(10)
	t20 := NewNode(20)
	t30 := NewNode(30)

	list.Append(t10)
	list.Append(t20)
	list.Append(t30)

	list.Delete(10)

	if list.head != t20 {
		t.Errorf("Expected `%+v` actual `%+v`", t20, list.head)
	}
	if list.head.next != t30 {
		t.Errorf("Expected `%+v` actual `%+v`", t30, list.head.next)
	}
	if list.head.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head.next.next)
	}
}

func TestDeleteMiddleNode(t *testing.T) {
	list := NewList()
	t10 := NewNode(10)
	t20 := NewNode(20)
	t30 := NewNode(30)

	list.Append(t10)
	list.Append(t20)
	list.Append(t30)

	list.Delete(20)

	if list.head != t10 {
		t.Errorf("Expected `%+v` actual `%+v`", t10, list.head)
	}
	if list.head.next != t30 {
		t.Errorf("Expected `%+v` actual `%+v`", t30, list.head.next)
	}
	if list.head.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head.next.next)
	}
}

func TestDeleteLastNode(t *testing.T) {
	list := NewList()
	t10 := NewNode(10)
	t20 := NewNode(20)
	t30 := NewNode(30)
	t40 := NewNode(40)

	list.Append(t10)
	list.Append(t20)
	list.Append(t30)
	list.Append(t40)

	list.Delete(40)

	if list.head != t10 {
		t.Errorf("Expected `%+v` actual `%+v`", t10, list.head)
	}
	if list.head.next != t20 {
		t.Errorf("Expected `%+v` actual `%+v`", t20, list.head.next)
	}
	if list.head.next.next != t30 {
		t.Errorf("Expected `%+v` actual `%+v`", t30, list.head.next.next)
	}
	if list.head.next.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head.next.next.next)
	}
}

func TestDeleteAllNodes(t *testing.T) {
	list := NewList()

	list.Append(NewNode(10))
	list.Append(NewNode(10))
	list.Append(NewNode(10))

	list.Delete(10)

	if list.head != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head)
	}
}

func TestDeleteOnEmptyList(t *testing.T) {
	list := NewList()
	list.Delete(43)
	if list.head != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head)
	}
}

func TestPrependOnEmptyList(t *testing.T) {
	list := NewList()
	list.Prepend(NewNode(10))
	if list.head.value != 10 {
		t.Errorf("Expected `%+v` actual `%+v`", 10, list.head.value)
	}
}

func TestPrependOnFilledList(t *testing.T) {
	list := NewList()
	list.Append(NewNode(10))
	list.Append(NewNode(20))

	list.Prepend(NewNode(0))
	if list.head.value != 0 {
		t.Errorf("Expected `%+v` actual `%+v`", 0, list.head.value)
	}
	if list.head.next.value != 10 {
		t.Errorf("Expected `%+v` actual `%+v`", 10, list.head.next.value)
	}
	if list.head.next.next.value != 20 {
		t.Errorf("Expected `%+v` actual `%+v`", 20, list.head.next.next.value)
	}
	if list.head.next.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head.next.next.next)
	}
}
