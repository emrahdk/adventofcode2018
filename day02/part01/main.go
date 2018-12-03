package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	threeLetterCounts := 0
	twoLetterCounts := 0

	strings, _ := readInputFile()

	for _, v := range strings {
		twos, threes := processString(v)

		twoLetterCounts += twos
		threeLetterCounts += threes
	}

	fmt.Printf("%v * %v = %v", twoLetterCounts, threeLetterCounts, twoLetterCounts*threeLetterCounts)
}

func processString(letters string) (twos int, threes int) {
	chars := make(map[string]int)
	twos = 0
	threes = 0

	// Group by letter occurences
	for i := 0; i < len(letters); i++ {
		currentChar := string(letters[i])

		if _, ok := chars[currentChar]; ok {
			chars[currentChar] = chars[currentChar] + 1
		} else {
			chars[currentChar] = 1
		}
	}

	// Filter everything below 2
	for k, v := range chars {
		if v < 2 {
			delete(chars, k)
		}
	}

	if findFirst(chars, 2) {
		twos = 1
	}

	if findFirst(chars, 3) {
		threes = 1
	}

	return twos, threes
}

func findFirst(chars map[string]int, val int) (res bool) {
	res = false
	for _, v := range chars {
		if v == val {
			res = true
			break
		}
	}
	return res
}

func readInputFile() (letters []string, err error) {
	bytes, err := ioutil.ReadFile("../input/input.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")

	// Assign cap to avoid resize on every append.
	letters = make([]string, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		letters = append(letters, l)
	}

	return letters, nil
}
