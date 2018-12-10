package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	edges := readInput()

	// Find all vertices
	vertexMap := map[string]bool{}
	for _, v := range edges {
		vertexMap[v.from] = true
		vertexMap[v.to] = true
	}

	res := make([]string, 0)
	for len(edges) > 0 {
		vertex := getVerticesWithNoIncommingEdges(vertexMap, edges)
		delete(vertexMap, vertex) // Visited
		res = append(res, vertex)
		edges = removeEdge(edges, vertex)

	}

	// Add last non visited vertex
	for k := range vertexMap {
		res = append(res, k)
	}

	// Print res
	for _, v := range res {
		fmt.Printf(v)
	}
}

// Find vertices with no incomming edges
func getVerticesWithNoIncommingEdges(notVisitedVertices map[string]bool, edges []edge) string {
	noIncommingEdges := make([]string, 0)
	for k := range notVisitedVertices {
		found := false
		for _, w := range edges {
			if w.to == k {
				found = true
				break
			}
		}

		if !found {
			noIncommingEdges = append(noIncommingEdges, k)
		}
	}

	sort.Strings(noIncommingEdges)

	return noIncommingEdges[0]
}
func removeEdge(edges []edge, vertex string) []edge {
	for i := 0; i < len(edges); i++ {
		if edges[i].from == vertex {
			edges = append(edges[:i], edges[i+1:]...)
			i--
		}
	}

	return edges
}

func readExample() []edge {
	lines := []string{
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
		"Step A must be finished before step B can begin.",
		"Step F must be finished before step E can begin.",
		"Step A must be finished before step D can begin."}

	edges := make([]edge, 0)

	for _, v := range lines {
		edge := edge{}
		fmt.Sscanf(v, "Step %v must be finished before step %v can begin.", &edge.from, &edge.to)

		edges = append(edges, edge)
	}

	return edges
}

type edge struct {
	from string
	to   string
}

func readInput() []edge {
	lines := []string{}

	file, err := os.Open("../input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	edges := make([]edge, 0)

	for _, v := range lines {
		edge := edge{}
		fmt.Sscanf(v, "Step %v must be finished before step %v can begin.", &edge.from, &edge.to)

		edges = append(edges, edge)
	}

	return edges
}
