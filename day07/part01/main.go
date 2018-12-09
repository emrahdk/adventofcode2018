package main

import (
	"fmt"
)

func main() {
	steps := readExample()

	for k, v := range steps {
		fmt.Printf("%v -> %v\n", k, v)
	}

	fmt.Println(len(steps))
	order := make([]string, 0)

	for k, v := range steps {
		index := getIndex(order, k)
		parent := 0

		if index == -1 {
			parent = 0
		}

		order[parent] = k

		for _, w := range v {
			index := getIndex(order, w)
			if index == -1 {
				order[parent+1] = k
			}
		}

	}

	fmt.Println(order)
}

func getIndex(order []string, letter string) int {
	for i, v := range order {
		if v == letter {
			return i
		}
	}

	return -1
}

func readExample() map[string][]string {
	lines := []string{
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
		"Step A must be finished before step B can begin.",
		"Step F must be finished before step E can begin.",
		"Step A must be finished before step D can begin."}

	steps := map[string][]string{}

	for _, v := range lines {
		name := ""
		blocks := ""
		fmt.Sscanf(v, "Step %v must be finished before step %v can begin.", &name, &blocks)

		if _, ok := steps[name]; ok {
			steps[name] = append(steps[name], blocks)
		} else {

			steps[name] = []string{blocks}
		}
	}

	return steps
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
