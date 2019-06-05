package graph

type Graph struct {
	vertexes map[int]Vertex
	edges map[string]Edge
}

func NewGraph(values ...int) *Graph {
	graph := &Graph{
		make(map[int]Vertex),
		make(map[string]Edge),
	}

	if len(values) > 0 {
		for _, value := range values {
			graph.AddVertex(value)
		}
	}
	return graph
}

func (g *Graph) AddVertex(value int) {
	g.vertexes[value] = Vertex{value}
}

func (g *Graph) RemoveVertex(value int) {
	delete(g.vertexes, value)
}
