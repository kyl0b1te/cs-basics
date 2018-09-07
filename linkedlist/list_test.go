package linkedlist

import (
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	list := NewList()
	for i := 0; i < b.N; i++ {
		list.Append(i)
	}
}

func BenchmarkPrepend(b *testing.B) {
	list := NewList()
	for i := 0; i < b.N; i++ {
		list.Prepend(i)
	}
}

func BenchmarkIndex(b *testing.B) {
	list := NewList()
	list.Append(43)
	list.Append(37)
	list.Append(83)

	for i := 0; i < b.N; i++ {
		list.Index(37)
	}
}

func BenchmarkFind(b *testing.B) {
	list := NewList()
	list.Append(43)
	list.Append(37)
	list.Append(83)

	for i := 0; i < b.N; i++ {
		list.Find(37)
	}
}

func TestNewList(t *testing.T) {
	list := NewList()
	if list.head != nil {
		t.Error("Expected list head is `nil` actual", list.head)
	}
	if list.tail != nil {
		t.Error("Expected list tail is `nil` actual", list.tail)
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
	node := list.Append(99)
	if list.head != node {
		t.Errorf("Expected list head is equals to `%x` for `%x`", node, list.head)
	}
}

func TestAppendMultipleNodes(t *testing.T) {
	list := NewList()
	values := [...]int{10, 20, 30, 40}

	for _, value := range values {
		list.Append(value)
	}

	listNode := list.head
	for _, value := range values {
		if listNode.value != value {
			t.Fatalf("Expected list value is `%+v` actual `%+v`", value, listNode.value)
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
		list.Append(node.value)
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
	list.Append(10)
	if index := list.Index(99); index != -1 {
		t.Errorf("Expected `%+v` actual `%+v`", -1, index)
	}
}

func TestDeleteFirstNode(t *testing.T) {
	list := NewList()

	list.Append(10)
	list.Append(20)
	list.Append(30)

	list.Delete(10)

	if list.head.value != 20 {
		t.Errorf("Expected `%+v` actual `%+v`", 20, list.head.value)
	}
	if list.head.next.value != 30 {
		t.Errorf("Expected `%+v` actual `%+v`", 30, list.head.next.value)
	}
	if list.head.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head.next.next)
	}
}

func TestDeleteMiddleNode(t *testing.T) {
	list := NewList()

	list.Append(10)
	list.Append(20)
	list.Append(30)

	list.Delete(20)

	if list.head.value != 10 {
		t.Errorf("Expected `%+v` actual `%+v`", 10, list.head.value)
	}
	if list.head.next.value != 30 {
		t.Errorf("Expected `%+v` actual `%+v`", 30, list.head.next.value)
	}
	if list.head.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head.next.next)
	}
}

func TestDeleteLastNode(t *testing.T) {
	list := NewList()

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)

	list.Delete(40)

	if list.head.value != 10 {
		t.Errorf("Expected `%+v` actual `%+v`", 10, list.head.value)
	}
	if list.head.next.value != 20 {
		t.Errorf("Expected `%+v` actual `%+v`", 20, list.head.next.value)
	}
	if list.head.next.next.value != 30 {
		t.Errorf("Expected `%+v` actual `%+v`", 30, list.head.next.next.value)
	}
	if list.head.next.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head.next.next.next)
	}
}

func TestDeleteAllNodes(t *testing.T) {
	list := NewList()

	list.Append(10)
	list.Append(10)
	list.Append(10)

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
	node := list.Prepend(10)
	if list.head != node {
		t.Errorf("Expected `%+v` actual `%+v`", node, list.head)
	}
}

func TestPrependOnFilledList(t *testing.T) {
	list := NewList()
	list.Append(10)
	list.Append(20)

	node := list.Prepend(0)
	if list.head != node {
		t.Errorf("Expected `%+v` actual `%+v`", node, list.head)
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

func TestFindOnEmptyList(t *testing.T) {
	list := NewList()
	if node := list.Find(34); node != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, node)
	}
}

func TestFindNotExisting(t *testing.T) {
	list := NewList()
	list.Append(43)
	if node := list.Find(34); node != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, node)
	}
}

func TestFindExisting(t *testing.T) {
	list := NewList()
	list.Append(34)
	if node := list.Find(34); node.value != 34 {
		t.Errorf("Expected `%+v` actual `%+v`", 34, node.value)
	}
}

func TestFindFirstFoundNode(t *testing.T) {
	list := NewList()

	list.Append(10)
	test := list.Append(20)
	list.Append(20)
	list.Append(30)

	if node := list.Find(20); node != test {
		t.Errorf("Expected `%+v` actual `%+v`", test, node)
	}
}
