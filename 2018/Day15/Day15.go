package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/brange/advent_of_code/shared"
	aoc_utils "github.com/brange/advent_of_code/utils"
)

type warriorType int
type cavern int

const (
	goblin warriorType = 0
	elf    warriorType = 1

	wall      cavern = 0
	openSpace cavern = 1
)

type warrior struct {
	point       shared.Point
	race        warriorType
	health      int
	attackPower int
}

func generateMap(input string) ([][]cavern, []warrior) {
	lines := strings.Split(input, "\n")
	world := make([][]cavern, len(lines))
	var warriors []warrior

	for y, line := range lines {
		world[y] = make([]cavern, len(line))
		chars := strings.Split(line, "")
		for x, char := range chars {
			if char == "#" {
				world[y][x] = wall
			} else {
				world[y][x] = openSpace
				if char == "E" || char == "G" {
					var race warriorType
					if char == "E" {
						race = elf
					} else {
						race = goblin
					}
					p := shared.Point{X: x, Y: y}
					warrior := warrior{p, race, 200, 3}
					warriors = append(warriors, warrior)
				}
			}
		}
	}

	return world, warriors
}

func printMap(world [][]cavern, warriors []warrior) {
	for y := 0; y < len(world); y++ {
		for x := 0; x < len(world[y]); x++ {
			index, warrior := findWarrior(warriors, shared.Point{X: x, Y: y})
			var symbol string
			if index >= 0 {
				if warrior.race == elf {
					symbol = "E"
				} else {
					symbol = "G"
				}
			} else {
				if world[y][x] == wall {
					symbol = "#"
				} else {
					symbol = "."
				}
			}
			fmt.Print(symbol)
		}
		if y == 0 {
			nbrElves, elfHP := countWarriors(warriors, elf)
			nbrGoblins, goblinHP := countWarriors(warriors, goblin)
			fmt.Printf("\tElves: %d (%d HP), Goblins: %d (%d HP)", nbrElves, elfHP, nbrGoblins, goblinHP)
		}
		fmt.Println("")
	}
}

func findWarrior(warriors []warrior, point shared.Point) (int, warrior) {
	var _warrior warrior
	for index, w := range warriors {
		if shared.Equals(w.point, point) && w.health > 0 {
			return index, w
		}
	}

	return -1, _warrior
}

func sortWarriors(warriors []warrior) {
	sort.Slice(warriors, func(i, j int) bool {
		var a = warriors[i]
		var b = warriors[j]
		return shared.FirstInReadingOrder(a.point, b.point)
	})
}

func countWarriors(warriors []warrior, race warriorType) (int, int) {
	sum := 0
	health := 0
	for _, w := range warriors {
		if w.race == race && w.health > 0 {
			sum++
			health += w.health
		}
	}

	return sum, health
}

func findPath(world [][]cavern, warriors []warrior, step int, attackerPoint shared.Point, pathPoints *[]shared.PathPoint, nextPoints []shared.Point, maxSteps int) ([]shared.Point, int) {
	if step > maxSteps {
		return nil, -1
	}

	//log.Println("FindPath: step :" + strconv.Itoa(step))

	var points []shared.Point
	for _, p := range nextPoints {
		//	log.Println(fmt.Sprintf("In loop 1, p: %v", p))
		for _, nearbyPoint := range getNearbyPoints(p) {
			if nearbyPoint.X < 0 ||
				nearbyPoint.Y < 0 ||
				nearbyPoint.Y > len(world) ||
				nearbyPoint.X > len(world[nearbyPoint.Y]) {
				continue
			}
			//		log.Println(fmt.Sprintf("In loop 2, nearbyPoint: %v", nearbyPoint))
			points = append(points, nearbyPoint)
		}
	}

	//log.Println(fmt.Sprintf("Going to check points: %v, got nextPoints: %v", points, nextPoints))

	var nextPoints2 []shared.Point

	for _, point := range points {

		_, foundPP := shared.FindPathPoint(*pathPoints, point)
		if foundPP {
			// Don't go back again..
			continue
		}

		//log.Println(fmt.Sprintf("Checking point %d, %d. Step: %d", point.X, point.Y, step))

		if shared.Equals(point, attackerPoint) {
			// Back at attackerPoint, was this the best way so far?
			//log.Println(fmt.Sprintf("Found attacker after %d steps from %v", step, pathPoints))
			var firstPoints []shared.Point
			for _, pp := range *pathPoints {
				if pp.Distance == (step-1) && shared.Distance(pp.Point, attackerPoint) == 1 {
					firstPoints = append(firstPoints, pp.Point)
				}
			}
			return firstPoints, step
		}

		var index, _ = findWarrior(warriors, point)
		if index >= 0 || world[point.Y][point.X] == wall {
			// Something in the way, don't go there!
		} else {
			// Open space, lets explore

			var pp, foundPP = shared.FindPathPoint(*pathPoints, point)
			if !foundPP || pp.Distance >= step {
				// Not been here before, or been here before but with a larger distance to the target (is that even possible??)
				*pathPoints = append(*pathPoints, shared.PathPoint{Point: point, Distance: step})
				nextPoints2 = append(nextPoints2, point)
			}
		}
	}
	if len(nextPoints2) > 0 {
		//log.Println(fmt.Sprintf("Didn't find the target, searching for it at %d", nextPoints2))
		return findPath(world, warriors, step+1, attackerPoint, pathPoints, nextPoints2, maxSteps)
	}
	return nil, -1
}

func findClosestEnemy(world [][]cavern, warriors []warrior, attacker warrior) (warrior, []shared.Point, bool) {
	var targetWalkingPoints []shared.Point
	distanceToTarget := 10000
	var target warrior
	foundATarget := false
	for _, warrior := range warriors {
		if warrior.race == attacker.race || warrior.health <= 0 {
			//if shared.Equals(attacker.point, warrior.point) {
			continue
		}
		var nextPoints []shared.Point
		nextPoints = append(nextPoints, warrior.point)
		var pathPoints []shared.PathPoint
		pathPoints = append(pathPoints, shared.PathPoint{Point: warrior.point, Distance: 0})
		walkingPoints, stepsToTarget := findPath(world, warriors, 1, attacker.point, &pathPoints, nextPoints, 200)
		if stepsToTarget > 0 && (stepsToTarget < distanceToTarget ||
			(stepsToTarget == distanceToTarget && shared.FirstInReadingOrder(warrior.point, target.point))) {
			//log.Println(fmt.Sprintf("Found a path from attacker at %v to target at %v, via %v", attacker.point, warrior.point, pathPoints))
			distanceToTarget = stepsToTarget
			targetWalkingPoints = walkingPoints
			target = warrior
			foundATarget = true
		}
	}

	return target, targetWalkingPoints, foundATarget
}

func findFirstEnemyWithinReach(warriors []warrior, attacker warrior) (int, warrior) {
	var target warrior
	targetIndex := -1
	for _, t := range getNearbyPoints(attacker.point) {
		index, _target := findWarrior(warriors, t)
		if index >= 0 && _target.race != attacker.race && _target.health > 0 {
			if targetIndex == -1 ||
				_target.health < target.health ||
				(_target.health == target.health && shared.FirstInReadingOrder(_target.point, target.point)) {
				// No previously target or, _target is "before" target in reading order
				target = _target
				targetIndex = index
			}
		}
	}

	return targetIndex, target
}

func getNearbyPoints(point shared.Point) [4]shared.Point {
	points := [...]shared.Point{
		shared.Point{point.X, point.Y - 1},
		shared.Point{point.X, point.Y + 1},
		shared.Point{point.X - 1, point.Y},
		shared.Point{point.X + 1, point.Y},
	}
	return points
}

func printWarriors(warriors []warrior) {
	for _, w := range warriors {
		log.Println(fmt.Sprintf("Warrior: %v", w))
	}
}

func battleRound(world [][]cavern, warriors []warrior) {
	sortWarriors(warriors)
	for attackerIndex, attacker := range warriors {
		if attacker.health <= 0 {
			continue
		}

		targetIndex, _ := findFirstEnemyWithinReach(warriors, attacker)
		if targetIndex >= 0 {
			warriors[targetIndex].health -= attacker.attackPower
			//log.Println(fmt.Sprintf("Found target before walking, %v is attacking %v", attacker.point, warriors[targetIndex].point))
		} else {
			// Lets walk
			_, walkingPoints, foundATarget := findClosestEnemy(world, warriors, attacker)
			if foundATarget {
				sort.Slice(walkingPoints, func(i, j int) bool {
					return shared.FirstInReadingOrder(walkingPoints[i], walkingPoints[j])
				})
				//log.Println(fmt.Sprintf("Moving attacker from %v to %v, all pathPoints %v", attacker.point, walkingPoints[0], walkingPoints))
				warriors[attackerIndex].point = walkingPoints[0]
				attacker.point = walkingPoints[0]

				// Something to hit on now?
				targetIndex, _ := findFirstEnemyWithinReach(warriors, attacker)
				if targetIndex >= 0 {
					warriors[targetIndex].health -= attacker.attackPower
					//log.Println(fmt.Sprintf("Found target before walking, %v is attacking %v", attacker.point, warriors[targetIndex].point))
				}
			}
		}
	}
}

func step1(input string) {
	world, warriors := generateMap(input)
	printMap(world, warriors)

	cnt := 0
	for true {
		if cnt > 10000 {
			break
		}

		battleRound(world, warriors)

		printMap(world, warriors)

		numberElfs, elfHP := countWarriors(warriors, elf)
		numberGoblins, goblinHp := countWarriors(warriors, goblin)
		if numberElfs == 0 || numberGoblins == 0 {
			log.Println("Ending the battle after " + strconv.Itoa(cnt) + " rounds with " + strconv.Itoa(numberElfs) + " Elves, and " +
				strconv.Itoa(numberGoblins) + " Goblins. Winning team has " +
				strconv.Itoa(elfHP+goblinHp) + " health")
			answer := (elfHP + goblinHp) * cnt
			log.Println("Step 1 answer: " + strconv.Itoa(answer))
			break
		}
		cnt++
	}
}

func step2(input string, elfAttackPower int) (int, int) {
	log.Println(fmt.Sprintf("Running step two with attack power %d", elfAttackPower))
	world, warriors := generateMap(input)
	//printMap(world, warriors)
	numberOfElves := 0
	for i, w := range warriors {
		if w.race == elf {
			warriors[i].attackPower = elfAttackPower
			numberOfElves++
		}
	}

	cnt := 0
	for true {
		if cnt > 10000 {
			break
		}

		battleRound(world, warriors)

		//printMap(world, warriors)

		numberElfs, elfHP := countWarriors(warriors, elf)
		numberGoblins, goblinHp := countWarriors(warriors, goblin)
		if numberElfs == 0 || numberGoblins == 0 {
			log.Println("Ending the battle after " + strconv.Itoa(cnt) + " rounds with " + strconv.Itoa(numberElfs) + " Elves, and " +
				strconv.Itoa(numberGoblins) + " Goblins. Winning team has " +
				strconv.Itoa(elfHP+goblinHp) + " health. Elf attack power: " + strconv.Itoa(elfAttackPower))
			answer := (elfHP + goblinHp) * cnt

			if numberElfs == numberOfElves {
				log.Println("Step 2 answer: " + strconv.Itoa(answer))
			}
			return numberElfs, numberOfElves
		}
		cnt++
	}
	return -1, -1
}

func main() {
	input := aoc_utils.FetchInput(2018, 15)
	step1(input)

	brute := false
	if brute {
		elfAttackPower := 4
		for true {
			elfIsWinningTeam, totalNumberOfElvesFromTheStart := step2(input, elfAttackPower)
			if elfIsWinningTeam < totalNumberOfElvesFromTheStart {
				elfAttackPower++
			} else {
				break
			}
		}
	} else {
		lowerLimit := 4
		upperLimit := 200
		elfAttackPower := upperLimit
		lowestAP := 200
		for true {
			elfIsWinningTeam, totalNumberOfElvesFromTheStart := step2(input, elfAttackPower)
			if elfIsWinningTeam == totalNumberOfElvesFromTheStart {
				if elfAttackPower < lowestAP {
					lowestAP = elfAttackPower
				}
				if upperLimit-lowerLimit == 1 {
					break
				} else {
					upperLimit = elfAttackPower
					elfAttackPower -= (upperLimit - lowerLimit) / 2
				}
			} else {
				if elfAttackPower == lowestAP-1 {
					log.Println("Lowest AP that doesn't kill any elfes: " + strconv.Itoa(lowestAP))
					break
				}
				lowerLimit = elfAttackPower
				if upperLimit-lowerLimit <= 0 {
					elfAttackPower *= 2
				} else {
					elfAttackPower += (upperLimit - lowerLimit) / 2
				}
			}

		}
	}
}
