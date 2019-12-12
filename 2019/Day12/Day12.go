package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	_ "strconv"
	"strings"
	_ "strings"

	_ "github.com/brange/advent_of_code/shared"
	_ "github.com/brange/advent_of_code/shared/math"
	"github.com/brange/advent_of_code/utils"
)

var YEAR = 2019
var DAY = 12

var debug bool

type moon struct {
	position coord
	velocity coord
}

func (m *moon) getEnergy() int {
	return m.position.getEnergy() * m.velocity.getEnergy()
}

type coord struct {
	x, y, z int
}

func (c *coord) add(other coord) {
	c.x += other.x
	c.y += other.y
	c.z += other.z
}

func (c *coord) getEnergy() int {
	return abs(c.x) + abs(c.y) + abs(c.z)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -1 * a
}

func getVelocity(a, b int) int {
	if a == b {
		return 0
	}
	if a > b {
		return -1
	}
	return 1
}

func main() {
	debugPtr := flag.Bool("debug", false, "Enable debug prints")
	flag.Parse()
	debug = *debugPtr
	if debug {
		log.Println("Enabling debug")
	}

	input := utils.FetchInput(YEAR, DAY)

	moonRegexp := regexp.MustCompile(`^<x=(\-?\d+), y=(\-?\d+), z=(\-?\d+)>$`)

	log.Printf("Got input\n%s", input)
	arr := strings.Split(input, "\n")
	moons := make([]moon, len(arr))
	originalMoons := make([]moon, len(moons))

	for index, line := range arr {
		match := moonRegexp.FindAllStringSubmatch(line, -1)[0]
		toInt := func(str string) int {
			i, _ := strconv.Atoi(str)
			return i
		}
		moons[index] = moon{position: coord{x: toInt(match[1]), y: toInt(match[2]), z: toInt(match[3])}}
		originalMoons[index] = moons[index]
	}
	log.Printf("Got moons: %v", moons)

	step := 0

	if debug {
		fmt.Printf("After step %d\n", step)
		for index, moon := range moons {
			fmt.Printf("Moon %d: %v\n", index, moon)
		}
		fmt.Println("")
	}

	reps := make([]int, 3)
	for k := 0; k < 3; k++ {
		reps[k] = -1
	}

	var totalEnergy int
	for i := 0; true; i++ {
		step++
		for index, moon := range moons {

			for _, other := range moons {
				moons[index].velocity.x += getVelocity(moon.position.x, other.position.x)
				moons[index].velocity.y += getVelocity(moon.position.y, other.position.y)
				moons[index].velocity.z += getVelocity(moon.position.z, other.position.z)
			}

			if debug {
				log.Printf("moon is now %v", moon)
			}
		}

		for index, moon := range moons {
			if debug {
				log.Printf("Adding %v to the moon %v", moon.velocity, moon)
			}
			moons[index].position.add(moon.velocity)
		}

		allX := true
		allY := true
		allZ := true
		for moonIndex, moon := range moons {
			if allX &&
				moon.position.x != originalMoons[moonIndex].position.x ||
				moon.velocity.x != originalMoons[moonIndex].velocity.x {
				allX = false
			}
			if allY &&
				moon.position.y != originalMoons[moonIndex].position.y ||
				moon.velocity.y != originalMoons[moonIndex].velocity.y {
				allY = false
			}
			if allZ &&
				moon.position.z != originalMoons[moonIndex].position.z ||
				moon.velocity.z != originalMoons[moonIndex].velocity.z {
				allZ = false
			}
		}
		if allX && reps[0] == -1 {
			reps[0] = step
		}
		if allY && reps[1] == -1 {
			reps[1] = step
		}
		if allZ && reps[2] == -1 {
			reps[2] = step
		}

		foundAllRepititions := reps[0] != -1 && reps[1] != -1 && reps[2] != -1

		if debug {
			fmt.Printf("After step %d\n", step)
			for index, moon := range moons {
				fmt.Printf("Moon %d: %v\n", index, moon)
			}
			fmt.Println("")
		}
		if step == 1000 {
			for index, moon := range moons {
				fmt.Printf("Energy for moon %d (%v) is %d\n", index, moon, moon.getEnergy())
				totalEnergy += moon.getEnergy()
			}
		}
		if foundAllRepititions && step > 1000 {
			break
		}
	}

	log.Printf("Total energy after 1000 steps (part one): %d", totalEnergy)

	log.Printf("reps: %v", reps)
	log.Printf("Part two: %d", lcm(reps...))
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (lcm) via GCD
func lcm(integers ...int) int {
	result := integers[0] * integers[1] / gcd(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
