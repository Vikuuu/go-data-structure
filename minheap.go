package datstr

import "log"

type minHeap struct {
	len  int
	heap []int
}

func NewMinHeap() *minHeap {
	return &minHeap{
		len:  0,
		heap: []int{},
	}
}

func (h *minHeap) Add(val int) {
	log.Printf("Add %d to heap\n", val)
	h.heap = append(h.heap, val)
	log.Println("Called HeapifyUp")
	h.heapifyUp(h.len)
	h.len++
}

func (h *minHeap) Remove() int {
	if h.len == 0 {
		return -1
	}
	out := h.heap[0]
	h.len--

	if h.len == 0 {
		h.heap = []int{}
		return out
	}

	h.heap[0] = h.heap[h.len]
	h.heap = h.heap[:h.len]
	log.Printf("The heap: %v\n", h.heap)
	log.Println("Heapifying Down")
	h.heapifyDown(0)
	log.Printf("After Heapifing: %v\n", h.heap)
	return out
}

func (h *minHeap) heapifyUp(idx int) {
	log.Printf("Heap : %v\n", h.heap)
	log.Printf("Idx : %d\n", idx)
	if idx == 0 {
		return
	}
	p := parent(idx)
	log.Printf("Parend Idx: %d\n", p)
	parentV := h.heap[p]
	v := h.heap[idx]

	if parentV > v {
		log.Printf("Heap h.heap[idx]: %d, Parent Val: %d\n", h.heap[idx], parentV)
		h.heap[idx] = parentV
		h.heap[p] = v
		h.heapifyUp(p)
	}
}

func (h *minHeap) heapifyDown(idx int) {
	log.Printf("Heap in Heapify Down: %v\n", h)
	log.Printf("The idx in Heapify Down: %d\n", idx)
	lIdx := leftChild(idx)
	rIdx := rightChild(idx)
	log.Printf("Left Idx: %d Right Idx: %d\n", lIdx, rIdx)
	if idx >= h.len || rIdx >= h.len {
		return
	}

	v := h.heap[idx]
	lV := h.heap[lIdx]
	rV := h.heap[rIdx]

	if lV > rV && v > rV {
		h.heap[idx] = rV
		h.heap[rIdx] = v
		h.heapifyDown(rIdx)
	} else if rV > lV && v > lV {
		h.heap[idx] = lV
		h.heap[lIdx] = v
		h.heapifyDown(lIdx)
	}
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return 2*idx + 1
}

func rightChild(idx int) int {
	return 2*idx + 2
}

func (h *minHeap) Len() int {
	return h.len
}
