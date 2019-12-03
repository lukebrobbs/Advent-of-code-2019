package main

import (
	"reflect"
	"testing"
)

func TestIfDirectionIsRightShouldAddIncreasedXCoordinate(t *testing.T) {
	d := []string{}
	d = append(d, "R1")
	expected := [][]int{[]int{0, 0}, []int{1, 0}}
	actual := drawCoordinates(d)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	e := []string{}
	e = append(e, "R5")
	otherExpected := [][]int{[]int{0, 0}, []int{5, 0}}
	otherActual := drawCoordinates(e)
	if !reflect.DeepEqual(otherExpected, otherActual) {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}

func TestIfDirectionIsLeftShouldAddDecreasedXCoordinate(t *testing.T) {
	d := []string{}
	d = append(d, "L1")
	expected := [][]int{[]int{0, 0}, []int{-1, 0}}
	actual := drawCoordinates(d)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	e := []string{}
	e = append(e, "L5")
	otherExpected := [][]int{[]int{0, 0}, []int{-5, 0}}
	otherActual := drawCoordinates(e)
	if !reflect.DeepEqual(otherExpected, otherActual) {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}

func TestIfDirectionIsUpShouldAddIncreasedYCoordinate(t *testing.T) {
	d := []string{}
	d = append(d, "U1")
	expected := [][]int{[]int{0, 0}, []int{0, 1}}
	actual := drawCoordinates(d)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	e := []string{}
	e = append(e, "U5")
	otherExpected := [][]int{[]int{0, 0}, []int{0, 5}}
	otherActual := drawCoordinates(e)
	if !reflect.DeepEqual(otherExpected, otherActual) {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}

func TestIfDirectionIsDownShouldAddDecreasedYCoordinate(t *testing.T) {
	d := []string{}
	d = append(d, "D1")
	expected := [][]int{[]int{0, 0}, []int{0, -1}}
	actual := drawCoordinates(d)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	e := []string{}
	e = append(e, "D5")
	otherExpected := [][]int{[]int{0, 0}, []int{0, -5}}
	otherActual := drawCoordinates(e)
	if !reflect.DeepEqual(otherExpected, otherActual) {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}
