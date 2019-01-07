package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/brange/advent_of_code/shared/math"
	"github.com/brange/advent_of_code/utils"
)

type position struct {
	x, y         int
	veloX, veloY int
}

func parseInput(input string) []position {
	re := regexp.MustCompile(`position=<\s*([\-\d]+),\s*([\-\d]+)> velocity=<\s*([\-\d]+),\s*([\-\d]+)>`)

	var positions []position
	for _, line := range strings.Split(input, "\n") {
		matches := re.FindAllStringSubmatch(line, 1)
		groups := matches[0]
		x, _ := strconv.Atoi(groups[1])
		y, _ := strconv.Atoi(groups[2])
		veloX, _ := strconv.Atoi(groups[3])
		veloY, _ := strconv.Atoi(groups[4])

		pos := position{x: x, y: y, veloX: veloX, veloY: veloY}
		positions = append(positions, pos)
	}

	return positions
}

func sumDistance(positions []position) int {
	minX, maxX, minY, maxY := getMinsAndMax(positions)
	return math.AbsInt(maxX-minX) + math.AbsInt(maxY-minY)
}

func movePositions(positions *[]position, multiplier int) {
	for index := range *positions {
		(*positions)[index].x += (*positions)[index].veloX * multiplier
		(*positions)[index].y += (*positions)[index].veloY * multiplier
	}
}

func getMinsAndMax(positions []position) (int, int, int, int) {
	maxX := positions[0].x
	minX := positions[0].x

	maxY := positions[0].y
	minY := positions[0].y
	for _, pos := range positions {
		minX = math.MinInt(pos.x, minX)
		maxX = math.MaxInt(pos.x, maxX)

		minY = math.MinInt(pos.y, minY)
		maxY = math.MaxInt(pos.y, maxY)
	}

	return minX, maxX, minY, maxY
}

func printPositions(positions []position) {

	hasPos := func(x, y int) bool {
		for _, pos := range positions {
			if pos.x == x && pos.y == y {
				return true
			}
		}
		return false
	}
	s := ""
	minX, maxX, minY, maxY := getMinsAndMax(positions)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if hasPos(x, y) {
				s += "#"
			} else {
				s += " "
			}
		}
		s += "\n"
	}
	fmt.Print(s)
}

func main() {
	input := utils.FetchInput(2018, 10)
	positions := parseInput(input)

	dist := sumDistance(positions)

	minutes := 0
	for true {
		movePositions(&positions, 1)
		d := sumDistance(positions)
		if d > dist {
			break
		}
		minutes++
		dist = d
	}
	movePositions(&positions, -1)
	printPositions(positions)
	log.Printf("Minutes (answer step 2): %d\n", minutes)
}
