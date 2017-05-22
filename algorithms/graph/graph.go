package main

import "fmt"

// graph demonstrates a simple directed graph data structure in go

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("A", "B")
	var b = hasEdge("A", "B")
	fmt.Println(graph, b)
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
