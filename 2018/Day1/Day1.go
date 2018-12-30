package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/brange/advent_of_code/utils"
)

func main() {

	var input = utils.FetchInput(2018, 1)

	freq := 0
	arr := strings.Split(input, "\n")

	frequences := make(map[int]bool)

	for _, change := range arr {
		i, _ := strconv.Atoi(change)
		freq += i
		frequences[freq] = true
	}

	log.Println("Answer step 1: ", freq)

	log.Println("List size: ", len(arr))
	cnt := 0
l:
	for true {
		for _, change := range arr {
			cnt++
			i, _ := strconv.Atoi(change)
			if i == 0 {
				continue
			}
			freq += i
			_, hasFreq := frequences[freq]
			if hasFreq {
				log.Println("Found the first frequency that repeats itself: ", freq)
				break l
			}
			frequences[freq] = true
		}
	}
	log.Println("Counter: ", cnt)

}
