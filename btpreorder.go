package datstr

type BNode struct {
	Val   int
	Left  *BNode
	Right *BNode
}

func preWalk(node *BNode, path []int) []int {
	if node == nil {
		return path
	}
	path = append(path, node.Val)
	path = preWalk(node.Left, path)
	path = preWalk(node.Right, path)

	return path
}

func PreOrderSearch(head *BNode) []int {
	return preWalk(head, []int{})
}
