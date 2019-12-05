package main

import (
	"fmt"
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

func runProgram(arr []string) int {
	ints := createIntArray(arr)

	for index := 0; index < len(ints); {
		input := ints[index]
		if input == 99 {
			//log.Println("Breaking")
			break
		}
		var optcode, mode1, mode2, mode3 int
		if input >= 1 && input <= 4 {
			optcode = input
		} else {
			optcode = input % 100
			mode1 = ((input - optcode) % 1000) / 100
			mode2 = ((input - mode1 - optcode) % 10000) / 1000
			mode3 = ((input - mode2 - mode1 - optcode) % 100000) / 10000
		}

		var param1 int
		if mode1 == 1 {
			param1 = ints[index+1]
		} else {
			param1 = ints[ints[index+1]]
		}

		log.Println("optcode:", optcode, "input", input, "mode1", mode1, "mode2", mode2, "mode3", mode3)
		if optcode == 1 || optcode == 2 {
			val1 := param1
			var val2, pos int

			if mode2 == 1 {
				val2 = ints[index+2]
			} else {
				val2 = ints[ints[index+2]]
			}
			if mode3 == 1 {
				// Parameters that an instruction writes to will never be in immediate mode.
				pos = index + 3
			} else {
				pos = ints[index+3]
			}
			if optcode == 1 {
				ints[pos] = val1 + val2
			} else if optcode == 2 {
				ints[pos] = val1 * val2
			}
			index += 4
		} else if optcode == 3 || optcode == 4 {
			if optcode == 3 {
				// Parameters that an instruction writes to will never be in immediate mode.
				ints[ints[index+1]] = 1
			} else if optcode == 4 {
				fmt.Printf("%d", param1)
			}
			index += 2
		} /*else if optcode == 5 || optcode == 6 {
			var param2 int
			if mode2 == 1 {
				param2 = ints[index+2]
			} else {
				param2 = ints[ints[index+2]]
			}
			if (optcode == 5 && param1 != 0) ||
				(optcode == 6 && param1 == 0) {
				index = param2
			} else {
				index += 2
			}
		} else if optcode == 7 || optcode == 8 {
			var param2, param3 int
			if mode2 == 1 {
				param2 = ints[index+2]
			} else {
				param2 = ints[ints[index+2]]
			}
			if mode3 == 1111 {
				// Parameters that an instruction writes to will never be in immediate mode.
				param3 = ints[index+3]
			} else {
				param3 = ints[ints[index+3]]
			}

			if (optcode == 7 && param1 < param2) ||
				(optcode == 8 && param1 == param2) {
				ints[param3] = 1
			} else {
				ints[param3] = 0
			}

		}
		*/
		/*
			Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
		*/
		if optcode == 0 {
			break
		}

	}

	fmt.Println("")
	return ints[0]
}

func main() {
	arr := strings.Split(utils.FetchInput(2019, 5), ",")

	partOne := runProgram(arr)

	log.Println("Part one: ", partOne)
}
