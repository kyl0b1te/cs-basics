package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	assert.NotNil(t, stack.list)
}

func TestIsEmptyOnEmptyStack(t *testing.T) {
	stack := NewStack()
	assert.True(t, stack.IsEmpty())
}

func TestIsEmptyOnFilledStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	assert.False(t, stack.IsEmpty())
}

func TestPush(t *testing.T) {
	stack := NewStack()

	tests := []int{1, 2, 3}
	for _, val := range tests {
		node := stack.Push(val)
		assert.Equal(t, node, stack.list.Head)
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	n10 := stack.Push(10)
	n20 := stack.Push(20)

	assert.Equal(t, n20, stack.Pop())
	assert.Equal(t, n10, stack.Pop())
	assert.Nil(t, stack.Pop())
}
