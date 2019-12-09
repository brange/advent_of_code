package main

import (
	"flag"
	"log"
	"math"
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

func runProgram(arr []string, inChannel <-chan int64, output chan<- int64, halt chan<- bool) {
	ints := createIntArray(arr)

	var relativeBase int64

	for index := 0; index < len(ints); {

		input := ints[index]

		getMode := func(paramIndex int) int {
			pow := func(a, b int) int {
				return int(math.Pow(float64(a), float64(b)))
			}
			return int(input) % pow(10, paramIndex+2) / pow(10, paramIndex+1)
		}

		getParameter := func(paramIndex int) int64 {
			_index := index + paramIndex
			mode := getMode(paramIndex)

			if int(mode) == ParameterMode {
				pos := ints[_index]
				return ints[pos]
			} else if int(mode) == ImmediateMode {
				return ints[_index]
			} else if int(mode) == RelativeMode {
				relativeIndex := relativeBase + ints[_index]
				return ints[relativeIndex]
			}
			log.Fatalln("Invalid mode", mode)
			return int64(-1)
		}

		getPosition := func(paramIndex int) int64 {
			// 'Parameters that an instruction writes to will never be in immediate mode.'
			mode := getMode(paramIndex)
			_index := index + paramIndex
			var position int64
			if int(mode) == ParameterMode {
				position = ints[_index]
			} else if int(mode) == RelativeMode {
				position = relativeBase + ints[_index]
			} else {
				log.Fatal("Invalid mode for getPosition:", mode)
				return -1
			}
			if int(position) >= len(ints) {
				// Find a better way to increase the slices
				//log.Println("Will expand from", len(ints), "to", pos+1)
				for k := len(ints); k <= (int(position) + 1); k++ {
					ints = append(ints, 0)
				}
			}
			return position
		}

		optcode := int(input % 100)

		if optcode == 99 {
			halt <- true
			break
		}

		param1 := getParameter(1)

		switch optcode {
		case 1:
			// Add
			pos := getPosition(3)
			ints[pos] = param1 + getParameter(2)
			index += 4
		case 2:
			// Multiplication
			pos := getPosition(3)
			ints[pos] = param1 * getParameter(2)
			index += 4
		case 3:
			// Opcode 3 takes a single integer as input and saves it to the address given by its only parameter.
			// For example, the instruction 3,50 would take an input value and store it at address 50.
			pos := getPosition(1)
			(ints)[pos] = <-inChannel
			index += 2
		case 4:
			// Opcode 4 outputs the value of its only parameter.
			// For example, the instruction 4,50 would output the value at address 50.
			output <- param1
			index += 2
		case 5:
			// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter.
			// Otherwise, it does nothing.
			param2 := getParameter(2)
			if param1 != 0 {
				index = int(param2)
			} else {
				index += 3
			}
		case 6:
			// Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			param2 := getParameter(2)
			if param1 == 0 {
				index = int(param2)
			} else {
				index += 3
			}
		case 7:
			// Opcode 7 is less than: if the first parameter is less than the second parameter,
			// it stores 1 in the position given by the third parameter. Otherwise, it stores 0.

			param2 := getParameter(2)
			pos := getPosition(3)

			if param1 < param2 {
				ints[pos] = 1
			} else {
				ints[pos] = 0
			}
			index += 4
		case 8:
			// Opcode 8 is equals: if the first parameter is equal to the second parameter,
			// it stores 1 in the position given by the third parameter. Otherwise, it stores 0.

			param2 := getParameter(2)
			pos := getPosition(3)

			if param1 == param2 {
				ints[pos] = 1
			} else {
				ints[pos] = 0
			}
			index += 4
		case 9:
			// Opcode 9 adjusts the relative base by the value of its only parameter.
			// The relative base increases (or decreases, if the value is negative) by the value of the parameter.
			relativeBase += param1
			index += 2
		default:
			log.Fatal("Invalid optcode:", optcode)
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
	<-halt

	go runProgram(arr, in, out, halt)
	in <- 2
	partTwo := <-out
	log.Println("Part two:", partTwo)
	<-halt

}
