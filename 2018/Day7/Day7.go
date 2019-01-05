package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/brange/advent_of_code/utils"
)

type step struct {
	name         string
	dependencies []string
}

type worker struct {
	currentWork step
	startMinute int
	working     bool
}

func (s *step) addDependency(dep string) {
	s.dependencies = append(s.dependencies, dep)
}

func (s *step) removeDependency(dep string) {
	for index, d := range s.dependencies {
		if d == dep {
			s.dependencies[index] = s.dependencies[len(s.dependencies)-1]
			s.dependencies = s.dependencies[:len(s.dependencies)-1]

			break
		}
	}
}

func parseInput(input string) []step {
	arr := strings.Split(input, "\n")
	re := regexp.MustCompile(`Step (\w+) must be finished before step (\w+) can begin.`)
	var steps []step
	indexOf := func(name string) int {
		for index := range steps {
			if steps[index].name == name {
				return index
			}
		}
		return -1
	}
	for _, r := range arr {
		groups := re.FindAllStringSubmatch(r, 1)[0]
		name := groups[1]
		dependency := groups[2]

		indexForName := indexOf(name)
		indexforDep := indexOf(dependency)

		if indexForName == -1 {
			steps = append(steps, step{name: name})
		}
		if indexforDep == -1 {
			dep := step{name: dependency}
			dep.addDependency(name)
			steps = append(steps, dep)
		} else {
			steps[indexforDep].addDependency(name)
		}

	}

	return steps
}

func findFirstStepsWithoutDependencies(steps []step) (step, int) {
	var s step
	index := -1
	for i, s2 := range steps {
		if len(s2.dependencies) == 0 {
			if s.name == "" || s2.name < s.name {
				s = s2
				index = i
			}
		}
	}

	return s, index
}

func calculateTimeForStep(name string) int {
	return 60 + int(name[0]) - 64
}

func main() {
	input := utils.FetchInput(2018, 7)
	steps := parseInput(input)

	var resultStep1 string
	for len(steps) > 0 {
		first, index := findFirstStepsWithoutDependencies(steps)
		for index := range steps {
			steps[index].removeDependency(first.name)
		}
		resultStep1 += first.name

		if index >= 0 {
			steps[index] = steps[len(steps)-1]
			steps = steps[:len(steps)-1]
		}
	}
	fmt.Println("Answer step 1: " + resultStep1)

	steps = parseInput(input)
	workers := make([]worker, 5)
	minute := -1
	resultStep2 := ""
	workLeft := func() bool {
		if len(steps) == 0 {
			for _, worker := range workers {
				if worker.working {
					return true
				}
			}
			return false
		}
		return true
	}
	for workLeft() {
		minute++
		for i, worker := range workers {
			if worker.working {
				if minute == worker.startMinute+calculateTimeForStep(worker.currentWork.name) {
					// Done with this one.

					resultStep2 += worker.currentWork.name
					for index := range steps {
						steps[index].removeDependency(worker.currentWork.name)
					}

					workers[i].working = false
				}
			}
		}

		for i, worker := range workers {
			if !worker.working {
				first, index := findFirstStepsWithoutDependencies(steps)
				if index >= 0 {
					workers[i].startMinute = minute
					workers[i].currentWork = first
					workers[i].working = true

					steps[index] = steps[len(steps)-1]
					steps = steps[:len(steps)-1]

				}
			}
		}
	}
	fmt.Println("Answer step 2:", minute, "minutes")
	fmt.Println("Work order step 2: " + resultStep2)
}
