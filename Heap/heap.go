package heap

import (
	"math"
)

// Heap represents abstract data structure
type Heap struct {
	values []int
}

// NewHeap is a function that creates a new heap
func NewHeap() *Heap {
	return &Heap{values: []int{}}
}

// Add is a function that add a new element into the heap
func (h *Heap) Add(val int) {
	h.values = append(h.values, val)

	lastIndex := len(h.values) - 1
	parentIndex := getParentIndex(lastIndex)
	for parentIndex >= 0 {
		if h.values[parentIndex] > h.values[lastIndex] {
			h.swap(parentIndex, lastIndex)
			lastIndex = parentIndex
		}
		parentIndex = getParentIndex(parentIndex)
	}
}

func (h *Heap) swap(a, b int) {
	c := h.values[a]
	h.values[a] = h.values[b]
	h.values[b] = c
}

func (h Heap) isIndexValid(idx int) bool {
	return idx >= 0 && idx < len(h.values)
}

func getParentIndex(idx int) int {
	if idx == 0 {
		return -1
	}
	return int(math.Floor(float64(idx)-1.0) / 2.0)
}

func getLeftIndex(idx int) int {
	return (2 * idx) + 1
}

func getRightIndex(idx int) int {
	return (2 * idx) + 2
}
