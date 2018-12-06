package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	input := []rune(readInput()[0])
	alphabet := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lengths := make([]int, 0)

	for _, v := range alphabet {
		tempInput := append([]rune(nil), input...)
		for i := 0; i < len(tempInput); i++ {

			if v == tempInput[i] || v == unicode.ToUpper(tempInput[i]) {
				tempInput = append(tempInput[:i], tempInput[i+1:]...)

				next := i - 2
				if next < -1 {
					next = -1
				}
				i = next
			}
		}

		lengths = append(lengths, dothingfrompart01(tempInput))
	}

	sort.Ints(lengths)

	fmt.Println(lengths[0])
	fmt.Println(time.Since(start))
}

func dothingfrompart01(input []rune) int {
	for i := 0; i < len(input); i++ {
		if i+1 == len(input) {
			break
		}

		if unicode.IsLower(input[i]) {
			if input[i+1] == unicode.ToUpper(input[i]) {
				input = append(input[:i], input[i+2:]...)

				next := i - 2
				if next < -1 {
					next = -1
				}
				i = next

				continue
			}
		} else {
			if input[i+1] == unicode.ToLower(input[i]) {
				input = append(input[:i], input[i+2:]...)

				next := i - 2
				if next < -1 {
					next = -1
				}
				i = next

				continue
			}
		}
	}

	return len(input)
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
