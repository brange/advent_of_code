package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/brange/advent_of_code/utils"
)

type rect struct {
	x      int
	y      int
	width  int
	height int
	id     int
}

func _int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func parseInput(input string) []rect {
	var rects []rect
	arr := strings.Split(input, "\n")

	var re = regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

	for _, row := range arr {
		groups := re.FindAllStringSubmatch(row, -1)[0]
		rects = append(rects, rect{id: _int(groups[1]), x: _int(groups[2]), y: _int(groups[3]), width: _int(groups[4]), height: _int(groups[5])})

	}

	return rects
}

func rectToString(rect rect) string {
	return fmt.Sprintf("#%d @ %d,%d: %dx%d", rect.id, rect.x, rect.y, rect.width, rect.height)
}

func checkRectangles(rects []rect) (string, string) {
	xMax := 0
	yMax := 0
	for _, rect := range rects {
		x := rect.x + rect.width
		y := rect.y + rect.height
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}
	var matrix = make([][]int, yMax+1)
	doubles := 0
	for _, rect := range rects {

		for y := rect.y; y < (rect.y + rect.height); y++ {
			if matrix[y] == nil {
				matrix[y] = make([]int, xMax+1)
			}
			for x := rect.x; x < (rect.x + rect.width); x++ {
				matrix[y][x] = matrix[y][x] + 1
				if matrix[y][x] == 2 {
					doubles++
				}
			}
		}
	}

	// Step 2
	step2Answer := 0
	for _, rect := range rects {
		overlap := false
	Coord:
		for y := rect.y; y < (rect.y + rect.height); y++ {
			for x := rect.x; x < (rect.x + rect.width); x++ {
				if matrix[y][x] > 1 {
					overlap = true
					break Coord
				}
			}
		}

		if !overlap {
			step2Answer = rect.id
			break
		}

	}

	return strconv.Itoa(doubles), fmt.Sprintf("ID %d", step2Answer)

}

func main() {
	input := utils.FetchInput(2018, 3)
	rects := parseInput(input)

	answerStepOne, answerStepTwo := checkRectangles(rects)
	fmt.Println("Answer step1: " + answerStepOne)
	fmt.Println("Answer step2: " + answerStepTwo)

}
