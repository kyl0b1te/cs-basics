package graph

type Graph struct {
	vertexes map[int]Vertex
	edges map[string]Edge
}

func NewGraph() *Graph {
	return &Graph{
		make(map[int]Vertex),
		make(map[string]Edge),
	}
}

func (g *Graph) AddVertex(value int) {
	g.vertexes[value] = Vertex{value}
}
