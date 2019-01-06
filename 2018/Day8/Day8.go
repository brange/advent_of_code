package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brange/advent_of_code/utils"
)

type node struct {
	children []node
	metadata []int
}

func (n *node) addMetadata(val string) {
	metadata, _ := strconv.Atoi(val)
	n.metadata = append(n.metadata, metadata)
}

func (n *node) addChild(child node) {
	n.children = append(n.children, child)
}

func readData(input []string, index int) (node, int) {
	var n node

	numberChildren, _ := strconv.Atoi(input[index])
	numberMetadata, _ := strconv.Atoi(input[index+1])
	index += 2
	for i := 0; i < numberChildren; i++ {
		child, _index := readData(input, index)
		n.addChild(child)
		index = _index
	}
	for i := 0; i < numberMetadata; i++ {
		n.addMetadata(input[index])
		index++
	}

	return n, index
}

func sumMetadata(n node) int {
	sum := 0

	for _, c := range n.children {
		sum += sumMetadata(c)
	}

	for _, c := range n.metadata {
		sum += c
	}
	return sum
}

func calcStep2(n node) int {
	sum := 0

	if len(n.children) == 0 {
		for _, c := range n.metadata {
			sum += c
		}
	} else {
		for _, c := range n.metadata {
			if c <= len(n.children) {
				sum += calcStep2(n.children[c-1])
			}
		}
	}

	return sum
}

func main() {
	input := utils.FetchInput(2018, 8)

	rootNode, _ := readData(strings.Split(input, " "), 0)

	fmt.Println("Answer step 1:", sumMetadata(rootNode))
	fmt.Println("Answer step 2:", calcStep2(rootNode))
}
