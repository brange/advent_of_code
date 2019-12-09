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

var YEAR = 2019
var DAY = 8

var debug bool

var width = 25
var height = 6
var size = width * height

func createLayers(arr []string) [][]int {
	numberOfLayers := len(arr) / size
	log.Println("array len", len(arr), "layer size:", size, "number of layers:", numberOfLayers)
	layers := make([][]int, numberOfLayers)

	var layer []int
	var layerI int
	for i, val := range arr {
		if i%size == 0 {
			layer = make([]int, size)
			layers[layerI] = layer
			layerI++
		}
		layer[i%size] = toInt(val)

	}

	return layers
}

func toInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func countDigits(layer []int, digit int) int {
	count := 0
	for _, d := range layer {
		if d == digit {
			count++
		}
	}
	return count
}

const (
	black       int = 0
	white       int = 1
	transparent int = 2
)

func main() {
	debugPtr := flag.Bool("debug", false, "Enable debug prints")
	flag.Parse()
	debug = *debugPtr
	if debug {
		log.Println("Enabling debug")
	}

	input := utils.FetchInput(YEAR, DAY)

	log.Println("Got input", input)

	layers := createLayers(strings.Split(input, ""))
	//what is the number of 1 digits multiplied by the number of 2 digits?
	numZeros := size
	answer := 0
	for _, layer := range layers {
		zeros := countDigits(layer, 0)
		if zeros < numZeros {
			answer = countDigits(layer, 1) * countDigits(layer, 2)
			numZeros = zeros
		}
	}

	log.Println("Part one", answer)

	numberOfLayers := len(layers)
	for h := 0; h < height; h++ {
		fmt.Println("")
		for w := 0; w < width; w++ {
			pixel := w + h*width
			//fmt.Print(pixel, ":")

		c:
			for l := 0; l < numberOfLayers; l++ {
				color := layers[l][pixel]
				if color == black || color == white {
					if color == black {
						fmt.Print(" ")
					} else {
						fmt.Print(color)
					}
					break c
				}
			}

			//fmt.Print(", ")
		}

	}
	fmt.Println("")

}
