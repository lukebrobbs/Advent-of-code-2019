package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func handleDirection(d string, cp *[][]int) {
	c := *cp
	a, err := strconv.Atoi(d[1:])
	newItem := []int{}

	if err != nil {
		fmt.Println("Unable to convert string to int")
	}
	for i := 0; i < a; i++ {
		last := c[len(c)-1]
		switch string(d[0]) {
		case "R":
			newItem = []int{last[0] + 1, last[1]}
		case "L":
			newItem = []int{last[0] - 1, last[1]}
		case "U":
			newItem = []int{last[0], last[1] + 1}
		case "D":
			newItem = []int{last[0], last[1] - 1}
		}
		c = append(c, newItem)

	}
	*cp = append(c)
}

func drawCoordinates(d []string) [][]int {
	c := [][]int{[]int{0, 0}}
	for _, direction := range d {
		handleDirection(direction, &c)
	}
	return c
}

func calculateDistance(wireOne []string, wireTwo []string) int {
	coordsOne := drawCoordinates(wireOne)
	coordsTwo := drawCoordinates(wireTwo)
	shortest := 0
	for _, c := range coordsTwo {
		for _, d := range coordsOne {
			if reflect.DeepEqual(c, d) {
				dist := int(math.Abs(float64(c[0])) + math.Abs(float64(c[1])))
				if shortest == 0 {
					shortest = dist
				}
				if dist < shortest {
					shortest = dist
				}
			}
		}
	}
	return shortest
}
