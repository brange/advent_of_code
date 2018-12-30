package main

import (
	"log"
	"strconv"
	"strings"

	aoc_utils "github.com/brange/advent_of_code/utils"
)

func main() {
	inputStr := aoc_utils.FetchInput(2018, 2)
	input := strings.Split(inputStr, "\n")

	numberTwos := 0
	numberThrees := 0

	var a, b string

	for _, row := range input {
		if row == "" {
			continue
		}
		m := make(map[string]int)
		chars := strings.Split(row, "")

		for _, char := range chars {
			i, ok := m[char]
			if ok {
				m[char] = i + 1
			} else {
				m[char] = 1
			}
		}

		hasTwos := false
		hasThrees := false
		for _, value := range m {
			if value == 2 {
				hasTwos = true
			} else if value == 3 {
				hasThrees = true
			}
		}
		if hasTwos {
			numberTwos++
		}
		if hasThrees {
			numberThrees++
		}

		if a == "" {
			for _, row2 := range input {
				chars2 := strings.Split(row2, "")
				nbr := 0
				for index, char := range chars2 {
					if chars[index] != char {
						nbr++
					}
					if nbr > 1 {
						break
					}
				}
				if nbr == 1 {
					a = row
					b = row2
					break
				}
			}
		}
	}

	sum := numberTwos * numberThrees
	log.Println("Result step 1: " + strconv.Itoa(sum))
	log.Println("String a: " + a)
	log.Println("String b: " + b)
	if a != "" && b != "" {
		step2Result := ""
		chars := strings.Split(a, "")
		charsB := strings.Split(b, "")
		for index, char := range chars {
			if charsB[index] == char {
				step2Result += char
			}
		}
		log.Println("Result step2: '" + step2Result + "'")
	}
}
