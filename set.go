package datstr

type set struct {
	data map[string]struct{}
	len  int
}

func NewSet() *set {
	return &set{
		data: make(map[string]struct{}),
		len:  0,
	}
}

func (s *set) Add(val string) {
	if _, found := s.data[val]; !found {
		s.len++
	}
	s.data[val] = struct{}{}
}

func (s *set) Delete(val string) {
	if _, found := s.data[val]; found {
		s.len--
	}
	delete(s.data, val)
}

func (s *set) Find(val string) bool {
	if _, found := s.data[val]; found {
		return true
	}
	return false
}
