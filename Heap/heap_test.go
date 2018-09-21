package heap

import (
	"testing"
)

func TestNewHeap(t *testing.T) {
	heap := NewHeap()
	if len(heap.values) != 0 {
		t.Error("Stack should not be empty")
	}
}

func TestAddOnEmptyHeap(t *testing.T) {
	heap := NewHeap()
	heap.Add(10)
	if heap.values[0] != 10 {
		t.Errorf("Expected '%+v' actual '%+v'", 10, heap.values[0])
	}
}

func TestAdd(t *testing.T) {
	heap := NewHeap()
	tests := [3]int{2, 3, 5}

	for _, val := range tests {
		heap.Add(val)
	}

	for i, val := range tests {
		if heap.values[i] != val {
			t.Errorf("Expected '%+v' actual '%+v'", val, heap.values[i])
		}
	}

	heap.Add(1)
	if heap.values[0] != 1 {
		t.Errorf("Expected '%+v' actual '%+v'", 1, heap.values[0])
	}
}

func TestIsIndexValidOnEmptyHeap(t *testing.T) {
	heap := NewHeap()
	if actual := heap.isIndexValid(0); actual {
		t.Errorf("Expected '%+v' actual '%+v'", false, actual)
	}
}

func TestIsIndexValidOnFilledHeap(t *testing.T) {
	heap := NewHeap()
	test := [3]int{1, 2, 5}
	for _, val := range test {
		heap.Add(val)
	}
	for i := range test {
		if actual := heap.isIndexValid(i); !actual {
			t.Errorf("Expected '%+v' actual '%+v'", true, actual)
		}
	}
}

func TestSwap(t *testing.T) {
	heap := NewHeap()
	for _, val := range [3]int{1, 2, 5} {
		heap.Add(val)
	}
	heap.swap(2, 1)
	for i, val := range [3]int{1, 5, 2} {
		if heap.values[i] != val {
			t.Errorf("Expected '%+v' actual '%+v'", val, heap.values[i])
		}
	}
}

func TestGetParentIndex(t *testing.T) {
	test := [3]int{-1, 0, 0}
	for i, val := range test {
		if actual := getParentIndex(i); actual != val {
			t.Errorf("Expected '%+v' actual '%+v'", val, actual)
		}
	}
}

func TestGetLeftIndex(t *testing.T) {
	if actual := getLeftIndex(1); actual != 3 {
		t.Errorf("Expected '%+v' actual '%+v'", 3, actual)
	}
}

func TestGetRightIndex(t *testing.T) {
	if actual := getRightIndex(1); actual != 4 {
		t.Errorf("Expected '%+v' actual '%+v'", 4, actual)
	}
}
