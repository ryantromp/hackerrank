package main

// Graph is a representation of a graph
type Graph struct {
	Nodes []*Node
}

// Node is a representation of a node in a graph
type Node struct {
	Value int
	Edges []*Edge
}

// Edge is a representation of an edge in a graph
type Edge struct {
	EndNode *Node
	Weight  int
}

func main() {

}
