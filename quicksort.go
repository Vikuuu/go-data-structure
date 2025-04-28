package datstr

func partition(arr *[]int, l, h int) int {
	pivot := (*arr)[h]
	// do the sorting here
	idx := l - 1
	for i := l; i < h; i++ {
		if (*arr)[i] <= pivot {
			idx++
			tmp := (*arr)[i]
			(*arr)[i] = (*arr)[idx]
			(*arr)[idx] = tmp
		}
	}
	idx++
	(*arr)[h] = (*arr)[idx]
	(*arr)[idx] = pivot

	return idx
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
