package datstr

type stackNode struct {
	val  any
	prev *stackNode
}
type Stack struct {
	head   *stackNode
	lenght int
}

func NewStack() *Stack {
	return &Stack{
		head:   nil,
		lenght: 0,
	}
}

func (s *Stack) Push(val any) {
	s.lenght++
	n := &stackNode{val: val, prev: nil}
	if s.head == nil {
		s.head = n
		return
	}

	n.prev = s.head
	s.head = n
}

func (s *Stack) Pop() any {
	if s.lenght == 0 {
		return nil
	}
	n := s.head
	s.lenght--
	if s.lenght == 0 {
		s.head = nil
		return n.val
	}
	s.head = n.prev
	// free
	n.prev = nil

	return n.val
}

func (s *Stack) Peek() any {
	return s.head.val
}

func (s *Stack) Len() int {
	return s.lenght
}
