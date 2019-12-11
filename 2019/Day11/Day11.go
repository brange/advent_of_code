package main

import (
	"flag"
	"fmt"
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

	Black int = 0
	White int = 1

	Left  int = 90
	Right int = -90
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
				if int(pos) >= len(ints) {
					return 0
				}
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

		if debug {
			log.Printf("executing optcode %d from input %d", optcode, input)
		}

		if optcode == 99 {
			log.Println("Halting the program...")
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
var DAY = 11

var debug bool

type robot struct {
	location, direction coord
}

func (r *robot) move() {
	r.location.x += r.direction.x
	r.location.y += r.direction.y
}

type coord struct {
	x, y int
}

func (c *coord) rotate(degrees int) {
	rad := float64(degrees) * math.Pi / 180.0
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	px := float64(c.x)*cos - float64(c.y)*sin
	py := float64(c.x)*sin + float64(c.y)*cos
	c.x = int(px)
	c.y = int(py)
}

type panel struct {
	coord                 coord
	color, numberOfLayers int
}

func findPanel(panels *[]panel, coord coord) (panel, bool) {
	for _, p := range *panels {
		if p.coord.x == coord.x && p.coord.y == coord.y {
			return p, true
		}
	}

	return panel{coord: coord}, false
}

func main() {
	debugPtr := flag.Bool("debug", false, "Enable debug prints")
	flag.Parse()
	debug = *debugPtr
	if debug {
		log.Println("Enabling debug")
	}

	input := utils.FetchInput(YEAR, DAY)
	arr := strings.Split(input, ",")

	provideInput := func(initialColor int, panels *[]panel, out <-chan int64, in chan<- int64, robot robot) {
		log.Printf("Starting provideInput, initial color: %d, robot is %v, num panels %d", initialColor, robot, len(*panels))
		for {
			//panel, existed := findPanel(&panels, robot.location)
			panelIndex := math.MaxInt32
			panelColor := initialColor
			for index, p := range *panels {
				if p.coord.x == robot.location.x && p.coord.y == robot.location.y {
					panelIndex = index
					panelColor = p.color
					break
				}
			}

			if debug {
				log.Printf("Panel under robot is %v", (*panels)[panelIndex])
			}
			if debug {
				log.Printf("Sending color %d", panelColor)
			}
			in <- int64(panelColor)

			if panelIndex == math.MaxInt32 {
				panelIndex = len(*panels)
				*panels = append(*panels, panel{coord: robot.location, color: initialColor})
			}

			color := <-out
			dir := <-out
			((*panels)[panelIndex]).color = int(color)
			((*panels)[panelIndex]).numberOfLayers++

			if debug {
				log.Printf("Got color %d and direction %d from program, panel is now %v", color, dir, (*panels)[panelIndex])
			}
			// 0 means it should turn left 90 degrees, and 1 means it should turn right 90 degrees.
			if dir == 0 {
				robot.direction.rotate(Left)
			} else if dir == 1 {
				robot.direction.rotate(Right)
			} else {
				log.Fatalf("Invalid direction %d", dir)
			}
			robot.move()
			if debug {
				log.Printf("Robot is now %v", robot)
			}
		}
	}

	panelsPartOne := make([]panel, 0)

	robot := robot{location: coord{x: 0, y: 0}, direction: coord{x: 0, y: 1}}
	log.Printf("Robot is at start %v", robot)
	in := make(chan int64)
	out := make(chan int64)
	halt := make(chan bool)

	go runProgram(arr, in, out, halt)
	go provideInput(Black, &panelsPartOne, out, in, robot)
	<-halt
	log.Printf("Part one, number of panels painted: %d", len(panelsPartOne))
	if debug {
		for i, panel := range panelsPartOne {
			fmt.Printf("Panel %d: %v => color %d, layers: %d\n", i, panel, panel.color, panel.numberOfLayers)
		}
	}

	panelsPartTwo := make([]panel, 0)
	robot.location = coord{x: 0, y: 0}
	robot.direction = coord{x: 0, y: 1}

	log.Printf("Robot is at start %v", robot)
	in = make(chan int64)
	out = make(chan int64)
	halt = make(chan bool)

	go runProgram(arr, in, out, halt)
	go provideInput(White, &panelsPartTwo, out, in, robot)
	<-halt

	minX := math.MaxInt32
	minY := math.MaxInt32
	maxX := math.MinInt32
	maxY := math.MinInt32
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	for _, panel := range panelsPartTwo {
		minX = min(minX, panel.coord.x)
		minY = min(minY, panel.coord.y)

		maxX = max(maxX, panel.coord.x)
		maxY = max(maxY, panel.coord.y)
	}
	log.Printf("Panels is between (%d,%d) and (%d,%d)", minX, minY, maxX, maxY)
	widht := maxX - minX
	height := maxY - minY
	log.Printf("Size: %d x %d", widht, height)

	log.Printf("number of panels, part2 = %d", len(panelsPartTwo))
	color := func(x, y int) int {
		for _, panel := range panelsPartTwo {
			if panel.coord.x == x && panel.coord.y == y {
				return panel.color
			}
		}
		return -1
	}
	for h := height; h >= 0; h-- {
		fmt.Println("")
		for w := 0; w < widht; w++ {
			c := color(w+minX, h+minY)
			if c == White {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println("")
}
