package datstr

func BFS(head *BNode, needle int) bool {
	queue := NewQueue[*BNode]()
	queue.Enqueue(head)

	for queue.Len() != 0 {
		next := queue.Deque()
		if next.Val == needle {
			return true
		}
		if next.Left != nil {
			queue.Enqueue(next.Left)
		}
		if next.Right != nil {
			queue.Enqueue(next.Right)
		}
	}
	return false
}
