package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	input := []rune(readInput()[0])

	for i := 0; i < len(input); i++ {
		if i+1 == len(input) {
			break
		}

		if unicode.IsLower(input[i]) {
			if input[i+1] == unicode.ToUpper(input[i]) {
				input = append(input[:i], input[i+2:]...)
				i = 0
				continue
			}
		} else {
			if input[i+1] == unicode.ToLower(input[i]) {
				input = append(input[:i], input[i+2:]...)
				i = 0
				continue
			}
		}
	}

	input = input[2:] // investigate this
	fmt.Println(len(input))
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
