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

func TestMinHeap(t *testing.T) {
	heap := ds.NewMinHeap()
	assert.Equal(t, 0, heap.Len())

	heap.Add(5)
	heap.Add(3)
	heap.Add(69)
	heap.Add(420)
	heap.Add(4)
	heap.Add(1)
	heap.Add(8)
	heap.Add(7)

	assert.Equal(t, 8, heap.Len())

	assert.Equal(t, 1, heap.Remove())
	assert.Equal(t, 3, heap.Remove())
	assert.Equal(t, 4, heap.Remove())
	assert.Equal(t, 5, heap.Remove())

	assert.Equal(t, 4, heap.Len())

	assert.Equal(t, 7, heap.Remove())
	assert.Equal(t, 8, heap.Remove())
	assert.Equal(t, 69, heap.Remove())
	assert.Equal(t, 420, heap.Remove())

	assert.Equal(t, 0, heap.Len())
}

func TestBFSGraphMatrix(t *testing.T) {
	assert.Equal(t, []int{0, 1, 4, 5, 6}, ds.BFSGraphMatrix(matrix2, 0, 6))
	assert.Equal(t, []int{}, ds.BFSGraphMatrix(matrix2, 6, 0))
}

func TestDFSGraphList(t *testing.T) {
	assert.Equal(t, []int{0, 1, 4, 5, 6}, ds.DFSGraphList(list2, 0, 6))
	assert.Equal(t, []int{}, ds.DFSGraphList(list2, 6, 0))
}

func TestDijkstraList(t *testing.T) {
	assert.Equal(t, []int{0, 1, 4, 5, 6}, ds.DijkstraList(0, 6, list1))
}

func TestLRU(t *testing.T) {
	lru := ds.NewLRU[string, int](3)

	assert.Nil(t, lru.Get("foo"))
	lru.Update("foo", 69)
	assert.Equal(t, 69, lru.Get("foo"))

	lru.Update("bar", 420)
	assert.Equal(t, 420, lru.Get("bar"))

	lru.Update("baz", 1337)
	assert.Equal(t, 1337, lru.Get("baz"))

	lru.Update("ball", 69420)
	assert.Equal(t, 69420, lru.Get("ball"))
	assert.Equal(t, nil, lru.Get("foo"))
	assert.Equal(t, 420, lru.Get("bar"))

	lru.Update("foo", 69)
	assert.Equal(t, 420, lru.Get("bar"))
	assert.Equal(t, 69, lru.Get("foo"))
}
