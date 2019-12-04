package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/brange/advent_of_code/utils"
)

var YEAR = 2019
var DAY = 4

func matches(pwd int) (bool, bool) {
	asString := strconv.Itoa(pwd)

	twoAdjecent := false
	increasing := true
	exactTwoAdjecent := false
	for i := 0; i < 4; i++ {
		var d int
		if i > 0 {
			d, _ = strconv.Atoi(asString[i-1 : i])
		} else {
			d = -1
		}
		a, _ := strconv.Atoi(asString[i : i+1])
		b, _ := strconv.Atoi(asString[i+1 : i+2])
		c, _ := strconv.Atoi(asString[i+2 : i+3])
		if a > b || b > c {
			increasing = false
			break
		}
		if !twoAdjecent && (a == b || b == c) {
			twoAdjecent = true
		}
		if !exactTwoAdjecent {
			exactTwoAdjecent = (a == b && b != c && d != a) ||
				(i == 3 && b == c && a != b)
		}
	}

	partOne := twoAdjecent && increasing
	partTwo := exactTwoAdjecent && increasing
	return partOne, partTwo
}

func main() {
	input := utils.FetchInput(YEAR, DAY)

	min, _ := strconv.Atoi(strings.Split(input, "-")[0])
	max, _ := strconv.Atoi(strings.Split(input, "-")[1])
	log.Println("min:", min, "max:", max)

	count := 0
	countPartTwo := 0
	for i := min; i <= max; i++ {
		partOneMatch, partTwoMatch := matches(i)
		if partOneMatch {
			count++
		}
		if partTwoMatch {
			countPartTwo++
		}
	}
	log.Println("Part one:", count)
	log.Println("Part two", countPartTwo)
}
