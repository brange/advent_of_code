package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	_ "strconv"
	"strings"
	_ "strings"

	_ "github.com/brange/advent_of_code/shared"
	_ "github.com/brange/advent_of_code/shared/math"
	"github.com/brange/advent_of_code/utils"
)

func createIntArray(arr []string) []int64 {
	ints := make([]int64, len(arr))
	for index, a := range arr {
		ints[index], _ = strconv.ParseInt(a, 10, 64)
	}
	return ints
}

const (
	ParameterMode int = 0
	ImmediateMode int = 1
	RelativeMode  int = 2
)

func getPosition(arr []int64, index int64) int64 {
	// 'Parameters that an instruction writes to will never be in immediate mode.'
	return arr[index]
}

func runProgram(arr []string, inChannel <-chan int64, output chan<- int64, halt chan<- bool) {
	ints := createIntArray(arr)

	//log.Println("phaseSetting:", phaseSetting, "userInput:", userInput, "startIndex:", startIndex, "len(ints)", len(*ints))
	for index := 0; index < len(ints); {

		input := (ints)[index]

		var optcode int64
		var mode1, mode2 int64
		if input >= 1 && input <= 8 {
			optcode = input
		} else if input == 99 {
			if debug {
				log.Println("Breaking, got input 99")
			}
			halt <- true
			break
		} else {
			optcode = input % 100
			mode1 = ((input - optcode) % 1000) / 100
			mode2 = ((input - mode1 - optcode) % 10000) / 1000
			//mode3 = ((input - mode2 - mode1 - optcode) % 100000) / 10000
		}

		if optcode == 99 {
			halt <- true
			break
		}

		expandIfNeeded := func(pos int64) {
			if int(pos) > len(ints) {
				// Find a better way to increase the slices
				for k := len(ints); k <= int(pos); k++ {
					ints = append(ints, 0)
				}
			}
		}
		var relativeBase int64

		getParameter := func(arr []int64, index, mode int64) int64 {
			if int(mode) == ParameterMode {
				return arr[arr[index]]
			} else if int(mode) == ImmediateMode {
				return arr[index]
			} else if int(mode) == RelativeMode {
				return arr[relativeBase+arr[index]]
			}
			log.Fatalln("Invalid mode", mode)
			return int64(-1)
		}
		param1 := getParameter(ints, int64(index+1), mode1)

		if optcode == 1 || optcode == 2 {
			param2 := getParameter(ints, int64(index+2), mode2)
			pos := getPosition(ints, int64(index+3))

			expandIfNeeded(pos)
			if optcode == 1 {
				(ints)[pos] = param1 + param2
			} else if optcode == 2 {
				(ints)[pos] = param1 * param2
			}
			index += 4
		} else if optcode == 3 || optcode == 4 {
			/*
			 * Opcode 3 takes a single integer as input and saves it to the address given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
			 * Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			 */
			if optcode == 3 {
				// Parameters that an instruction writes to will never be in immediate mode.
				log.Print("Waiting on input ")
				read := <-inChannel
				fmt.Println(", got: ", read)
				(ints)[getPosition(ints, int64(index+1))] = read
			} else if optcode == 4 {
				//if param1 != 0 {
				output <- param1
				//}
			}
			index += 2
		} else if optcode == 5 || optcode == 6 {
			/*
			 * Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			 * Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			 */
			param2 := getParameter(ints, int64(index+2), mode2)

			if (optcode == 5 && param1 != 0) ||
				(optcode == 6 && param1 == 0) {
				index = int(param2)
			} else {
				index += 3
			}
		} else if optcode == 7 || optcode == 8 {
			/*
			 * Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			 * Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			 */
			param2 := getParameter(ints, int64(index+2), mode2)
			pos := getPosition(ints, int64(index+3))

			expandIfNeeded(pos)

			if (optcode == 7 && param1 < param2) ||
				(optcode == 8 && param1 == param2) {
				(ints)[pos] = 1
			} else {
				(ints)[pos] = 0
			}

			index += 4
		} else if optcode == 9 {
			// Opcode 9 adjusts the relative base by the value of its only parameter.
			// The relative base increases (or decreases, if the value is negative) by the value of the parameter.
			relativeBase = relativeBase + param1
			index += 2
		} else {
			log.Fatalln("Invalid opcode: ", optcode)
		}
		if optcode == 0 {
			break
		}

	}
}

var YEAR = 2019
var DAY = 9

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

	in := make(chan int64)
	out := make(chan int64)

	halt := make(chan bool)

	go runProgram(arr, in, out, halt)
	in <- 1
	partOne := <-out
	log.Println("Part one:", partOne)
}
