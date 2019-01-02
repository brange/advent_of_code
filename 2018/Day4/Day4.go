package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/brange/advent_of_code/utils"
)

type inputRow struct {
	date  time.Time
	event string
}

type guard struct {
	id              string
	sleepingMinutes map[int]int
}

func (g guard) totalSleeping() int {
	sum := 0
	for _, v := range g.sleepingMinutes {
		sum += v
	}
	return sum
}

func parseInput(input string) []inputRow {
	arr := strings.Split(input, "\n")
	re := regexp.MustCompile(`\[(\d+)-(\d+)-(\d+) (\d+):(\d+)\] (.*)`)

	i := func(i string) int {
		res, _ := strconv.Atoi(i)
		return res
	}

	var rows []inputRow

	for _, r := range arr {
		g := re.FindAllStringSubmatch(r, -1)[0]
		date := time.Date(i(g[1]), time.Month(i(g[2])), i(g[3]), i(g[4]), i(g[5]), 0, 0, time.UTC)
		rows = append(rows, inputRow{date, g[6]})
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i].date.Before(rows[j].date)
	})

	return rows
}

func checkSleeping(rows []inputRow) []guard {

	findGuard := func(guards *[]guard, id string) int {
		for index, guard := range *guards {
			if guard.id == id {
				return index
			}
		}

		return -1
	}
	var guards []guard

	var currentGuard guard
	var currentSleepingMinute int
	re := regexp.MustCompile(`Guard #(\d+) begins shift`)
	for _, row := range rows {
		matches := re.FindAllStringSubmatch(row.event, 1)
		if len(matches) > 0 {
			id := matches[0][1]
			index := findGuard(&guards, id)
			if index == -1 {
				currentGuard = guard{id, make(map[int]int)}
				guards = append(guards, currentGuard)
			} else {
				currentGuard = guards[index]
			}
		} else {
			if row.event == "falls asleep" {
				currentSleepingMinute = row.date.Minute()
			} else {
				// Wakes up
				endSleep := row.date.Minute()
				for m := currentSleepingMinute; m < endSleep; m++ {
					currentGuard.sleepingMinutes[m]++
				}
			}
		}
	}

	return guards
}

func main() {
	input := utils.FetchInput(2018, 4)

	rows := parseInput(input)
	guards := checkSleeping(rows)
	sort.Slice(guards, func(i, j int) bool {
		return guards[i].totalSleeping() > guards[j].totalSleeping()
	})
	sleepyGuard := guards[0]
	var min int
	a := 0
	for minute, amount := range sleepyGuard.sleepingMinutes {
		if amount > a {
			min = minute
			a = amount
		}
	}
	id, _ := strconv.Atoi(sleepyGuard.id)
	answer1 := strconv.Itoa(id * min)
	fmt.Println("Answer step 1: " + answer1)

	// step 2
	a = 0
	var g guard
	min = 0
	for _, guard := range guards {
		for minute, amount := range guard.sleepingMinutes {
			if amount > a {
				g = guard
				a = amount
				min = minute
			}
		}
	}
	id, _ = strconv.Atoi(g.id)
	answer2 := strconv.Itoa(id * min)
	fmt.Println("Answer step 2: " + answer2)

}
