package main

import (
	"log"
	_ "strconv"
	_ "strings"

	_ "github.com/brange/advent_of_code/shared"
	_ "github.com/brange/advent_of_code/shared/math"
	"github.com/brange/advent_of_code/utils"
)

var YEAR = 2019
var DAY = -1

func main() {
	input := utils.FetchInput(YEAR, DAY)

	log.Println("Got input", input)
}
