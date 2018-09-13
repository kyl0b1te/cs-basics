package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	if stack.list == nil {
		t.Error("Stack should not be empty")
	}
}

func TestIsEmptyOnEmptyStack(t *testing.T) {
	stack := NewStack()
	if stack := stack.IsEmpty(); stack != true {
		t.Errorf("Expected '%+v' actual '%+v'", true, stack)
	}
}

func TestIsEmptyOnFilledStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	if notEmpty := stack.IsEmpty(); notEmpty != false {
		t.Errorf("Expected '%+v' actual '%+v'", false, notEmpty)
	}
}

func TestPush(t *testing.T) {
	stack := NewStack()
	n10 := stack.Push(10)
	if stack.list.Head != n10 {
		t.Errorf("Expected '%+v' actual '%+v'", n10, stack.list.Head)
	}
	if stack.Push(20); stack.list.Head != n10 {
		t.Errorf("Expected '%+v' actual '%+v'", n10, stack.list.Head)
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	n10 := stack.Push(10)
	n20 := stack.Push(20)

	if node := stack.Pop(); node != n20 {
		t.Errorf("Expected '%+v' actual '%+v'", n20, node)
	}
	if node := stack.Pop(); node != n10 {
		t.Errorf("Expected '%+v' actual '%+v'", n10, node)
	}
	if node := stack.Pop(); node != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, node)
	}
}
