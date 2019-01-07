/**
Find the fuel cell's rack ID, which is its X coordinate plus 10.
Begin with a power level of the rack ID times the Y coordinate.
Increase the power level by the value of the grid serial number (your puzzle input).
Set the power level to itself multiplied by the rack ID.
Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
Subtract 5 from the power level.
*/
package main

import (
	"log"
	"strconv"

	"github.com/brange/advent_of_code/utils"
)

func calculatePower(x, y, gridSerialNumber int) int {
	rackID := x + 10
	power := rackID * y
	power += gridSerialNumber
	power *= rackID

	hundred := ((power % 1000) - (power % 100)) / 100
	return hundred - 5
}

func makeGrid(gridSerialNumber int) [][]int {
	grid := make([][]int, 301)
	grid[0] = make([]int, 301)
	for y := 1; y <= 300; y++ {
		grid[y] = make([]int, 301)
		for x := 1; x <= 300; x++ {
			power := calculatePower(x, y, gridSerialNumber)
			grid[y][x] = power + grid[y-1][x] + grid[y][x-1] - grid[y-1][x-1]
		}
	}
	return grid
}

func calculateSquarePower(grid [][]int, squareSize int) (int, int, int) {
	var bestX, bestY, best int
	for y := squareSize; y <= 300; y++ {
		for x := squareSize; x <= 300; x++ {
			r := grid[y][x] - grid[y-squareSize][x] - grid[y][x-squareSize] + grid[y-squareSize][x-squareSize]
			if r > best {
				best = r
				bestX = x - (squareSize - 1)
				bestY = y - (squareSize - 1)
			}
		}
	}

	return bestX, bestY, best

}

func main() {

	gridSerialNumber, _ := strconv.Atoi(utils.FetchInput(2018, 11))

	log.Println("Using grid serial number:", gridSerialNumber)

	grid := makeGrid(gridSerialNumber)
	log.Println("Grid created")
	x, y, _ := calculateSquarePower(grid, 3)
	log.Printf("Answer step 1 (x,y): %d,%d\n", x, y)

	// Step 2
	var bestX, bestY, best, bestSize int
	for size := 1; size <= 300; size++ {
		x, y, res := calculateSquarePower(grid, size)
		if res > best {
			bestX = x
			bestY = y
			best = res
			bestSize = size
		}
	}

	log.Printf("Answer step 2 (x,y,size): %d,%d,%d\n", bestX, bestY, bestSize)

}
