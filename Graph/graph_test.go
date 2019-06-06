package graph

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, condition bool, err string) {
	if !condition {
		t.Error(err)
	}
}

func TestNewGraphWithoutValues(t *testing.T) {
	graph := NewGraph()
	if graph.vertexes == nil {
		t.Error("`graph.vertexes` should not be nil")
	}
	if graph.edges == nil {
		t.Error("`graph.edges` should not be nil")
	}
}

func TestNewGraphWithValues(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	graph := NewGraph(values...)

	if len(graph.vertexes) != len(values) {
		t.Errorf("`graph.vertexes` should be filled with %d values", len(values))
	}

	for _, value := range values {
		if _, ok := graph.vertexes[value]; ok != true {
			t.Errorf("Value %+v should be present in graph", value)
		}
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

func TestRemoveVertex(t *testing.T) {
	tests := []int{3, 5, 1, 8, 4}
	graph := NewGraph(tests...)

	test := func (value, expected int) {
		graph.RemoveVertex(value)
		if actual := len(graph.vertexes); expected != actual {
			t.Errorf("Expected %+v, actual %+v", expected, actual)
		}
	}

	i := len(tests)
	for _, value := range tests {
		i--
		test(value, i)
	}
}

func TestAddEdgeForExistingVertexes(t *testing.T) {
	graph := NewGraph(1, 2, 3, 4, 5)

	graph.AddEdge("t01", 1, 3)
	assert(t, len(graph.edges) == 1, "1 graph edge should be created")

	edge, exists := graph.edges["t01"]
	assert(t, exists, "`t01` edge label should exist")

	msg := fmt.Sprintf("Edge vertexes values should be `%d` and `%d`", 1, 3)
	assert(t, edge.X.Value == 1 && edge.Y.Value == 3, msg)

	msg = fmt.Sprintf("X Vertex should hold edge reference")
	assert(t, *graph.vertexes[1].Edges[0] == edge, msg)

	msg = fmt.Sprintf("Y Vertex should hold edge reference")
	assert(t, *graph.vertexes[3].Edges[0] == edge, msg)
}

func TestAddEdgeForNewVertexes(t *testing.T) {
	graph := NewGraph(1, 2, 3, 4, 5)

	graph.AddEdge("t01", 8, 9)

	edge, exists := graph.edges["t01"]
	assert(t, exists, "`t01` edge label should exist")

	msg := fmt.Sprintf("Edge vertexes values should be `%d` and `%d`", 8, 9)
	assert(t, edge.X.Value == 8 && edge.Y.Value == 9, msg)

	x, xOk := graph.vertexes[8];
	assert(t, xOk == true, "`X` vertex should be created automatically")
	assert(t, *x.Edges[0] == edge, "`X` vertex should should have edge")

	y, yOk := graph.vertexes[9];
	assert(t, yOk == true, "`Y` vertex should be created automatically")
	assert(t, *y.Edges[0] == edge, "`Y` vertex should should have edge")
}

func TestRemoveEdge(t *testing.T) {
	graph := NewGraph(4, 7, 2, 6, 1, 9)
	tests := map[string][]int{ "t01": []int{4, 2}, "t02": []int{7, 6}, "t03": []int{1, 9} }
	for label, values := range tests {
		graph.AddEdge(label, values[0], values[1])
	}

	for label, values := range tests {
		graph.RemoveEdge(label)

		_, exists := graph.edges[label]
		assert(t, !exists, fmt.Sprintf("`%s` edge label shouldn't exist", label))

		msg := fmt.Sprintf("Edges array should be empty, %+v", graph.vertexes[values[0]].Edges)
		assert(t, len(graph.vertexes[values[0]].Edges) == 0, msg)

		msg = fmt.Sprintf("Edges array should be empty, %+v", graph.vertexes[values[1]].Edges)
		assert(t, len(graph.vertexes[values[1]].Edges) == 0, msg)
	}
}
