package datstr

func cmpWalk(n1, n2 *BNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	if n1.Val != n2.Val {
		return false
	}

	return cmpWalk(n1.Left, n2.Left) && cmpWalk(n1.Right, n2.Right)
}

func CompareBT(a, b *BNode) bool {
	return cmpWalk(a, b)
}
