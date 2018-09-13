package queue

import "github.com/zhikiri/data-structs/LinkedList"

// Queue represents abstract data structure
type Queue struct {
	list *linkedlist.LinkedList
}

// NewQueue is a function that creates a new queue
func NewQueue() *Queue {
	list := linkedlist.NewList()
	return &Queue{list: &list}
}

// IsEmpty is a function that check is queue currently empty
func (queue Queue) IsEmpty() bool {
	return queue.list.Head == nil
}

// Enqueue is a function that add a new element into the queue
func (queue *Queue) Enqueue(value int) *linkedlist.Node {
	return queue.list.Append(value)
}

// Peek is a function that retrieve next element from the queue
func (queue Queue) Peek() *linkedlist.Node {
	return queue.list.Head
}

// Dequeue is a function that removes element at the front of the queue
func (queue *Queue) Dequeue() *linkedlist.Node {
	return queue.list.DeleteHead()
}
