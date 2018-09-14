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
	if list.Head != nil {
		t.Error("Expected list Head is `nil` actual", list.Head)
	}
	if list.Tail != nil {
		t.Error("Expected list Tail is `nil` actual", list.Tail)
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
	if list.Head != node {
		t.Errorf("Expected list Head is equals to `%x` for `%x`", node, list.Head)
	}
}

func TestAppendMultipleNodes(t *testing.T) {
	list := NewList()
	values := [...]int{10, 20, 30, 40}

	for _, value := range values {
		list.Append(value)
	}

	listNode := list.Head
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

	if list.Head.value != 20 {
		t.Errorf("Expected `%+v` actual `%+v`", 20, list.Head.value)
	}
	if list.Head.next.value != 30 {
		t.Errorf("Expected `%+v` actual `%+v`", 30, list.Head.next.value)
	}
	if list.Head.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.Head.next.next)
	}
}

func TestDeleteMiddleNode(t *testing.T) {
	list := NewList()

	list.Append(10)
	list.Append(20)
	list.Append(30)

	list.Delete(20)

	if list.Head.value != 10 {
		t.Errorf("Expected `%+v` actual `%+v`", 10, list.Head.value)
	}
	if list.Head.next.value != 30 {
		t.Errorf("Expected `%+v` actual `%+v`", 30, list.Head.next.value)
	}
	if list.Head.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.Head.next.next)
	}
}

func TestDeleteLastNode(t *testing.T) {
	list := NewList()

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)

	list.Delete(40)

	if list.Head.value != 10 {
		t.Errorf("Expected `%+v` actual `%+v`", 10, list.Head.value)
	}
	if list.Head.next.value != 20 {
		t.Errorf("Expected `%+v` actual `%+v`", 20, list.Head.next.value)
	}
	if list.Head.next.next.value != 30 {
		t.Errorf("Expected `%+v` actual `%+v`", 30, list.Head.next.next.value)
	}
	if list.Head.next.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.Head.next.next.next)
	}
}

func TestDeleteAllNodes(t *testing.T) {
	list := NewList()

	list.Append(10)
	list.Append(10)
	list.Append(10)

	list.Delete(10)

	if list.Head != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.Head)
	}
}

func TestDeleteOnEmptyList(t *testing.T) {
	list := NewList()
	list.Delete(43)
	if list.Head != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.Head)
	}
}

func TestPrependOnEmptyList(t *testing.T) {
	list := NewList()
	node := list.Prepend(10)
	if list.Head != node {
		t.Errorf("Expected `%+v` actual `%+v`", node, list.Head)
	}
}

func TestPrependOnFilledList(t *testing.T) {
	list := NewList()
	list.Append(10)
	list.Append(20)

	node := list.Prepend(0)
	if list.Head != node {
		t.Errorf("Expected `%+v` actual `%+v`", node, list.Head)
	}
	if list.Head.next.value != 10 {
		t.Errorf("Expected `%+v` actual `%+v`", 10, list.Head.next.value)
	}
	if list.Head.next.next.value != 20 {
		t.Errorf("Expected `%+v` actual `%+v`", 20, list.Head.next.next.value)
	}
	if list.Head.next.next.next != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.Head.next.next.next)
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

func TestDeleteHeadOnEmptyList(t *testing.T) {
	list := NewList()
	if node := list.DeleteHead(); node != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, node)
	}
}

func TestDeleteHeadOnFilledList(t *testing.T) {
	list := NewList()
	n10 := list.Append(10)
	n20 := list.Append(20)
	n30 := list.Append(30)

	if node := list.DeleteHead(); node != n10 {
		t.Errorf("Expected `%+v` actual `%+v`", n10, node)
	}
	if list.Head != n20 {
		t.Errorf("Expected `%+v` actual `%+v`", n20, list.Head)
	}
	if node := list.DeleteHead(); node != n20 {
		t.Errorf("Expected `%+v` actual `%+v`", n20, node)
	}
	if list.Head != n30 {
		t.Errorf("Expected `%+v` actual `%+v`", n30, list.Head)
	}
	if node := list.DeleteHead(); node != n30 {
		t.Errorf("Expected `%+v` actual `%+v`", n30, node)
	}
	if list.Head != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, list.Head)
	}
}

func TestMapOnEmptyList(t *testing.T) {
	list := NewList()
	res := list.Map(func(node *Node) bool {
		return node.value == 1
	})
	if res != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, res)
	}
}

func TestMapWithFalsyPredicate(t *testing.T) {
	list := NewList()
	nodes := [...]*Node{list.Append(10), list.Append(20)}

	counter := 0
	res := list.Map(func(node *Node) bool {
		if nodes[counter] != node {
			t.Errorf("Expected `%+v` actual `%+v`", nodes[counter], node)
		}
		counter++
		return false
	})

	if res != nil {
		t.Errorf("Expected `%+v` actual `%+v`", nil, res)
	}
}

func TestMapWithTruePredicate(t *testing.T) {
	list := NewList()
	list.Append(10)
	n20 := list.Append(20)

	res := list.Map(func(node *Node) bool {
		return node == n20
	})

	if res != n20 {
		t.Errorf("Expected `%+v` actual `%+v`", n20, res)
	}
}
