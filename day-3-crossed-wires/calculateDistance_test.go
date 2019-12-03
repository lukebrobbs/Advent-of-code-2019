package main

import (
	"reflect"
	"testing"
)

func TestIfDirectionIsRightShouldAddIncreasedXCoordinate(t *testing.T) {
	d := []string{}
	d = append(d, "R1")
	expected := [][]int{[]int{0, 0, 0}, []int{1, 0, 1}}
	actual := drawCoordinates(d)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	e := []string{}
	e = append(e, "R5")
	otherExpected := [][]int{[]int{0, 0, 0}, []int{1, 0, 1}, []int{2, 0, 2}, []int{3, 0, 3}, []int{4, 0, 4}, []int{5, 0, 5}}
	otherActual := drawCoordinates(e)
	if !reflect.DeepEqual(otherExpected, otherActual) {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}

func TestIfDirectionIsLeftShouldAddDecreasedXCoordinate(t *testing.T) {
	d := []string{}
	d = append(d, "L1")
	expected := [][]int{[]int{0, 0, 0}, []int{-1, 0, 1}}
	actual := drawCoordinates(d)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	e := []string{}
	e = append(e, "L5")
	otherExpected := [][]int{[]int{0, 0, 0}, []int{-1, 0, 1}, []int{-2, 0, 2}, []int{-3, 0, 3}, []int{-4, 0, 4}, []int{-5, 0, 5}}
	otherActual := drawCoordinates(e)
	if !reflect.DeepEqual(otherExpected, otherActual) {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}

func TestIfDirectionIsUpShouldAddIncreasedYCoordinate(t *testing.T) {
	d := []string{}
	d = append(d, "U1")
	expected := [][]int{[]int{0, 0, 0}, []int{0, 1, 1}}
	actual := drawCoordinates(d)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	e := []string{}
	e = append(e, "U5")
	otherExpected := [][]int{[]int{0, 0, 0}, []int{0, 1, 1}, []int{0, 2, 2}, []int{0, 3, 3}, []int{0, 4, 4}, []int{0, 5, 5}}
	otherActual := drawCoordinates(e)
	if !reflect.DeepEqual(otherExpected, otherActual) {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}

func TestIfDirectionIsDownShouldAddDecreasedYCoordinate(t *testing.T) {
	d := []string{}
	d = append(d, "D1")
	expected := [][]int{[]int{0, 0, 0}, []int{0, -1, 1}}
	actual := drawCoordinates(d)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	e := []string{}
	e = append(e, "D5")
	otherExpected := [][]int{[]int{0, 0, 0}, []int{0, -1, 1}, []int{0, -2, 2}, []int{0, -3, 3}, []int{0, -4, 4}, []int{0, -5, 5}}
	otherActual := drawCoordinates(e)
	if !reflect.DeepEqual(otherExpected, otherActual) {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}

func TestReturnsLowestCrossover(t *testing.T) {
	expected := 159
	actual, _ := calculateDistance([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"})
	if actual != expected {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	otherExpected := 135
	otherActual, _ := calculateDistance([]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"})
	if otherActual != otherExpected {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}

func TestShouldreturnLeastAmountOfCombinedSteps(t *testing.T) {
	expected := 610
	_, actual := calculateDistance([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"})
	if actual != expected {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	otherExpected := 410
	_, otherActual := calculateDistance([]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"})
	if otherActual != otherExpected {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}
