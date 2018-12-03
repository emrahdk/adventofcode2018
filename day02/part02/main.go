package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	strings, _ := readInputFile()
	curr := ""
	next := ""

	curr, next = process(strings)

	if curr != "" && next != "" {
		result := stripOutDifference(curr, next)
		fmt.Println(result)
	}
}

func process(strings []string) (curr string, next string) {
	for i := 0; i < len(strings); i++ {
		if i+1 == len(strings) {
			break
		}

		curr := strings[i]

		for j := i + 1; j < len(strings); j++ {
			next := strings[j]
			res := compareStrings(curr, next)

			if res {
				return curr, next
			}

		}
	}

	return "", ""
}

func stripOutDifference(current string, test string) (res string) {
	position := 0

	for i := 0; i < len(current); i++ {
		if string(current[i]) != string(test[i]) {
			position = i
			break
		}
	}

	return current[:position] + current[position+1:]
}

func compareStrings(current string, test string) (res bool) {
	maxDiff := 1

	for i := 0; i < len(current); i++ {
		if string(current[i]) != string(test[i]) {
			maxDiff--
			if maxDiff < 0 {
				return false
			}
		}
	}

	return true
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
