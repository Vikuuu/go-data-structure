package datstr

func search(curr *BNode, needle int) bool {
	if curr == nil {
		return false
	}

	if curr.Val == needle {
		return true
	}

	if curr.Val < needle {
		return search(curr.Right, needle)
	}
	return search(curr.Left, needle)
}

func DFSonBST(head *BNode, needle int) bool {
	return search(head, needle)
}
