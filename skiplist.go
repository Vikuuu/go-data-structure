package datstr

import (
	"bytes"
	"errors"
	"math"
	_ "unsafe"
)

const (
	MaxHeight = 16
	PValue    = 0.5 // p = 1/2
)

var probabilities [MaxHeight]uint32

type node struct {
	key   []byte
	value []byte
	tower [MaxHeight]*node
}

type SkipList struct {
	head   *node
	height int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head: &node{}, height: 1,
	}
}

func (s *SkipList) search(key []byte) (*node, [MaxHeight]*node) {
	var next *node
	var journey [MaxHeight]*node

	prev := s.head
	for lvl := s.height - 1; lvl >= 0; lvl-- {
		for next = prev.tower[lvl]; next != nil; next = prev.tower[lvl] {
			if bytes.Compare(key, next.key) <= 0 {
				break
			}
			prev = next
		}
		journey[lvl] = prev
	}

	if next != nil && bytes.Equal(key, next.key) {
		return next, journey
	}
	return nil, journey
}

func (s *SkipList) Find(key []byte) ([]byte, error) {
	found, _ := s.search(key)
	if found == nil {
		return nil, errors.New("key not found")
	}
	return found.value, nil
}

func init() {
	probability := 1.0

	for lvl := 0; lvl < MaxHeight; lvl++ {
		probabilities[lvl] = uint32(probability * float64(math.MaxUint32))
		probability *= PValue
	}
}

func randomHeight() int {
	seed := Uint32()
	height := 1
	for height < MaxHeight && seed <= probabilities[height] {
		height++
	}
	return height
}

func (s *SkipList) Insert(key []byte, val []byte) {
	found, journey := s.search(key)

	if found != nil {
		// update value of existing key
		found.value = val
		return
	}
	height := randomHeight()
	n := &node{key: key, value: val}

	for lvl := 0; lvl < height; lvl++ {
		prev := journey[lvl]

		if prev == nil {
			// prev is nil if we are extending the height of the tree
			// because that level did not exist while the journey was
			// being recorded
			prev = s.head
		}
		n.tower[lvl] = prev.tower[lvl]
		prev.tower[lvl] = n
	}

	if height > s.height {
		s.height = height
	}
}

func (s *SkipList) Delete(key []byte) bool {
	found, journey := s.search(key)
	if found == nil {
		return false
	}

	for lvl := 0; lvl < s.height; lvl++ {
		if journey[lvl].tower[lvl] != found {
			break
		}
		journey[lvl].tower[lvl] = found.tower[lvl]
		found.tower[lvl] = nil
	}
	found = nil
	s.shrink()
	return true
}

func (s *SkipList) shrink() {
	for lvl := s.height - 1; lvl >= 0; lvl-- {
		if s.head.tower[lvl] == nil {
			s.height--
		}
	}
}

//go:linkname Uint32 runtime.fastrand
func Uint32() uint32
