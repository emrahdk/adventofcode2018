package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	coords := readInput()

	const HMAX int = 300
	const VMAX int = 300
	const MaxDist int = 10000
	const SafeArea string = "#"

	rect := [HMAX][VMAX]string{}

	for i := 0; i < VMAX; i++ {
		for j := 0; j < HMAX; j++ {
			current := rect[i][j]
			if current == SafeArea {
				continue
			}

			dist := 0

			for _, v := range coords {
				if dist > MaxDist {
					break
				}
				dist += manhattanDist(coord{x: j, y: i}, v)
			}

			if dist < MaxDist {
				rect[i][j] = SafeArea
			}
		}
	}

	safeRegionSize := 0
	for i := 0; i < VMAX; i++ {
		for j := 0; j < HMAX; j++ {
			if rect[i][j] == SafeArea {
				safeRegionSize++
			}
		}
	}

	fmt.Println(safeRegionSize)

}

func manhattanDist(a coord, b coord) int {
	return int(math.Abs(float64(b.x-a.x)) + math.Abs(float64(b.y-a.y)))
}

type coord struct {
	x int
	y int
}

func readInput() []coord {
	lines := make([]string, 0)

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

	coords := make([]coord, 0)

	for _, v := range lines {
		c := coord{}
		fmt.Sscanf(v, "%d, %d", &c.x, &c.y)
		coords = append(coords, c)
	}

	return coords
}
