package main

import (
	"fmt"
	"sort"
)

func main() {
	edges := readExample()

	for _, v := range edges {
		fmt.Printf("%v -> %v\n", v.from, v.to)
	}

	// Find all vertices
	vertices := map[string]bool{}
	for _, v := range edges {
		vertices[v.from] = true
		vertices[v.to] = true
	}

	// Find vertices with no incomming edges
	noIncommingEdges := make([]string, 0)
	for k, _ := range vertices {
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

	// SORT
	sort.Strings(noIncommingEdges)

	// BEGIN ALGO

	queue := make([]string, 0)
	queue = noIncommingEdges

	for len(edges) > 0 {
		for _, v := range queue {
			// Find edges from C
			verticeEdges := make([]string, 0)

			for i, w := range edges {
				if w.from == v {
					verticeEdges = append(verticeEdges, v)
					if len(edges) > 1 {
						edges = append(edges[:i], edges[i+1:]...)
					} else {
						edges = append(edges[:i])
					}
				}
			}

			// ADD TO QUEUE
			sort.Strings(verticeEdges)
		}
	}

}

func getIndex(order []string, letter string) int {
	for i, v := range order {
		if v == letter {
			return i
		}
	}

	return -1
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

// func readInput() map[string]string {
// 	lines := make([]string, 0)

// 	file, err := os.Open("../input/input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	coords := make([]coord, 0)

// 	for _, v := range lines {
// 		c := coord{}
// 		fmt.Sscanf(v, "%d, %d", &c.x, &c.y)
// 		coords = append(coords, c)
// 	}

// 	return coords
// }
