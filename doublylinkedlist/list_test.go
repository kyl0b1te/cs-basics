package doublylinkedlist

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

func TestAppendOneNode(t *testing.T) {
	list := NewList()
	node := list.Append(99)
	if list.head != node {
		t.Errorf("Expected list head is equals to `%x` for `%x`", node, list.head)
	}
}

func TestAppendMultipleNodes(t *testing.T) {
	list := NewList()

	var nodes []*Node
	for i, value := range [...]int{10, 20, 30, 40} {
		node := list.Append(value)
		if i > 0 && nodes[i-1].next != node {
			t.Fatalf("Expected `%+v` actual `%+v`", nodes[i-1].next, node)
		}
		nodes = append(nodes, node)
	}

	for i, node := range nodes {
		if i < 3 && nodes[i+1].prev != node {
			t.Fatalf("Expected `%+v` actual `%+v`", nodes[i+1].prev, node)
		}
	}

	listNode := list.head
	for _, node := range nodes {
		if listNode != node {
			t.Fatalf("Expected `%+v` actual `%+v`", node, listNode)
		}
		listNode = listNode.next
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
	n10 := list.Append(10)
	n20 := list.Append(20)

	node := list.Prepend(0)
	if list.head != node {
		t.Errorf("Expected `%+v` actual `%+v`", node, list.head)
	}
	if list.head.next != n10 {
		t.Errorf("Expected `%+v` actual `%+v`", n10, list.head.next)
	}
	if list.head.next.next != n20 {
		t.Errorf("Expected `%+v` actual `%+v`", n20, list.head.next.next)
	}
	if list.head.next.next.prev != n10 {
		t.Errorf("Expected `%+v` actual `%+v`", n10, list.head.next.next.prev)
	}
	if list.head.next.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head.next.next.next)
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

func TestDeleteFirstNode(t *testing.T) {
	list := NewList()

	list.Append(10)
	n20 := list.Append(20)
	n30 := list.Append(30)

	list.Delete(10)

	if list.head != n20 {
		t.Errorf("Expected `%+v` actual `%+v`", n20, list.head)
	}
	if list.head.next != n30 {
		t.Errorf("Expected `%+v` actual `%+v`", n30, list.head.next)
	}
	if list.head.next.prev != n20 {
		t.Errorf("Expected `%+v` actual `%+v`", n20, list.head.next.prev)
	}
	if list.head.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head.next.next)
	}
}

func TestDeleteMiddleNode(t *testing.T) {
	list := NewList()

	n10 := list.Append(10)
	list.Append(20)
	n30 := list.Append(30)

	list.Delete(20)

	if list.head != n10 {
		t.Errorf("Expected `%+v` actual `%+v`", n10, list.head)
	}
	if list.head.next != n30 {
		t.Errorf("Expected `%+v` actual `%+v`", n30, list.head.next)
	}
	if list.head.next.prev != n10 {
		t.Errorf("Expected `%+v` actual `%+v`", n10, list.head.next.prev)
	}
	if list.head.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.head.next.next)
	}
}

func TestDeleteLastNode(t *testing.T) {
	list := NewList()

	n10 := list.Append(10)
	n20 := list.Append(20)
	n30 := list.Append(30)
	list.Append(40)

	list.Delete(40)

	if list.head != n10 {
		t.Errorf("Expected `%+v` actual `%+v`", n10, list.head)
	}
	if list.head.next != n20 {
		t.Errorf("Expected `%+v` actual `%+v`", n20, list.head.next)
	}
	if list.head.next.next != n30 {
		t.Errorf("Expected `%+v` actual `%+v`", n30, list.head.next.next)
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
