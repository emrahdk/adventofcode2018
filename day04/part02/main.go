package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type logline struct {
	time time.Time
	val  string
}

func main() {
	// lines := []string{"[1518-11-01 00:00] Guard #10 begins shift",
	// 	"[1518-11-01 00:25] wakes up",
	// 	"[1518-11-01 00:05] falls asleep",
	// 	"[1518-11-01 00:30] falls asleep",
	// 	"[1518-11-01 00:55] wakes up",
	// 	"[1518-11-01 23:58] Guard #99 begins shift",
	// 	"[1518-11-02 00:40] falls asleep",
	// 	"[1518-11-02 00:50] wakes up",
	// 	"[1518-11-03 00:05] Guard #10 begins shift",
	// 	"[1518-11-03 00:24] falls asleep",
	// 	"[1518-11-03 00:29] wakes up",
	// 	"[1518-11-04 00:02] Guard #99 begins shift",
	// 	"[1518-11-04 00:36] falls asleep",
	// 	"[1518-11-04 00:46] wakes up",
	// 	"[1518-11-05 00:03] Guard #99 begins shift",
	// 	"[1518-11-05 00:45] falls asleep",
	// 	"[1518-11-05 00:55] wakes up"}

	lines := readInput()

	kvp := make([]logline, 0)

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

		x := logline{
			time: date,
			val:  fmt.Sprintf("%s %s", val, val2)}

		kvp = append(kvp, x)
	}

	// Sort by time
	sort.Slice(kvp, func(i, j int) bool {
		return kvp[i].time.Before(kvp[j].time)
	})

	// Group by guards
	guardid2 := ""
	guards := map[string][]logline{}

	for _, v := range kvp {
		if strings.Contains(v.val, "Guard") {
			// fmt.Sscanf("%s #%d", v.val, &guardid2)
			guardid2 = strings.Split(v.val, "#")[1]
			continue
		}

		guards[guardid2] = append(guards[guardid2], v)
	}

	guardsSleep := map[string][]int{}

	for k, v := range guards {
		guardsSleep[k] = getGuardSleep(v)
	}

	mostSleepMinuteDups := map[string]map[int]int{}

	for k, v := range guardsSleep {
		hval, count := dup_count(v)

		mostSleepMinuteDups[k] = map[int]int{hval: count}
	}

	for k, v := range mostSleepMinuteDups {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func dup_count(list []int) (int, int) {

	duplicate_frequency := make(map[int]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}

	hvalue := -1
	highest := 0
	for k, v := range duplicate_frequency {
		if highest < v {
			highest = v
			hvalue = k
		}
	}

	return hvalue, highest
}

func getGuardSleep(loglines []logline) []int {
	minutes := make([]int, 0)
	sleepTime := time.Time{}

	currentTime := time.Time{}

	for _, v := range loglines {

		if strings.Contains(v.val, "falls") {
			sleepTime = v.time
			currentTime = v.time
		} else {
			currentTruncated := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
			thisTruncated := time.Date(v.time.Year(), v.time.Month(), v.time.Day(), 0, 0, 0, 0, v.time.Location())
			if !currentTruncated.Equal(thisTruncated) {
				continue
			}
			// sd := sleepTime.Day()
			// wd := v.time.Day()

			// if sd != wd {
			// 	continue
			// }
			s := sleepTime.Minute()
			w := v.time.Minute()

			for i := s; i < w; i++ {
				minutes = append(minutes, i)
			}

		}

	}

	return minutes
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
