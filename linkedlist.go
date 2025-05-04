package datstr

import "errors"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func NewNode(val int) *Node {
	return &Node{value: val, next: nil}
}

func NewLinkedList() *LinkedList {
	n := NewNode(0)
	return &LinkedList{
		head: n,
		tail: n,
		len:  0,
	}
}

// Add value to the end of the linked list
func (l *LinkedList) Append(val int) error {
	n := NewNode(val)
	l.tail.next = n
	l.tail = n
	l.len++
	return nil
}

// Add value after the head node
func (l *LinkedList) AppendAtHead(val int) {
	n := NewNode(val)
	if l.len == 0 {
		l.Append(val)
		return
	}
	n.next = l.head.next
	l.head.next = n
	l.len++
}

// Remove value from the end of the list
func (l *LinkedList) Pop() error {
	if l.len == 0 {
		return errors.New("Len 0")
	}
	currNode := l.head
	for currNode.next != l.tail {
		currNode = currNode.next
	}
	currNode.next = nil
	l.tail = currNode
	l.len--
	return nil
}

// Remove value after the head
func (l *LinkedList) PopFromHead() error {
	if l.len == 0 {
		return errors.New("Len 0")
	}
	if l.len == 1 {
		l.tail = l.head
		l.head.next = nil
		l.len--
		return nil
	}
	l.head.next = l.head.next.next
	l.len--
	return nil
}

// Add value before the mentioned idx
func (l *LinkedList) Add(idx, val int) error {
	if idx > l.len || idx < 0 {
		return errors.New("Invalid Index")
	}
	if idx == 0 {
		l.AppendAtHead(val)
		return nil
	}
	head := l.head
	for range idx {
		head = head.next
	}
	n := NewNode(val)
	n.next = head.next
	head.next = n
	l.len++
	return nil
}

// Delete the node at the index
func (l *LinkedList) Delete(idx int) error {
	if idx >= l.len || idx < 0 {
		return errors.New("Invalid index")
	}
	if idx == 0 {
		return l.PopFromHead()
	} else if idx == l.len-1 {
		return l.Pop()
	}
	prev := l.head
	node := l.head.next
	for range idx {
		prev = node
		node = node.next
	}
	prev.next = node.next
	l.len--
	return nil
}

func (l *LinkedList) Len() int {
	return l.len
}
