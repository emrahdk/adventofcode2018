package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	points := readExample()

	GRID_SIZE := 1000
	grid := [1000][1000]string{}
	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			grid[i][j] = "."
		}
	}

	for _, v := range points {
		x, y := convert(v.position.x, v.position.y, GRID_SIZE)
		grid[y][x] = "#"
	}

	// PrintGrid
	for _, v := range grid {
		fmt.Println(v)
	}

	for {
		// Rest grid
		for i := 0; i < GRID_SIZE; i++ {
			for j := 0; j < GRID_SIZE; j++ {
				grid[i][j] = "."
			}
		}

		for i := 0; i < len(points); i++ {
			current := points[i]
			newpoint := factorVelocity(current)
			points[i] = newpoint
			x, y := convert(newpoint.position.x, newpoint.position.y, GRID_SIZE)
			grid[y][x] = "#"
		}

		// PrintGrid
		for _, v := range grid {
			fmt.Println(v)
		}
		fmt.Println()

		time.Sleep(time.Second)
	}

}

func factorVelocity(p point) point {
	x := p.position.x + p.velocity.x
	y := p.position.y + p.velocity.y

	position := coord{x: x, y: y}

	return point{
		position: position,
		velocity: p.velocity}
}

func convert(x int, y int, grid int) (int, int) {
	middle := (grid / 2) - 1
	newX := x + middle
	newY := y + middle

	return newX, newY
}

type point struct {
	position coord
	velocity coord
}

type coord struct {
	x int
	y int
}

func readExample() []point {
	lines := []string{
		"position=< 9,  1> velocity=< 0,  2>",
		"position=< 7,  0> velocity=<-1,  0>",
		"position=< 3, -2> velocity=<-1,  1>",
		"position=< 6, 10> velocity=<-2, -1>",
		"position=< 2, -4> velocity=< 2,  2>",
		"position=<-6, 10> velocity=< 2, -2>",
		"position=< 1,  8> velocity=< 1, -1>",
		"position=< 1,  7> velocity=< 1,  0>",
		"position=<-3, 11> velocity=< 1, -2>",
		"position=< 7,  6> velocity=<-1, -1>",
		"position=<-2,  3> velocity=< 1,  0>",
		"position=<-4,  3> velocity=< 2,  0>",
		"position=<10, -3> velocity=<-1,  1>",
		"position=< 5, 11> velocity=< 1, -2>",
		"position=< 4,  7> velocity=< 0, -1>",
		"position=< 8, -2> velocity=< 0,  1>",
		"position=<15,  0> velocity=<-2,  0>",
		"position=< 1,  6> velocity=< 1,  0>",
		"position=< 8,  9> velocity=< 0, -1>",
		"position=< 3,  3> velocity=<-1,  1>",
		"position=< 0,  5> velocity=< 0, -1>",
		"position=<-2,  2> velocity=< 2,  0>",
		"position=< 5, -2> velocity=< 1,  2>",
		"position=< 1,  4> velocity=< 2,  1>",
		"position=<-2,  7> velocity=< 2, -2>",
		"position=< 3,  6> velocity=<-1, -1>",
		"position=< 5,  0> velocity=< 1,  0>",
		"position=<-6,  0> velocity=< 2,  0>",
		"position=< 5,  9> velocity=< 1, -2>",
		"position=<14,  7> velocity=<-2,  0>",
		"position=<-3,  6> velocity=< 2, -1>"}

	points := make([]point, 0)

	for _, v := range lines {
		point := point{
			position: coord{},
			velocity: coord{}}

		fmt.Sscanf(v, "position=<%d, %d> velocity=<%d, %d>", &point.position.x, &point.position.y, &point.velocity.x, &point.velocity.y)

		points = append(points, point)
	}

	return points
}

func readInput() []point {
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

	points := make([]point, 0)

	for _, v := range lines {
		point := point{
			position: coord{},
			velocity: coord{}}

		fmt.Sscanf(v, "position=<%d, %d> velocity=<%d, %d>", &point.position.x, &point.position.y, &point.velocity.x, &point.velocity.y)

		points = append(points, point)
	}

	return points
}
