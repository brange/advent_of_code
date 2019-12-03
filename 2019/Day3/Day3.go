package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/brange/advent_of_code/shared"
	"github.com/brange/advent_of_code/shared/math"
	"github.com/brange/advent_of_code/utils"
)

func travel(vect string) (int, int) {
	value, _ := strconv.Atoi(vect[1:])
	dir := vect[0:1]
	if dir == "U" {
		return 0, value
	} else if dir == "D" {
		return 0, -1 * value
	} else if dir == "R" {
		return value, 0
	} else if dir == "L" {
		return -1 * value, 0
	}

	return 0, 0
}
func main() {
	input := utils.FetchInput(2019, 3)

	wires := strings.Split(input, "\n")

	paths := make([]map[shared.Point]int, len(wires))
	for wireIndex, wire := range wires {
		m := make(map[shared.Point]int)
		paths[wireIndex] = m

		currX := 0
		currY := 0
		steps := 0

		for _, tr := range strings.Split(wire, ",") {
			x, y := travel(tr)

			//log.Println("Going to traverse ", x, ",", y)
			negX := 1
			negY := 1
			if x < 0 {
				negX *= -1
			}
			if y < 0 {
				negY *= -1
			}

			x *= negX
			y *= negY
			for _x := 1; _x <= x; _x++ {
				steps++
				key := shared.Point{X: currX + _x*negX, Y: currY}
				if m[key] == 0 {
					m[key] = steps
				}
				//log.Println("Marked", key, ":", m[key])
			}
			for _y := 1; _y <= y; _y++ {
				steps++
				key := shared.Point{X: currX, Y: currY + _y*negY}
				if m[key] == 0 {
					m[key] = steps
				}
				//log.Println("Marked", key, ":", m[key])
			}
			currX += x * negX
			currY += y * negY
			//log.Println("CurrX:", currX, "currY:", currY)
		}
	}

	distancePartOne := -1
	distancePartTwo := -1

	for key, _ := range paths[0] {
		_, hasKey := paths[1][key]
		if hasKey {
			x := key.X
			y := key.Y

			_dist := math.AbsInt(x) + math.AbsInt(y)

			//log.Println("_dist", _dist, "coord: ", x, ", ", y)
			if distancePartOne == -1 || (_dist != 0 && _dist < distancePartOne) {
				distancePartOne = _dist
			}
			_dist2 := paths[0][key] + paths[1][key]
			if distancePartTwo == -1 || (_dist2 != 0 && _dist2 < distancePartTwo) {
				distancePartTwo = _dist2
			}
		}
	}

	log.Println("distance part one", distancePartOne)
	log.Println("distance part two", distancePartTwo)
}
