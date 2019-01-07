package main

import (
	"log"
	"regexp"
	"strconv"

	"github.com/brange/advent_of_code/utils"
)

type marble struct {
	value uint64
	prev  *marble
	next  *marble
}

func parseInput(input string) (int, int) {
	re := regexp.MustCompile(`(\d+) players; last marble is worth (\d+) points`)
	groups := re.FindAllStringSubmatch(input, 1)[0]
	players, _ := strconv.Atoi(groups[1])
	points, _ := strconv.Atoi(groups[2])

	return players, points
}

func playGame(numberPlayers, maxPoint int) uint64 {
	log.Printf("Players: %d, points: %d\n", numberPlayers, maxPoint)

	players := make([]uint64, numberPlayers)

	var currentMarble *marble
	currentMarble = &marble{value: 0}
	currentMarble.next = currentMarble
	currentMarble.prev = currentMarble

	nextMarbleValue := uint64(1)

start:
	for true {
		for index := range players {

			if nextMarbleValue%23 == 0 {

				players[index] += nextMarbleValue
				toBeRemoved := currentMarble.prev
				for i := 6; i > 0; i-- {
					toBeRemoved = toBeRemoved.prev
				}
				prev := toBeRemoved.prev
				next := toBeRemoved.next
				players[index] += toBeRemoved.value
				prev.next = next
				next.prev = prev
				currentMarble = next
			} else {
				first := currentMarble.next
				second := first.next
				marble := &marble{value: nextMarbleValue}
				marble.next = second
				marble.prev = first
				first.next = marble
				second.prev = marble
				currentMarble = marble
			}

			nextMarbleValue++

			if nextMarbleValue >= uint64(maxPoint) {
				break start
			}
		}
	}
	highScore := uint64(0)
	for _, value := range players {
		if value > highScore {
			highScore = value
		}
	}

	return highScore
}

func main() {
	input := utils.FetchInput(2018, 9)

	numberPlayers, maxPoint := parseInput(input)
	answerStepOne := playGame(numberPlayers, maxPoint)

	log.Printf("High score step 1: %d\n", answerStepOne)

	answerStepTwo := playGame(numberPlayers, maxPoint*100)
	log.Printf("High score step 2: %d\n", answerStepTwo)
}
