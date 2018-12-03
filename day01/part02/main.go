package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	frequency := 0
	repeatedFrequency := 0
	changes, err := readInputFile()
	results := make(map[int]int)
	results[0] = 0
	runloop := true

	if err != nil {
		fmt.Println(err)
		return
	}

	for runloop {
		for i := 0; i < len(changes); i++ {
			frequency += changes[i]
			if _, ok := results[frequency]; ok {
				repeatedFrequency = frequency
				runloop = false
				break
			}

			results[frequency] = frequency
		}
	}

	fmt.Println(repeatedFrequency)
}

func readInputFile() (frequencies []int, err error) {
	bytes, err := ioutil.ReadFile("../input/input.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\r\n")

	// Assign cap to avoid resize on every append.
	frequencies = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		n, err := strconv.Atoi(l)

		if err != nil {
			return nil, err
		}
		frequencies = append(frequencies, n)
	}

	return frequencies, nil
}
