package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	points := readInput()

	var maxX, maxY, gridX, gridY int
	second := 0

	for {
		currentMaxX, currentMaxY, currentGridX, currentGridY := getGridData(points)

		// dry run
		copy := append(points[:0:0], points...)
		for j := 0; j < len(copy); j++ {
			current := copy[j]
			newpoint := factorVelocity(current)
			copy[j] = newpoint
		}

		_, _, _, nextGridY := getGridData(copy)

		if currentGridY < nextGridY {
			maxX = currentMaxX
			maxY = currentMaxY
			gridX = currentGridX
			gridY = currentGridY
			break
		}

		for j := 0; j < len(points); j++ {
			current := points[j]
			newpoint := factorVelocity(current)
			points[j] = newpoint
		}
		second++
	}

	// Initialize grid
	grid := make([][]string, gridY, gridY)
	for i := 0; i < gridY; i++ {
		grid[i] = make([]string, gridX, gridX)
	}

	for _, v := range points {
		x := v.position.x + gridX - maxX - 1
		y := v.position.y + gridY - maxY - 1
		grid[y][x] = "#"
	}

	fmt.Println("Part 1")
	fmt.Println()
	for _, v := range grid {
		fmt.Printf("%1v\n", v)
	}

	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println(strconv.Itoa(second))

	fmt.Println()
	fmt.Println(time.Since(start))
}

func getGridData(points []point) (int, int, int, int) {
	minX, minY, maxX, maxY := getMinMaxXY(points)
	gridX := int(math.Abs(float64(minX)-float64(maxX))) + 1
	gridY := int(math.Abs(float64(minY)-float64(maxY))) + 1

	return maxX, maxY, gridX, gridY
}

func factorVelocity(p point) point {
	x := p.position.x + p.velocity.x
	y := p.position.y + p.velocity.y

	position := coord{x: x, y: y}

	return point{
		position: position,
		velocity: p.velocity}
}

func getMinMaxXY(points []point) (int, int, int, int) {
	minX := int(^uint(0) >> 1)
	minY := int(^uint(0) >> 1)
	maxX := -int(^uint(0)>>1) - 1
	maxY := -int(^uint(0)>>1) - 1

	for _, v := range points {
		if v.position.x < minX {
			minX = v.position.x
		}
		if v.position.y < minY {
			minY = v.position.y
		}
		if v.position.x > maxX {
			maxX = v.position.x
		}
		if v.position.y > maxY {
			maxY = v.position.y
		}
	}

	return minX, minY, maxX, maxY
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
