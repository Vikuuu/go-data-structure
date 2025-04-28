package datstr

import (
	"errors"
)

type dnode struct {
	val  any
	next *dnode
	prev *dnode
}

type DoublyLinkedList struct {
	lenght int
	head   *dnode
	tail   *dnode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		lenght: 0,
		head:   nil,
		tail:   nil,
	}
}

func (d *DoublyLinkedList) Length() int { return d.lenght }

func (d *DoublyLinkedList) InsertAt(item any, idx int) error {
	if idx > d.lenght {
		return errors.New("Error")
	} else if idx == d.lenght {
		d.Append(item)
		return nil
	} else if idx == 0 {
		d.Prepend(item)
		return nil
	}

	d.lenght++
	curr := d.getAt(idx)
	n := &dnode{val: item}
	n.next = curr
	n.prev = curr.prev

	curr.prev.next = n
	curr.prev = n

	return nil
}

func (d *DoublyLinkedList) Remove(item any) any {
	curr := d.head
	for i := 0; curr != nil && i < d.lenght; i++ {
		if curr.val == item {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		return nil
	}

	return d.removeNode(curr)
}

func (d *DoublyLinkedList) RemoveAt(idx int) any {
	node := d.getAt(idx)

	if node == nil {
		return nil
	}

	return d.removeNode(node)
}

func (d *DoublyLinkedList) Append(item any) {
	n := &dnode{val: item}

	d.lenght++
	if d.tail == nil {
		d.tail = n
		d.head = n
		return
	}

	n.prev = d.tail
	d.tail.next = n
	d.tail = n
}

func (d *DoublyLinkedList) Prepend(item any) {
	n := &dnode{val: item, next: nil, prev: nil}

	d.lenght++
	if d.head == nil {
		d.head = n
		d.tail = n
		return
	}

	n.next = d.head
	d.head.prev = n
	d.head = n
}

func (d *DoublyLinkedList) Get(idx int) any {
	return d.getAt(idx).val
}

func (d *DoublyLinkedList) getAt(idx int) *dnode {
	curr := d.head
	for i := 0; curr.val != nil && i < idx; i++ {
		curr = curr.next
	}
	return curr
}

func (d *DoublyLinkedList) removeNode(n *dnode) any {
	d.lenght--
	if d.lenght == 0 {
		out := d.head.val
		d.head, d.tail = nil, nil
		return out
	}

	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	if n == d.head {
		d.head = n.next
	}
	if n == d.tail {
		d.tail = n.prev
	}

	n.prev, n.next = nil, nil
	return n.val
}
