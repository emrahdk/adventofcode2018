package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	overlapping := 0
	claims := []claim{}
	fabric := [1000][1000]string{}

	// Initialize matrix with default values
	for i, v := range fabric {
		for j, _ := range v {
			fabric[i][j] = "."
		}
	}

	// Get claims
	for _, v := range readInput() {
		claims = append(claims, getClaim(v))
	}

	for _, v := range claims {
		for i := v.row; i < v.row+v.height; i++ {
			for j := v.col; j < v.col+v.width; j++ {
				if fabric[i][j] == "." {
					fabric[i][j] = strconv.Itoa(v.id)
				} else {
					fabric[i][j] = "X"
				}
			}
		}
	}

	for _, v := range fabric {
		for _, w := range v {
			if w == "X" {
				overlapping++
			}
		}
	}

	fmt.Println(overlapping)
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

func getClaim(input string) claim {
	c := claim{}
	fmt.Sscanf(input, "#%d @ %d,%d: %dx%d", &c.id, &c.col, &c.row, &c.width, &c.height)
	return c
}

type claim struct {
	id     int
	row    int
	col    int
	width  int
	height int
}
