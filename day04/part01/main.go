package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	lines := []string{"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:05] falls asleep",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:55] wakes up",
		"[1518-11-01 23:58] Guard #99 begins shift",
		"[1518-11-02 00:40] falls asleep",
		"[1518-11-02 00:50] wakes up",
		"[1518-11-03 00:05] Guard #10 begins shift",
		"[1518-11-03 00:24] falls asleep",
		"[1518-11-03 00:29] wakes up",
		"[1518-11-04 00:02] Guard #99 begins shift",
		"[1518-11-04 00:36] falls asleep",
		"[1518-11-04 00:46] wakes up",
		"[1518-11-05 00:03] Guard #99 begins shift",
		"[1518-11-05 00:45] falls asleep",
		"[1518-11-05 00:55] wakes up"}


	kvp := map[time.Time]string{}

	for _, v := range lines {
		year := 0
		month := 0
		day := 0
		hour := 0
		minute := 0
		val := ""
		val2 := ""
		fmt.Sscanf(v, "[%d-%d-%d %d:%d] %s %s", &year, &month, &day, &hour, &minute, &val, &val2)

		date := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

		kvp[date] = 
	}
	fmt.Println(kvp)
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
