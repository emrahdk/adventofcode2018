package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	input := []rune(readInput()[0])
	alphabet := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphabetMap := map[string]int{}

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

		alphabetMap[string(v)] = dothingfrompart01(tempInput)
	}

	lowestValue := int(^uint(0) >> 1)
	for _, v := range alphabetMap {
		if v < lowestValue {
			lowestValue = v
		}
	}

	fmt.Println(lowestValue)
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
