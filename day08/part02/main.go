package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(readInput()[0], " ")
	root := getNode(&input)

	fmt.Println(getMetadataSum(root))
}

func getMetadataSum(root node) int {
	if root.children == 0 {
		sum := 0
		for _, v := range root.metadata {
			sum += v
		}

		return sum
	}

	sum := 0
	for i := 0; i < len(root.metadata); i++ {
		index := root.metadata[i] - 1

		if index < len(root.childNodes) {
			sum += getMetadataSum(root.childNodes[index])
		}
	}

	return sum
}

func getNode(tree *[]string) node {
	root := node{
		children:      getInt((*tree)[0]),
		metadataCount: getInt((*tree)[1])}

	*tree = (*tree)[2:]

	nodes := make([]node, 0)

	for k := 0; k < root.children; k++ {
		nodes = append(nodes, getNode(tree))
	}

	metadata := make([]int, 0)
	for i := 0; i < root.metadataCount; i++ {
		metadata = append(metadata, getInt((*tree)[i]))
	}
	root.metadata = metadata
	*tree = (*tree)[root.metadataCount:]
	root.childNodes = nodes

	return root
}

func getInt(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

type node struct {
	children      int
	metadataCount int
	metadata      []int
	childNodes    []node
}

func readExample() string {
	return "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"
	// A: 2 3 - 1 1 2
	// B: 0 3 - 10 11 12
	// C: 1 1 - 2
	// D: 0 1 - 99
}

func readInput() []string {
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

	return lines
}
