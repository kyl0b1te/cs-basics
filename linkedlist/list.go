package linkedlist

// LinkedList represents linked list data structure itself
type LinkedList struct {
	head *Node
}

// NewList is a function for create new "instance" of linked list
func NewList() LinkedList {
	return LinkedList{nil}
}

// Index is a function for get first node index that has specific value
func (list LinkedList) Index(searchValue int) int {
	if list.head == nil {
		return -1
	}
	listNode := list.head
	for i := 0; listNode != nil; i++ {
		if listNode.value == searchValue {
			return i
		}
		listNode = listNode.next
	}
	return -1
}

// Append is a function for add new node into the end of linked list
func (list *LinkedList) Append(node *Node) {
	if list.head == nil {
		list.head = node
		return
	}
	listNode := list.head
	for listNode.next != nil {
		listNode = listNode.next
	}
	listNode.next = node
}

// Prepend is a function for add new node into the begging of linked list
func (list *LinkedList) Prepend(node *Node) {
	if list.head == nil {
		list.head = node
		return
	}
	node.next = list.head
	list.head = node
}

// Delete is a function for delete node(s) from the linked list
func (list *LinkedList) Delete(searchValue int) {
	if list.head == nil {
		return
	}

	// Remove nodes from the begging of the list
	for list.head != nil && list.head.value == searchValue {
		list.head = list.head.next
	}

	listNode := list.head
	for listNode != nil {
		if listNode.next != nil && listNode.next.value == searchValue {
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

// todo : do not use Node as a parameter, use scalar data type and return Node

// todo : add Find function
