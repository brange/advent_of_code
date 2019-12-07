package main

import (
	"flag"
	"log"
	"strconv"
	_ "strconv"
	"strings"
	_ "strings"

	_ "github.com/brange/advent_of_code/shared"
	_ "github.com/brange/advent_of_code/shared/math"
	"github.com/brange/advent_of_code/utils"
)

func createIntArray(arr []string) []int {
	ints := make([]int, len(arr))
	for index, a := range arr {
		ints[index], _ = strconv.Atoi(a)
	}
	return ints
}

const (
	ParameterMode int = 0
	ImmediateMode int = 1
)

func getParameter(arr []int, index, mode int) int {
	if mode == ParameterMode {
		return arr[arr[index]]
	} else if mode == ImmediateMode {
		return arr[index]
	}
	log.Fatalln("Invalid mode", mode)
	return -1
}
func getPosition(arr []int, index int) int {
	// 'Parameters that an instruction writes to will never be in immediate mode.'
	return arr[index]
}

func runProgram(arr []string, phaseSetting, userInput int) int {
	ints := createIntArray(arr)

	for index := 0; index < len(ints); {
		input := ints[index]

		var optcode int
		var mode1, mode2, mode3 int
		if input >= 1 && input <= 8 {
			optcode = input
		} else if input == 99 {
			if debug {
				log.Println("Breaking, got input 99")
			}
			break
		} else {
			optcode = input % 100
			mode1 = ((input - optcode) % 1000) / 100
			mode2 = ((input - mode1 - optcode) % 10000) / 1000
			mode3 = ((input - mode2 - mode1 - optcode) % 100000) / 10000
		}

		if userInput >= 1 {
			if debug {
				log.Println("optcode:", optcode, "input", input, "mode1", mode1, "mode2", mode2, "mode3", mode3, "index:", index)
			}
		}
		if optcode == 99 {
			if debug {
				log.Println("Breaking")
			}
			break
		}

		param1 := getParameter(ints, index+1, mode1)

		if optcode == 1 || optcode == 2 {
			param2 := getParameter(ints, index+2, mode2)
			pos := getPosition(ints, index+3)

			if optcode == 1 {
				ints[pos] = param1 + param2
			} else if optcode == 2 {
				ints[pos] = param1 * param2
			}
			index += 4
		} else if optcode == 3 || optcode == 4 {
			/*
			 * Opcode 3 takes a single integer as input and saves it to the address given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
			 * Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			 */
			if optcode == 3 {
				// Parameters that an instruction writes to will never be in immediate mode.
				ints[getPosition(ints, index+1)] = phaseSetting
				phaseSetting = userInput
			} else if optcode == 4 {
				if param1 != 0 {
					return param1
				}
			}
			index += 2
		} else if optcode == 5 || optcode == 6 {
			/*
			 * Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			 * Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			 */
			param2 := getParameter(ints, index+2, mode2)

			if (optcode == 5 && param1 != 0) ||
				(optcode == 6 && param1 == 0) {
				index = param2
			} else {
				index += 3
			}
		} else if optcode == 7 || optcode == 8 {
			/*
			 * Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			 * Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			 */
			param2 := getParameter(ints, index+2, mode2)
			pos := getPosition(ints, index+3)

			if (optcode == 7 && param1 < param2) ||
				(optcode == 8 && param1 == param2) {
				ints[pos] = 1
			} else {
				ints[pos] = 0
			}

			index += 4
		}
		if optcode == 0 {
			break
		}

	}

	return -1
}

var YEAR = 2019
var DAY = 7

var debug bool

func main() {
	debugPtr := flag.Bool("debug", false, "Enable debug prints")
	flag.Parse()
	debug = *debugPtr
	if debug {
		log.Println("Enabling debug")
	}

	input := utils.FetchInput(YEAR, DAY)
	arr := strings.Split(input, ",")

	log.Println("Got input", input)

	maxThrust := 0
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			if b == a {
				continue
			}
			for c := 0; c < 5; c++ {
				if c == a || c == b {
					continue
				}
				for d := 0; d < 5; d++ {
					if d == a || d == b || d == c {
						continue
					}
					for e := 0; e < 5; e++ {
						if e == a || e == b || e == c || e == d {
							continue
						}
						ampA := runProgram(arr, a, 0)
						ampB := runProgram(arr, b, ampA)
						ampC := runProgram(arr, c, ampB)
						ampD := runProgram(arr, d, ampC)
						ampE := runProgram(arr, e, ampD)

						//log.Println("ampA:", ampA)
						if ampE > maxThrust {
							maxThrust = ampE
						}
					}
				}
			}
		}
	}

	log.Println("Part one:", maxThrust)

}
