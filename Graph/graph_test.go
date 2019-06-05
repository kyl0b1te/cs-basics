package graph

import "testing"

func TestNewGraph(t *testing.T) {
	graph := NewGraph()
	if graph.vertexes == nil {
		t.Error("`graph.vertexes` should not be nil")
	}
	if graph.edges == nil {
		t.Error("`graph.edges` should not be nil")
	}
}

func TestAddVertex(t *testing.T) {
	graph := NewGraph()

	test := func (value, expected int) {
		graph.AddVertex(value)
		if actual := len(graph.vertexes); expected != actual {
			t.Errorf("Expected %+v, actual %+v", expected, actual)
		}
	}

	for i, value := range [4]int{1, 5, 9, 3} {
		test(value, i + 1)
	}
	test(1, 4)
}
