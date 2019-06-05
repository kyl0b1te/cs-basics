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
	g.vertexes[value] = Vertex{value, []*Edge{}}
}

func (g *Graph) RemoveVertex(value int) {
	delete(g.vertexes, value)
}

func (g *Graph) AddEdge(label string, x, y int) {
	xVertex := g.getVertex(x)
	yVertex := g.getVertex(y)

	if _, exists := g.edges[label]; exists == true {
		// Edge is already exist
		return
	}
	edge := Edge{label, xVertex, yVertex}

	g.edges[label] = edge
	xVertex.Edges = append(xVertex.Edges, &edge)
	yVertex.Edges = append(yVertex.Edges, &edge)

	g.vertexes[x] = *xVertex
	g.vertexes[y] = *yVertex
}

func (g *Graph) getVertex(value int) *Vertex {
	vertex, ok := g.vertexes[value]
	if !ok {
		g.AddVertex(value)
		return g.getVertex(value)
	}
	return &vertex
}
