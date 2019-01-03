package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/brange/advent_of_code/utils"
)

func reactToPolymers(input string) int {
	length := len(input)
	for true {
		for c := 'a'; c <= 'z'; c++ {
			lower := string(c)
			upper := strings.ToUpper(lower)
			input = strings.Replace(input, lower+upper, "", -1)
			input = strings.Replace(input, upper+lower, "", -1)
		}
		if len(input) == length {
			break
		}
		length = len(input)
	}
	return length
}

func main() {
	input := utils.FetchInput(2018, 5)

	step1 := reactToPolymers(input)
	log.Println("Answer step 1: ", step1)

	// step 2
	shortest := len(input)
	var removed string
	for c := 'a'; c <= 'z'; c++ {
		lower := string(c)
		upper := strings.ToUpper(lower)
		reduced := strings.Replace(input, lower, "", -1)
		reduced = strings.Replace(reduced, upper, "", -1)
		l := reactToPolymers(reduced)
		if l < shortest {
			shortest = l
			removed = lower
		}
	}

	log.Println(fmt.Sprintf("Answer step 2: %d (removed: %s)", shortest, removed))
}
