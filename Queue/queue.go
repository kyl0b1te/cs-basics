package queue

import "github.com/zhikiri/data-structs/LinkedList"

// Queue represents queue data structure
type Queue struct {
	list *linkedlist.LinkedList
}

// NewQueue is a function for create a new queue
func NewQueue() *Queue {
	list := linkedlist.NewList()
	return &Queue{list: &list}
}

// IsEmpty is a function that check is queue is currently empty
func (queue Queue) IsEmpty() bool {
	return queue.list.Head == nil
}

// Enqueue is a function to add a new element into the queue
func (queue *Queue) Enqueue(value int) *linkedlist.Node {
	return queue.list.Append(value)
}

// Peek is a function for get next element
func (queue Queue) Peek() *linkedlist.Node {
	return queue.list.Head
}

// Dequeue is a function for remove element at the front of the queue
func (queue *Queue) Dequeue() *linkedlist.Node {
	return queue.list.DeleteHead()
}
