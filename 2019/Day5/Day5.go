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

func runProgram(arr []string, userInput int) int {
	ints := createIntArray(arr)

	for index := 0; index < len(ints); {
		input := ints[index]

		var optcode, mode1, mode2, mode3 int
		if input >= 1 && input <= 8 {
			optcode = input
		} else if input == 99 {
			log.Println("Breaking, got input 99")
			break
		} else {
			optcode = input % 100
			mode1 = ((input - optcode) % 1000) / 100
			mode2 = ((input - mode1 - optcode) % 10000) / 1000
			mode3 = ((input - mode2 - mode1 - optcode) % 100000) / 10000
		}

		if userInput >= 1 {
			log.Println("optcode:", optcode, "input", input, "mode1", mode1, "mode2", mode2, "mode3", mode3, "index:", index)
		}
		if optcode == 99 {
			log.Println("Breaking")
			break
		}

		var param1 int
		if mode1 == 1 {
			param1 = ints[index+1]
		} else {
			param1 = ints[ints[index+1]]
		}

		if optcode == 1 || optcode == 2 {
			val1 := param1
			var val2, pos int

			if mode2 == 1 {
				val2 = ints[index+2]
			} else {
				val2 = ints[ints[index+2]]
			}
			if mode3 == 1123123 /*ignored*/ {
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
			/*
			 * Opcode 3 takes a single integer as input and saves it to the address given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
			 * Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			 */
			if optcode == 3 {
				// Parameters that an instruction writes to will never be in immediate mode.
				ints[ints[index+1]] = userInput
			} else if optcode == 4 {
				if param1 != 0 {
					return param1
				}
				//fmt.Printf("%d ", param1)
			}
			index += 2
		} else if optcode == 5 || optcode == 6 {
			/*
			 * Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			 * Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			 */

			var param2 int
			_i := index
			if mode2 == 1 {
				param2 = ints[index+2]
			} else {
				param2 = ints[ints[index+2]]
			}
			if (optcode == 5 && param1 != 0) ||
				(optcode == 6 && param1 == 0) {
				index = param2
			} else {
				index += 3
			}
			log.Println("index, before:", _i, "after:", index)
		} else if optcode == 7 || optcode == 8 {
			/*
			 * Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			 * Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			 */
			var param2, param3 int
			if mode2 == 1 {
				param2 = ints[index+2]
			} else {
				param2 = ints[ints[index+2]]
			}

			if mode3 == 111 {
				// Parameters that an instruction writes to will never be in immediate mode.
				param3 = index + 3
			} else {
				param3 = ints[index+3]
			}

			if (optcode == 7 && param1 < param2) ||
				(optcode == 8 && param1 == param2) {
				ints[param3] = 1
			} else {
				ints[param3] = 0
			}

			index += 4
		}
		if optcode == 0 {
			break
		}

	}

	fmt.Println("")
	return -1
}

func main() {
	input := utils.FetchInput(2019, 5)
	arr := strings.Split(input, ",")

	log.Println("Part one")
	partOne := runProgram(arr, 1)
	log.Println("------------")
	log.Println("Part two")
	partTwo := runProgram(arr, 5)

	log.Println("Part one answer:", partOne)
	log.Println("Part two answer:", partTwo)

}
