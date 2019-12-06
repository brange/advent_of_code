package main

import (
	"flag"
	"log"
	_ "strconv"
	"strings"
	_ "strings"

	_ "github.com/brange/advent_of_code/shared"
	_ "github.com/brange/advent_of_code/shared/math"
	"github.com/brange/advent_of_code/utils"
)

type orbitObject struct {
	name   string
	parent *orbitObject
}

func (o *orbitObject) distanceTo(parentName string) int {
	distance := 0
	for parent := o.parent; parent.parent != nil; parent = parent.parent {
		if parent.name == parentName {
			return distance
		}
		distance++
	}
	return -1
}

var YEAR = 2019
var DAY = 6

var debug bool

func main() {
	debugPtr := flag.Bool("debug", false, "Enable debug prints")
	flag.Parse()
	debug = *debugPtr
	if debug {
		log.Println("Enabling debug")
	}

	input := utils.FetchInput(YEAR, DAY)
	arr := strings.Split(input, "\n")

	objectsMap := make(map[string]*orbitObject)

	for _, a := range arr {
		parts := strings.Split(a, ")")
		parent, child := parts[0], parts[1]
		if debug {
			log.Println(child, "orbits", parent, " (", a, ")")
		}

		childPtr, foundChild := objectsMap[child]
		parentPtr, foundParent := objectsMap[parent]

		if !foundChild {
			childPtr = &orbitObject{name: child}
			objectsMap[child] = childPtr
		}

		if !foundParent {
			parentPtr = &orbitObject{name: parent}
			objectsMap[parent] = parentPtr
		}

		childPtr.parent = parentPtr
	}

	if debug {
		log.Println("objects:", objectsMap)
	}

	hash := 0
	for _, o := range objectsMap {
		child := o
		for parent := o.parent; parent != nil; parent = parent.parent {
			if debug {
				log.Println(child.name, "orbits", parent.name)
				child = parent
			}
			hash++
		}
	}

	log.Println("Part one:", hash)

	YOU, foundYou := objectsMap["YOU"]
	SAN, foundSan := objectsMap["SAN"]
	if !foundYou || !foundSan {
		log.Fatalln("Missing SAN or YOU")
	}

	var commonParent *orbitObject
l:
	for you := YOU.parent; you.parent != nil; you = you.parent {
		for san := SAN.parent; san.parent != nil; san = san.parent {
			if san.name == you.name {
				commonParent = san
				break l
			}
		}
	}
	if debug {
		log.Println("Found common parent ", commonParent.name)
	}

	part2 := YOU.distanceTo(commonParent.name) + SAN.distanceTo(commonParent.name)

	log.Println("Part two:", part2)
}
