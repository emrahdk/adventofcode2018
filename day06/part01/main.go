package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	// coords := []coord{
	// 	coord{name: "A", x: 1, y: 1},
	// 	coord{name: "B", x: 1, y: 6},
	// 	coord{name: "C", x: 8, y: 3},
	// 	coord{name: "D", x: 3, y: 4},
	// 	coord{name: "E", x: 5, y: 5},
	// 	coord{name: "F", x: 8, y: 9}}

	const HMAX int = 8 + 1
	const VMAX int = 9 + 1

	rect := [10][10]grid{}

	// for _, v := range coords {
	v := coord{name: "A", x: 1, y: 1}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if v.y == i && v.x == j {
				// START POS
				rect[i][j] = grid{v.name, -1}
			} else {
				// CHECK IF CURRENT HAS LOWER OR EQUAL (.) DISTANCE
				// CALC DISTANCE
				rect[i][j] = grid{strings.ToLower(v.name), manhattanDist(v, coord{x: j, y: i})}
			}
		}
	}
	// }

	for _, v := range rect {
		fmt.Println(v)
	}

	// manhattanDist(coord{x: 1, y: 1}, coord{x: 1, y: 6})

}

func manhattanDist(a coord, b coord) int {
	return int(math.Abs(float64(b.x-a.x) + math.Abs(float64(b.y-a.y))))
}

type coord struct {
	name string
	x    int
	y    int
}

type grid struct {
	name string
	dist int
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
