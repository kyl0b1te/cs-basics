package linkedlist

// LinkedList represents linked list data structure itself
type LinkedList struct {
	Head *Node
	Tail *Node
}

// NewList is a function for create new "instance" of linked list
func NewList() LinkedList {
	return LinkedList{nil, nil}
}

// Index is a function for get first node index that has specific value
func (list LinkedList) Index(value int) int {
	if list.Head == nil {
		return -1
	}
	listNode := list.Head
	for i := 0; listNode != nil; i++ {
		if listNode.value == value {
			return i
		}
		listNode = listNode.next
	}
	return -1
}

// Find is a function for get first node with specific value
func (list LinkedList) Find(value int) *Node {
	listNode := list.Head
	for listNode != nil {
		if listNode.value == value {
			return listNode
		}
		listNode = listNode.next
	}
	return nil
}

// Append is a function for add new node into the end of linked list
func (list *LinkedList) Append(value int) *Node {
	node := NewNode(value)
	if list.Head == nil {
		list.Head = node
		list.Tail = node
		return node
	}
	previous := list.Tail
	previous.next = node
	list.Tail = node
	return node
}

// Prepend is a function for add new node into the begging of linked list
func (list *LinkedList) Prepend(value int) *Node {
	node := NewNode(value)
	if list.Head == nil {
		list.Head = node
		return list.Head
	}
	node.next = list.Head
	list.Head = node
	return node
}

// Delete is a function for delete node(s) from the linked list
func (list *LinkedList) Delete(value int) {
	if list.Head == nil {
		return
	}

	// Remove nodes from the begging of the list
	for list.Head != nil && list.Head.value == value {
		list.Head = list.Head.next
	}

	listNode := list.Head
	for listNode != nil {
		if listNode.next != nil && listNode.next.value == value {
			if listNode.next.next != nil {
				// Next for current node equal to the node after next one
				listNode.next = listNode.next.next
			} else {
				// Made current node is end of the list
				listNode.next = nil
			}
		}
		listNode = listNode.next
	}
}

// DeleteHead is a function for remove list head element
func (list *LinkedList) DeleteHead() *Node {
	if list.Head == nil {
		return nil
	}
	head := list.Head
	list.Head = head.next
	return head
}
