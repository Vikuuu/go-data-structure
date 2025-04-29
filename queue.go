package datstr

type queueNode struct {
	val  any
	next *queueNode
}

type Queue struct {
	head   *queueNode
	tail   *queueNode
	length int
}

func NewQueue() *Queue {
	return &Queue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (q *Queue) Enqueue(val any) {
	q.length++
	n := &queueNode{val: val, next: nil}

	if q.tail == nil {
		q.tail = n
		q.head = n
		return
	}

	tail := q.tail
	tail.next = n
	q.tail = n
}

func (q *Queue) Deque() any {
	if q.head == nil {
		return nil
	}

	q.length--
	n := q.head
	q.head = n.next

	// free
	n.next = nil

	if q.length == 0 {
		q.tail = nil
	}

	return n.val
}

func (q *Queue) Peek() any {
	return q.head.val
}

func (q *Queue) Len() int {
	return q.length
}
