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

	// Requirements
	availableWorkers := 5
	workers := map[string]int{}
	stepTime := 60

	time := 0
	for {
		verticesWithNoIncommingEdges := getVerticesWithNoIncommingEdges(vertexMap, edges)

		if len(verticesWithNoIncommingEdges) == 0 {
			break
		}

		// Remove currently processing
		reduced := make([]string, 0)
		for _, v := range verticesWithNoIncommingEdges {
			if _, ok := workers[v]; !ok {
				reduced = append(reduced, v)
			}
		}

		verticesWithNoIncommingEdges = reduced

		for _, v := range verticesWithNoIncommingEdges {
			time := int(v[0] - 64 + byte(stepTime))

			if len(workers) >= availableWorkers {
				continue
			}

			workers[v] = time

		}

		time++
		edges = decrementWorkersUpdateGraph(workers, edges, vertexMap)
	}

	fmt.Println(time)
}

func decrementWorkersUpdateGraph(workers map[string]int, edges []edge, vertexMap map[string]bool) []edge {
	for k, v := range workers {
		if v-1 < 1 {
			delete(workers, k)
			delete(vertexMap, k)         // Visited
			edges = removeEdge(edges, k) // return edges
			continue
		}

		workers[k]--
	}

	return edges
}

// Find vertices with no incomming edges
func getVerticesWithNoIncommingEdges(notVisitedVertices map[string]bool, edges []edge) []string {
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

	return noIncommingEdges
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
