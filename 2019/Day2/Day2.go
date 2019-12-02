package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/brange/advent_of_code/utils"
)

func createIntArray(arr []string) []int {
	ints := make([]int, len(arr))
	for index, a := range arr {
		ints[index], _ = strconv.Atoi(a)
	}
	return ints
}

func runProgram(arr []string, noun, verb int) int {
	ints := createIntArray(arr)
	// Part 1
	// replace position 1 with the value 12 and replace position 2 with the value 2
	ints[1] = noun
	ints[2] = verb
	//log.Println("Input: ", ints)

	for index, optcode := range ints {
		if index%4 != 0 {
			continue
		}
		if optcode == 99 {
			//log.Println("Breaking")
			break
		}
		val1 := ints[ints[index+1]]
		val2 := ints[ints[index+2]]
		pos := ints[index+3]

		if optcode == 1 {
			ints[pos] = val1 + val2
		} else if optcode == 2 {
			ints[pos] = val1 * val2
		}
	}

	return ints[0]
}

func main() {
	arr := strings.Split(utils.FetchInput(2019, 2), ",")

	partOne := runProgram(arr, 12, 2)

	log.Println("Part one: ", partOne)

l:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			result := runProgram(arr, noun, verb)
			if result == 19690720 {
				log.Println("Got result (Part two):", (100*noun + verb), "with noun", noun, "and verb", verb)
				break l
			}
		}
	}
}
