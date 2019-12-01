package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/brange/advent_of_code/utils"
)

func calculateFuel(mass int) int {
	fuel := mass / 3
	fuel -= 2
	if fuel < 0 {
		fuel = 0
	}
	return fuel
}

func main() {

	var input = utils.FetchInput(2019, 1)

	arr := strings.Split(input, "\n")

	totalFuel := 0

	for _, line := range arr {
		mass, _ := strconv.Atoi(line)

		fuel := calculateFuel(mass)
		fuelForFuel := fuel

		for ok := true; ok; ok = (fuelForFuel > 0) {
			fuelForFuel = calculateFuel(fuelForFuel)
			totalFuel += fuelForFuel
		}

		totalFuel += fuel
	}

	log.Println("totalFueal: ", totalFuel)
}
