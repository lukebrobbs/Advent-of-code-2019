package main

import (
	"reflect"
	"testing"
)

func TestIntCodeOneAddsTogetherNumber(t *testing.T) {
	expected := []int{1, 1, 1, 2}
	actual := intCode{1, 1, 1, 3}.run()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	otherExpected := []int{2, 0, 0, 0}
	otherActual := intCode{1, 0, 0, 0}.run()
	if !reflect.DeepEqual(otherExpected, otherActual) {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}

func TestNinetyNineIntCodeExitsProgram(t *testing.T) {
	expected := []int{2, 4, 4, 5, 99, 9801}
	actual := intCode{2, 4, 4, 5, 99, 0}.run()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	otherExpected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	otherActual := intCode{1, 1, 1, 4, 99, 5, 6, 0, 99}.run()
	if !reflect.DeepEqual(otherExpected, otherActual) {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}

func TestTwoIntCodeMultipliesNumbers(t *testing.T) {
	expected := []int{2, 1, 1, 1}
	actual := intCode{2, 1, 1, 3}.run()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, instead got %v", expected, actual)
	}
	otherExpected := []int{2, 3, 0, 6, 99}
	otherActual := intCode{2, 3, 0, 3, 99}.run()
	if !reflect.DeepEqual(otherExpected, otherActual) {
		t.Fatalf("expected %v, instead got %v", otherExpected, otherActual)
	}
}
