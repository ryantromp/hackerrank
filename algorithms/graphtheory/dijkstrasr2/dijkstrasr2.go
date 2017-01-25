package main

import "fmt"

// Graph is a representation of a graph
type Graph struct {
	Nodes []*Node
}

// Node is a representation of a node in a graph
type Node struct {
	Distance int
	Previous *Node
	Edges    []*Edge
}

// Edge is a representation of an edge in a graph
type Edge struct {
	EndNode *Node
	Weight  int
}

// NewGraph creates a new Graph with the specified number of nodes
// and returns a reference to the graph
func NewGraph(numNodes int) *Graph {
	nodes := make([]*Node, numNodes)
	for i := 0; i < numNodes; i++ {
		node := Node{
			Distance: -1,
			Previous: nil,
			Edges:    make([]*Edge, 0),
		}
		nodes[i] = &node
	}

	g := Graph{Nodes: nodes}

	return &g
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(x, y, r int) {
	ex := Edge{EndNode: g.Nodes[y-1], Weight: r}
	ey := Edge{EndNode: g.Nodes[x-1], Weight: r}
	g.Nodes[x-1].Edges = append(g.Nodes[x-1].Edges, &ey)
	g.Nodes[y-1].Edges = append(g.Nodes[y-1].Edges, &ex)
}

// Implement dijkstra's algorithm here
func dijkstra(g *Graph, start int) {
	var nodes []*Node
	nodes.
}

func main() {

	// Get the number of test cases
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		processTestCase()
	}
}

func processTestCase() {

	// Get the number of nodes(n) and edges (m)
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	graph := NewGraph(n)

	// Process the input to get the edges
	for i := 0; i < m; i++ {
		var x, y, r int
		fmt.Scanf("%d %d %d\n", &x, &y, &r)
		graph.AddEdge(x, y, r)
	}

	// Get the starting node
	var s int
	fmt.Scanf("%d\n", &s)

}
