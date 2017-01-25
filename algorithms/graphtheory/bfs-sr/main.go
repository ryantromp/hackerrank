package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Node is a representation of a node in a graph
type Node struct {
	Value    int
	Edges    []*Edge
	Distance int
	Parent   *Node
}

// Graph is a representation of a graph
type Graph struct {
	Nodes []*Node
}

// Edge is a representation of an edge in a graph
type Edge struct {
	EndNode *Node
	Weight  int
}

// Queue is a LIFO list of Nodes
type Queue struct {
	Nodes []*Node
	Size  int
	Count int
	Head  int
	Tail  int
}

var solutions []string

// NewQueue creates a new Queue and returns it
func NewQueue() *Queue {
	return &Queue{Nodes: make([]*Node, 0)}
}

// Enqueue adds a new element to the queue
func (q *Queue) Enqueue(n *Node) {
	q.Nodes = append(q.Nodes, n)
	q.Tail++
	q.Count++
}

// Dequeue removes an element from the queue
func (q *Queue) Dequeue() *Node {
	if q.Count == 0 {
		return nil
	}

	node := q.Nodes[q.Head]
	q.Head++
	q.Count--

	return node
}

func main() {

	var q int
	fmt.Scanf("%d\n", &q)

	solutions = make([]string, 0)
	for i := 0; i < q; i++ {
		processQueryInput()
	}

	for _, s := range solutions {
		fmt.Println(s)
	}

}

func processQueryInput() {

	var graph = Graph{
		Nodes: make([]*Node, 0),
	}

	// Get the number of nodes and the number of edges
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	// Populate the (unlinked) nodes in the graph
	for i := 0; i < n; i++ {
		graph.Nodes = append(graph.Nodes, &Node{
			Value:    i + 1,
			Distance: -1,
			Parent:   nil,
			Edges:    make([]*Edge, 0),
		})
	}

	// For each edge, get the two nodes that are connected
	for j := 0; j < m; j++ {
		var u, v int

		fmt.Scanf("%d %d\n", &u, &v)
		eu := Edge{EndNode: graph.Nodes[v-1]}
		ev := Edge{EndNode: graph.Nodes[u-1]}
		graph.Nodes[u-1].Edges = append(graph.Nodes[u-1].Edges, &eu)
		graph.Nodes[v-1].Edges = append(graph.Nodes[v-1].Edges, &ev)
	}

	// Get the starting node
	var s int
	fmt.Scanf("%d\n", &s)

	q := NewQueue()
	graph.Nodes[s-1].Distance = 0
	q.Enqueue(graph.Nodes[s-1])

	var current *Node
	for q.Count > 0 {
		current = q.Dequeue()
		for _, e := range current.Edges {
			n := e.EndNode
			if n.Distance == -1 {
				n.Distance = current.Distance + 6
				n.Parent = current
				q.Enqueue(n)
			}
		}
	}

	var str string
	for i := 0; i < n; i++ {
		if i != (s - 1) {
			node := graph.Nodes[i]
			str = str + strconv.Itoa(node.Distance) + " "
		}
	}
	solutions = append(solutions, strings.TrimSpace(str))

	//fmt.Printf("%#v\n", graph)
}
