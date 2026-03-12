package main

type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

// less(a, b) should return true if a should come before b.
// For a min-heap: less = func(a, b T) bool { return a < b }
// For a max-heap: less = func(a, b T) bool { return a > b }
func NewHeap[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data: []T{},
		less: less,
	}
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Peek() T {
	return h.data[0]
}

func (h *Heap[T]) Push(val T) {
	h.data = append(h.data, val)
	h.up(len(h.data) - 1)
}

func (h *Heap[T]) Pop() T {
	n := len(h.data)
	top := h.data[0]

	last := h.data[n-1]
	h.data = h.data[:n-1]

	if len(h.data) > 0 {
		h.data[0] = last
		h.down(0)
	}

	return top
}

func (h *Heap[T]) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if !h.less(h.data[i], h.data[parent]) {
			break
		}
		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

func (h *Heap[T]) down(i int) {
	n := len(h.data)
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.less(h.data[left], h.data[smallest]) {
			smallest = left
		}
		if right < n && h.less(h.data[right], h.data[smallest]) {
			smallest = right
		}
		if smallest == i {
			break
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}
