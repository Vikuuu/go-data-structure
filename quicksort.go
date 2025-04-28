package datstr

func partition(list *[]int, l, h int) int {
	pivot := h
	// do the sorting here
	i := l
	for j := l; j < h+1; j++ {
		if (*list)[j] < (*list)[pivot] {
			(*list)[j], (*list)[i] = (*list)[i], (*list)[j]
			i++
		}
	}
	(*list)[pivot], (*list)[i] = (*list)[i], (*list)[pivot]

	return i
}

func qsort(list *[]int, l, h int) {
	// Base condition
	if l >= h {
		return
	}

	p := partition(list, l, h)

	qsort(list, l, p-1)
	qsort(list, p+1, h)
}

func QuickSort(list *[]int) []int {
	l := list
	// fmt.Print("Length of l: ", len(*l))

	qsort(l, 0, len(*l)-1)

	return *l
}
