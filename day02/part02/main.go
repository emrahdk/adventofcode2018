package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	strings, _ := readInputFile()
	a := ""
	b := ""

	a, b = process(strings)

	fmt.Println(stripOutDifference(a, b))
}

func process(strings []string) (a string, b string) {
	for i, v := range strings {
		if i+1 == len(strings) {
			break
		}

		a := v

		for j := i + 1; j < len(strings); j++ {
			b := strings[j]

			if compareStrings(a, b) {
				return a, b
			}

		}
	}

	return "", ""
}

func stripOutDifference(a string, b string) (res string) {
	position := 0

	for i, v := range a {
		if string(v) != string(b[i]) {
			position = i
			break
		}
	}

	return a[:position] + a[position+1:]
}

func compareStrings(a string, b string) (res bool) {
	diff := 1

	for i, v := range a {
		if string(v) != string(b[i]) {
			diff--
			if diff < 0 {
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
