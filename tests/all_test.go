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
	q := ds.NewQueue[int]()
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
	assert.Equal(t, 0, q.Deque())
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

func TestMazeSolver(t *testing.T) {
	maze := [6]string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeResult := []ds.Point{
		{X: 10, Y: 0},
		{X: 10, Y: 1},
		{X: 10, Y: 2},
		{X: 10, Y: 3},
		{X: 10, Y: 4},
		{X: 9, Y: 4},
		{X: 8, Y: 4},
		{X: 7, Y: 4},
		{X: 6, Y: 4},
		{X: 5, Y: 4},
		{X: 4, Y: 4},
		{X: 3, Y: 4},
		{X: 2, Y: 4},
		{X: 1, Y: 4},
		{X: 1, Y: 5},
	}

	result := ds.MazeSolver(maze[:], "x", ds.Point{X: 10, Y: 0}, ds.Point{X: 1, Y: 5})
	assert.Equal(t, mazeResult, result)
}

func TestQuickSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	out := ds.QuickSort(&arr)
	assert.Equal(t, out, []int{3, 4, 7, 9, 42, 69, 420})
}

func ListTest(t *testing.T, list *ds.DoublyLinkedList) {
	list.Append(5)
	list.Append(7)
	list.Append(9)

	assert.Equal(t, 9, list.Get(2))
	assert.Equal(t, 7, list.RemoveAt(1))
	assert.Equal(t, 2, list.Length())

	list.Append(11)
	assert.Equal(t, 9, list.RemoveAt(1))
	assert.Nil(t, list.Remove(9))
	assert.Equal(t, 5, list.RemoveAt(0))
	assert.Equal(t, 11, list.RemoveAt(0))
	assert.Equal(t, 0, list.Length())

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	assert.Equal(t, 5, list.Get(2))
	assert.Equal(t, 9, list.Get(0))
	assert.Equal(t, 9, list.Remove(9))
	assert.Equal(t, 2, list.Length())
	assert.Equal(t, 7, list.Get(0))
}

func TestDoublyLinkedList(t *testing.T) {
	list := ds.NewDoublyLinkedList()
	ListTest(t, list)
}

func TestBTTraversal(t *testing.T) {
	assert.Equal(t,
		[]int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100},
		ds.InOrderSearch(tree),
	)
	assert.Equal(t,
		[]int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100},
		ds.PreOrderSearch(tree),
	)
	assert.Equal(t,
		[]int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20},
		ds.PostOrderSearch(tree),
	)
}

func TestBTBFS(t *testing.T) {
	assert.True(t, ds.BFS(tree, 45))
	assert.True(t, ds.BFS(tree, 7))
	assert.False(t, ds.BFS(tree, 69))
}

func TestCompareBT(t *testing.T) {
	assert.True(t, ds.CompareBT(tree, tree))
	assert.False(t, ds.CompareBT(tree, tree2))
}

func TestDFSonBST(t *testing.T) {
	assert.True(t, ds.DFSonBST(tree, 45))
	assert.True(t, ds.DFSonBST(tree, 7))
	assert.False(t, ds.DFSonBST(tree, 69))
}
