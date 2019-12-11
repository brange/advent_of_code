package main

import (
	"flag"
	"log"
	"math"
	"sort"
	"strings"

	"github.com/brange/advent_of_code/utils"
)

var YEAR = 2019
var DAY = 10

var debug bool

type asteroid struct {
	x, y int
}

func (ast *asteroid) angle(other asteroid) float64 {

	atan2 := math.Atan2(float64(ast.y-other.y), float64(other.x-ast.x))
	return atan2 * 180 / math.Pi
}

func (ast *asteroid) distance(other asteroid) int {
	return abs(ast.x-other.x) + abs(ast.y-other.y)
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func main() {
	debugPtr := flag.Bool("debug", false, "Enable debug prints")
	flag.Parse()
	debug = *debugPtr
	if debug {
		log.Println("Enabling debug")
	}

	input := utils.FetchInput(YEAR, DAY)

	log.Printf("Got input:\n%s", input)

	asteroids := make([]asteroid, 0)
	for y, xs := range strings.Split(input, "\n") {
		for x, ast := range strings.Split(xs, "") {
			if ast == "#" {
				a := asteroid{x: x, y: y}
				asteroids = append(asteroids, a)
			}
		}
	}
	log.Printf("Found %d asteroids", len(asteroids))

	var bestMatch asteroid
	numberOfLOS := 0
	var bestMatchAngles []float64

	for _, asteroid := range asteroids {
		angles := make([]float64, 0)

		for _, other := range asteroids {
			if other == asteroid {
				continue
			}
			angle := asteroid.angle(other)
			canSee := true

			for _, a := range angles {
				if a == angle {
					canSee = false
					break
				}
			}

			if debug {
				log.Printf("Angle between %v and %v is %f => %t", asteroid, other, angle, canSee)
			}

			if canSee {
				angles = append(angles, angle)
			}

		}

		if debug {
			log.Printf("%v got %d angles", asteroid, len(angles))
		}

		if len(angles) > numberOfLOS {
			bestMatch = asteroid
			numberOfLOS = len(angles)
			bestMatchAngles = angles
		}
	}

	log.Printf("Part one: %v (%d asteorids in sight)", bestMatch, numberOfLOS)

	destroyed := make([]float64, 0)

	sort.Float64s(bestMatchAngles)

	start := 90.0
	log.Printf("Starting at %f from %v", start, bestMatch)
	index := len(bestMatchAngles)
	started := false

	findAsteroid := func(angle float64) asteroid {
		distnace := 99999
		var asteroid asteroid
		for _, other := range asteroids {
			if bestMatch.angle(other) == angle &&
				bestMatch.distance(other) < distnace {
				distnace = bestMatch.distance(other)
				asteroid = other
			}
		}
		return asteroid
	}
	for {
		index--
		if !started && bestMatchAngles[index] <= start {
			started = true
		}

		if started {
			if debug {
				log.Printf("Destroying %f => %v", bestMatchAngles[index], findAsteroid(bestMatchAngles[index]))
			}
			destroyed = append(destroyed, bestMatchAngles[index])
		}

		if len(destroyed) == 200 {
			log.Printf("The 200 destroyed has angle %f", bestMatchAngles[index])
			for _, other := range asteroids {
				if bestMatch.angle(other) == bestMatchAngles[index] {

					// Find closest with that angle
					twoHundredDestroyed := other
					distance := bestMatch.distance(twoHundredDestroyed)
					for _, closest := range asteroids {
						if bestMatch.angle(closest) == bestMatchAngles[index] &&
							bestMatch.distance(closest) < distance {
							distance = bestMatch.distance(closest)
							twoHundredDestroyed = closest
						}
					}
					log.Printf("The 200 destroyed asteroid is %v", twoHundredDestroyed)
					log.Printf("Answer part two: %d", (twoHundredDestroyed.x*100 + twoHundredDestroyed.y))
				}
			}
			break
		}
		if index == 0 {
			index = len(bestMatchAngles)
		}
	}
}
