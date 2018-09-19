package linkedlist

// LinkedList represents abstract data structure
type LinkedList struct {
	Head    *Node
	Tail    *Node
	isEqual func(a, b interface{}) bool
}

// NewList is a function that creates a new "instance" of linked list
func NewList(compare func(a, b interface{}) bool) LinkedList {
	return LinkedList{nil, nil, compare}
}

// DefaultCompare is a function that represents integer comparassion
func DefaultCompare(a, b interface{}) bool {
	aInt := a.(int)
	bInt := b.(int)
	return aInt == bInt
}

// Index is a function that gets index of first node where value equal index
func (list LinkedList) Index(value interface{}) int {
	if list.Head == nil {
		return -1
	}
	listNode := list.Head
	for i := 0; listNode != nil; i++ {
		if list.isEqual(listNode.Value, value) {
			return i
		}
		listNode = listNode.next
	}
	return -1
}

// Find is a function that gets a first node with specific value
func (list LinkedList) Find(value int) *Node {
	listNode := list.Head
	for listNode != nil {
		if list.isEqual(listNode.Value, value) {
			return listNode
		}
		listNode = listNode.next
	}
	return nil
}

// Append is a function that add new node into the end of linked list
func (list *LinkedList) Append(value interface{}) *Node {
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

// Prepend is a function that add new node into the begging of linked list
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

// Delete is a function that removes node(s) from the linked list
func (list *LinkedList) Delete(value int) {
	if list.Head == nil {
		return
	}

	// Remove nodes from the begging of the list
	for list.Head != nil && list.isEqual(list.Head.Value, value) {
		list.Head = list.Head.next
	}

	listNode := list.Head
	for listNode != nil {
		if listNode.next != nil && list.isEqual(listNode.next.Value, value) {
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

// DeleteHead is a function that removes list head
func (list *LinkedList) DeleteHead() *Node {
	if list.Head == nil {
		return nil
	}
	head := list.Head
	list.Head = head.next
	return head
}

// Map is a function that iterate through list via custom predicate function
func (list LinkedList) Map(predicate func(*Node) bool) *Node {
	if list.Head == nil {
		return nil
	}
	listNode := list.Head
	for listNode != nil {
		if predicate(listNode) == true {
			return listNode
		}
		listNode = listNode.next
	}
	return nil
}
