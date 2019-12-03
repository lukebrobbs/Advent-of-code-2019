package main

import (
	"fmt"
	"strconv"
)

func handleDirection(d string, cp *[][]int) {
	c := *cp
	a, err := strconv.Atoi(d[1:])
	newItem := []int{}
	last := c[len(c)-1]

	if err != nil {
		fmt.Println("Unable to convert string to int")
	}
	switch string(d[0]) {
	case "R":
		newItem = []int{last[0] + a, last[1]}
	case "L":
		newItem = []int{last[0] - a, last[1]}
	case "U":
		newItem = []int{last[0], last[1] + a}
	case "D":
		newItem = []int{last[0], last[1] - a}
	}
	*cp = append(c, newItem)
}

func drawCoordinates(d []string) [][]int {
	c := [][]int{[]int{0, 0}}
	for _, direction := range d {
		handleDirection(direction, &c)
	}
	return c
}

func calculateDistance() {

}
