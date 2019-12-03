package main

import (
	"fmt"
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
	switch string(d[0]) {
	case "R":
		for i := 0; i <= a; i++ {
			last := c[len(c)-1]
			newItem = []int{last[0] + 1, last[1]}
			c = append(c, newItem)
		}
	case "L":
		for i := 0; i <= a; i++ {
			last := c[len(c)-1]
			newItem = []int{last[0] - 1, last[1]}
			c = append(c, newItem)
		}
	case "U":
		for i := 0; i <= a; i++ {
			last := c[len(c)-1]
			newItem = []int{last[0], last[1] + 1}
			c = append(c, newItem)
		}
	case "D":
		for i := 0; i <= a; i++ {
			last := c[len(c)-1]
			newItem = []int{last[0], last[1] - 1}
			c = append(c, newItem)
		}
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

func calculateDistance(wireOne []string, wireTwo []string) {
	coordsOne := drawCoordinates(wireOne)
	coordsTwo := drawCoordinates(wireTwo)
	// fmt.Println(coordsOne)
	// fmt.Println(coordsTwo)

	x := [][]int{}
	for _, c := range coordsTwo {
		for _, d := range coordsOne {
			if reflect.DeepEqual(c, d) {
				fmt.Println("Tague", c)
				x = append(x, c)
			}
		}
	}
	fmt.Println(x)
}
