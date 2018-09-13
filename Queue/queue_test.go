package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	if queue.list == nil {
		t.Error("List should not be empty")
	}
}

func TestIsEmptyOnEmptyList(t *testing.T) {
	queue := NewQueue()
	if empty := queue.IsEmpty(); empty != true {
		t.Errorf("Expected '%+v' actual '%+v'", true, empty)
	}
}

func TestIsEmptyOnFilledList(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	if notEmpty := queue.IsEmpty(); notEmpty != false {
		t.Errorf("Expected '%+v' actual '%+v'", false, notEmpty)
	}
}

func TestEnqueueAddNewElement(t *testing.T) {
	queue := NewQueue()
	node := queue.Enqueue(1)
	if queue.list.Head != node {
		t.Errorf("Expected '%+v' actual '%+v'", node, queue.list.Head)
	}
}

func TestPeekOnEmptyList(t *testing.T) {
	queue := NewQueue()
	if node := queue.Peek(); node != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, node)
	}
}

func TestPeekOnFilledList(t *testing.T) {
	queue := NewQueue()
	n10 := queue.Enqueue(10)
	n20 := queue.Enqueue(20)

	if peeked := queue.Peek(); peeked != n10 {
		t.Errorf("Expected '%+v' actual '%+v'", n10, peeked)
	}
	queue.Dequeue()
	if peeked := queue.Peek(); peeked != n20 {
		t.Errorf("Expected '%+v' actual '%+v'", n20, peeked)
	}
	queue.Dequeue()
	if peeked := queue.Peek(); peeked != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, peeked)
	}
}

func TestDequeueOnEmptyList(t *testing.T) {
	queue := NewQueue()
	if node := queue.Dequeue(); node != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, node)
	}
}

func TestDequeueOnFilledList(t *testing.T) {
	queue := NewQueue()
	n10 := queue.Enqueue(10)

	if node := queue.Dequeue(); node != n10 {
		t.Errorf("Expected '%+v' actual '%+v'", n10, node)
	}
	if queue.list.Head != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, queue.list.Head)
	}
}
