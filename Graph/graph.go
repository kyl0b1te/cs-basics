package graph

// Graph represents a corresponding abstract data structure
type Graph struct {
	vertexes map[int]Vertex
	edges map[string]Edge
}

// Vertex represents a graph node
type Vertex struct {
	Value int
	Edges []*Edge
}

// Vertex represents a graph relation
type Edge struct {
	Label string
	X, Y *Vertex
}

// NewGraph is a function that creates new "instance" for structure
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

// AddVertex is a function that creates new vertex with provided value
func (g *Graph) AddVertex(value int) {
	g.vertexes[value] = Vertex{value, []*Edge{}}
}

// RemoveVertex is a function that removes vertex by provided value
func (g *Graph) RemoveVertex(value int) {
	delete(g.vertexes, value)
}

// AddEdge is a function that set relation between 2 vertexes
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

// RemoveEdge is a function that removes previously set relation
func (g *Graph) RemoveEdge(label string) {
	edge, exist := g.edges[label]
	if !exist {
		// Edge is not exist
		return
	}
	delete(g.edges, label)
	g.removeVertexEdge(edge.X, edge)
	g.removeVertexEdge(edge.Y, edge)
}

// Neighbors is a function that returns list of vertex related values
func (g *Graph) Neighbors(value int) []int {
	vertex := g.getVertex(value)

	var neighbors []int
	for _, edge := range vertex.Edges {
		if edge.X.Value == value {
			neighbors = append(neighbors, edge.Y.Value)
			continue;
		}
		neighbors = append(neighbors, edge.X.Value)
	}
	return neighbors
}

func (g *Graph) getVertex(value int) *Vertex {
	vertex, ok := g.vertexes[value]
	if !ok {
		g.AddVertex(value)
		return g.getVertex(value)
	}
	return &vertex
}

func (g *Graph) removeVertexEdge(vertex *Vertex, edge Edge) {
	for i, vEdge := range vertex.Edges {
		if *vEdge == edge {
			vertex.Edges[i] = vertex.Edges[len(vertex.Edges) - 1] // Copy last element to target position
			vertex.Edges = vertex.Edges[:len(vertex.Edges)-1]
			break
		}
	}
	g.vertexes[vertex.Value] = *vertex
}
