package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
)

func main() {
	coords := readInput()

	allIds := make([]string, 0)
	for _, v := range coords {
		allIds = append(allIds, strings.ToLower(v.name))
	}

	const HMAX int = 360
	const VMAX int = 360

	rect := [HMAX][VMAX]grid{}

	// Build rect
	for _, v := range coords {
		for i := 0; i < VMAX; i++ {
			for j := 0; j < HMAX; j++ {
				if v.y == i && v.x == j {
					// START POS
					rect[i][j] = grid{v.name, 0}
				} else {
					// CHECK IF CURRENT HAS LOWER OR EQUAL (.) DISTANCE
					newGrid := grid{strings.ToLower(v.name), manhattanDist(v, coord{x: j, y: i})}
					currentGrid := rect[i][j]

					if currentGrid.name == "" {
						rect[i][j] = newGrid
					} else {
						if newGrid.dist < currentGrid.dist {
							rect[i][j] = newGrid
						} else if newGrid.dist == currentGrid.dist {
							rect[i][j] = grid{".", currentGrid.dist}
						}
					}
				}
			}
		}
	}

	// Print map
	// for _, v := range rect {
	// 	for _, w := range v {
	// 		fmt.Printf("%2v", w.name)
	// 	}
	// 	fmt.Println()
	// }

	infinites := getInfiniteAreas(rect)
	finites := make([]string, 0)

	for _, v := range allIds {
		if !contains(infinites, v) {
			finites = append(finites, v)
		}
	}

	areaSizes := map[string]int{}

	for _, v := range finites {
		for i := 0; i < VMAX; i++ {
			for j := 0; j < HMAX; j++ {
				if strings.ToLower(rect[i][j].name) == v {
					areaSizes[v]++
				}
			}
		}
	}

	largestArea := 0
	for _, v := range areaSizes {
		if v > largestArea {
			largestArea = v
		}
	}

	fmt.Println(largestArea)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func generateTwoLetterString() string {
	bytes := make([]byte, 2)

	for i := 0; i < 2; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}

	return strings.ToUpper(string(bytes))
}

func getInfiniteAreas(rect [360][360]grid) []string {
	//infites
	infinites := make([]string, 0)

	// FIRST ROW
	for _, v := range rect[0] {
		if v.name != "." {
			infinites = append(infinites, strings.ToLower(v.name))
		}
	}

	// LAST ROW
	for _, v := range rect[9] {
		if v.name != "." {
			infinites = append(infinites, strings.ToLower(v.name))
		}
	}

	// FIRST AND LAST COL
	for i := 0; i < 360; i++ {
		if rect[i][0].name != "." {
			infinites = append(infinites, strings.ToLower(rect[i][0].name))
		}

		if rect[i][9].name != "." {
			infinites = append(infinites, strings.ToLower(rect[i][9].name))
		}
	}

	return removeDuplicates(infinites)
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func manhattanDist(a coord, b coord) int {
	return int(math.Abs(float64(b.x-a.x)) + math.Abs(float64(b.y-a.y)))
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

func getCoord(id string, input string) coord {
	c := coord{name: id}
	fmt.Sscanf(input, "%d, %d", &c.x, &c.y)
	return c
}

func readInput() []coord {
	names := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "BB", "CC", "DD", "EE", "FF", "GG", "HH", "II", "JJ", "KK", "LL", "MM", "NN", "OO", "PP", "QQ", "RR", "SS", "TT", "UU", "VV", "WW", "XX"}

	lines := make([]string, 0)
	coordMap := map[string]coord{}

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

	for i, v := range lines {
		// TODO figure out why random names breaks code
		// name := ""
		// notUnique := true

		// for notUnique {
		// 	if _, ok := coordMap[name]; !ok {
		// 		name = generateTwoLetterString()
		// 		notUnique = false
		// 	}
		// }

		coordMap[names[i]] = getCoord(names[i], v)
	}

	coords := make([]coord, 0)

	for _, v := range coordMap {
		coords = append(coords, v)
	}

	return coords
}
