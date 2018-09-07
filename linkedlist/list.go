package linkedlist

// LinkedList represents linked list data structure itself
type LinkedList struct {
	head *Node
	tail *Node
}

// NewList is a function for create new "instance" of linked list
func NewList() LinkedList {
	return LinkedList{nil, nil}
}

// Index is a function for get first node index that has specific value
func (list LinkedList) Index(value int) int {
	if list.head == nil {
		return -1
	}
	listNode := list.head
	for i := 0; listNode != nil; i++ {
		if listNode.value == value {
			return i
		}
		listNode = listNode.next
	}
	return -1
}

// Find is a funcation for get first node with specific value
func (list LinkedList) Find(value int) *Node {
	listNode := list.head
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
	if list.head == nil {
		list.head = node
		list.tail = node
		return node
	}
	previous := list.tail
	previous.next = node
	list.tail = node
	return node
}

// Prepend is a function for add new node into the begging of linked list
func (list *LinkedList) Prepend(value int) *Node {
	node := NewNode(value)
	if list.head == nil {
		list.head = node
		return list.head
	}
	node.next = list.head
	list.head = node
	return node
}

// Delete is a function for delete node(s) from the linked list
func (list *LinkedList) Delete(value int) {
	if list.head == nil {
		return
	}

	// Remove nodes from the begging of the list
	for list.head != nil && list.head.value == value {
		list.head = list.head.next
	}

	listNode := list.head
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
