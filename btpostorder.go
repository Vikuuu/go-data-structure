package datstr

func postWalk(node *BNode, path []int) []int {
	if node == nil {
		return path
	}
	path = postWalk(node.Left, path)
	path = postWalk(node.Right, path)
	path = append(path, node.Val)

	return path
}

func PostOrderSearch(head *BNode) []int {
	return postWalk(head, []int{})
}
