package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lukebrobbs/Advent-of-code-2019/utils"
)

type node struct {
	name     string
	children []*node
}

func splitString(s string) []string {
	return strings.Split(s, ")")
}

func getDirectChildren(m []string) []node {
	search := func(x string) int {
		return sort.Search(len(m), func(i int) bool { return splitString(x)[0] <= splitString(m[i])[0] })
	}
	for _, orbit := range m {
		i := search(orbit)
		fmt.Println(i)
	}
	return []node{}
}

func createTree(m []string) []node {

	return []node{}
}

func calculateOrbits(m []string) int {
	getDirectChildren(m)
	return 0
}

func main() {
	const (
		year = 2019
		day  = 6
		goal = 19690720
	)
	filePath := fmt.Sprintf("input.txt")
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := utils.ReadInput(filePath)

	solution1 := calculateOrbits(lines)

	fmt.Printf("%s\nSolution: \nPart1:  %v\n",
		header, solution1)
}
