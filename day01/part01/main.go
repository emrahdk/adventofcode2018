package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	frequency := 0
	changes, err := readInputFile()

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(changes); i++ {
		frequency += changes[i]
	}

	fmt.Println(frequency)
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
