package datstr

type queueNode[T any] struct {
	val  T
	next *queueNode[T]
}

type Queue[T any] struct {
	head   *queueNode[T]
	tail   *queueNode[T]
	length int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (q *Queue[T]) Enqueue(val T) {
	q.length++
	n := &queueNode[T]{val: val, next: nil}

	if q.tail == nil {
		q.tail = n
		q.head = n
		return
	}

	tail := q.tail
	tail.next = n
	q.tail = n
}

func (q *Queue[T]) Deque() T {
	if q.head == nil {
		var zero T
		return zero
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

func (q *Queue[T]) Peek() T {
	return q.head.val
}

func (q *Queue[T]) Len() int {
	return q.length
}
