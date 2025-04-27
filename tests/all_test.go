package tests

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	ds "github.com/Vikuuu/go-datastructure"
)

func TestBinarySearch(t *testing.T) {}

func TestTwoCrystalBalls(t *testing.T) {
	idx := rand.Intn(10000)
	data := make([]bool, 10000)

	assert.Equal(t, ds.TwoCrystalBalls(data), -1)
	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	assert.Equal(t, ds.TwoCrystalBalls(data), idx)
}

func TestQueue(t *testing.T) {
	q := ds.NewQueue()
	q.Enqueue(5)
	q.Enqueue(7)
	q.Enqueue(9)
	assert.Equal(t, q.Deque(), 5)
	assert.Equal(t, q.Len(), 2)

	q.Enqueue(11)
	assert.Equal(t, q.Deque(), 7)
	assert.Equal(t, q.Deque(), 9)
	assert.Equal(t, q.Peek(), 11)
	assert.Equal(t, q.Deque(), 11)
	assert.Nil(t, q.Deque())
	assert.Equal(t, q.Len(), 0)
}

func TestStack(t *testing.T) {
	s := ds.NewStack()

	s.Push(5)
	s.Push(7)
	s.Push(9)

	assert.Equal(t, s.Pop(), 9)
	assert.Equal(t, s.Len(), 2)

	s.Push(11)

	assert.Equal(t, s.Pop(), 11)
	assert.Equal(t, s.Pop(), 7)
	assert.Equal(t, s.Peek(), 5)
	assert.Equal(t, s.Pop(), 5)
	assert.Nil(t, s.Pop())

	s.Push(69)
	assert.Equal(t, s.Peek(), 69)
	assert.Equal(t, s.Len(), 1)
}
