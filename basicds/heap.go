package basicds

import "golang.org/x/exp/constraints"

func HeapSort[T constraints.Ordered](data []T) {
	HeapSortFunc(data, func(a, b T) bool {
		return a < b
	})
}

func HeapSortFunc[T any](data []T, less func(a, b T) bool) {
	h := Heap[T]{data[:0], less}
	for i := range data {
		h.Push(data[i])
	}
	for i := range data {
		data[len(data)-1-i] = h.PopMax()
	}
}

type Heap[T any] struct {
	Data []T
	Less func(a, b T) bool
}

func left(i int) int   { return 2*i + 1 }
func right(i int) int  { return 2*i + 2 }
func parent(i int) int { return (i - 1) / 2 }

func (h *Heap[T]) Len() int { return len(h.Data) }

func (h *Heap[T]) Push(val T) {
	h.Data = append(h.Data, val)
	h.maxHeapifyUp(len(h.Data) - 1)
}

func (h *Heap[T]) swap(a, b int) {
	h.Data[a], h.Data[b] = h.Data[b], h.Data[a]
}

func (h *Heap[T]) maxHeapifyUp(i int) {
	for i > 0 {
		j := parent(i)
		if h.Less(h.Data[j], h.Data[i]) {
			h.swap(i, j)
			i = j
		} else {
			return
		}
	}
}

func (h *Heap[T]) PopMax() T {
	v := h.Data[0]
	h.Data[0] = h.Data[len(h.Data)-1]
	h.Data = h.Data[:len(h.Data)-1]
	h.maxHeapifyDown(0)
	return v
}

func (h *Heap[T]) maxHeapifyDown(i int) {
	for i < len(h.Data) {
		left := left(i)
		right := right(i)
		if left >= len(h.Data) {
			return
		}
		if h.Less(h.Data[i], h.Data[left]) {
			if right < len(h.Data) && h.Less(h.Data[left], h.Data[right]) {
				h.swap(right, i)
				i = right
			} else {
				h.swap(left, i)
				i = left
			}
		} else {
			if right >= len(h.Data) {
				return
			}
			if h.Less(h.Data[i], h.Data[right]) {
				h.swap(right, i)
				i = right
			} else {
				return
			}
		}
	}
}
