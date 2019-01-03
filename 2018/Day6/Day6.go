package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/brange/advent_of_code/shared/coord"
	"github.com/brange/advent_of_code/utils"
)

func parseInput(input string) ([]coord.Coord, int, int) {
	arr := strings.Split(input, "\n")
	var coords []coord.Coord
	re := regexp.MustCompile(`(\d+), (\d+)`)
	maxX := 0
	maxY := 0
	for _, row := range arr {
		co := re.FindAllStringSubmatch(row, 1)[0]
		x, _ := strconv.Atoi(co[1])
		y, _ := strconv.Atoi(co[2])
		coords = append(coords, coord.New(x, y))
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	return coords, maxX, maxY
}

func main() {
	input := utils.FetchInput(2018, 6)
	coords, maxX, maxY := parseInput(input)
	log.Println(fmt.Sprintf("maxX: %d, maxY: %d, numberCoords: %d", maxX, maxY, len(coords)))

	field := make([][]coord.Coord, maxY+1)

	var nullCoord coord.Coord
	for y := 0; y <= maxY; y++ {
		field[y] = make([]coord.Coord, maxX+1)
		for x := 0; x <= maxX; x++ {
			var c coord.Coord
			distance := maxX + maxY
			for _, coord := range coords {
				dist := coord.DistanceTo(x, y)
				if dist < distance {
					distance = dist
					c = coord
				} else if dist == distance {
					distance = dist
					c = nullCoord
				}
			}
			field[y][x] = c
		}
	}
	var infiniteCoords []coord.Coord
	add := func(c coord.Coord) {
		if !coord.Contains(infiniteCoords, c) {
			infiniteCoords = append(infiniteCoords, c)
		}
	}
	for y := 0; y <= maxY; y++ {
		add(field[y][0])
		add(field[y][maxX])
	}
	for x := 0; x <= maxX; x++ {
		add(field[0][x])
		add(field[maxY][x])
	}

	largestArea := 0
	for _, c := range coords {
		if !coord.Contains(infiniteCoords, c) {
			area := 0
			for y := 0; y <= maxY; y++ {
				for x := 0; x <= maxX; x++ {
					if c == field[y][x] {
						area++
					}
				}
			}

			if area > largestArea {
				largestArea = area
			}
		}
	}

	log.Println("Answer step 1:", largestArea)

	numberWithinRange := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			distance := 0
			for _, c := range coords {
				distance += c.DistanceTo(x, y)
				if distance > 10000 {
					break
				}
			}
			if distance <= 10000 {
				numberWithinRange++
			}
		}
	}

	log.Println("Answer step 2:", numberWithinRange)

}
