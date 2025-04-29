package datstr

func inWalk(node *BNode, path []int) []int {
	if node == nil {
		return path
	}
	path = inWalk(node.Left, path)
	path = append(path, node.Val)
	path = inWalk(node.Right, path)

	return path
}

func InOrderSearch(head *BNode) []int {
	return inWalk(head, []int{})
}
