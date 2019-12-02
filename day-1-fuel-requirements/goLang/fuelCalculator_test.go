package main

import "testing"

func assertIsEqual(t *testing.T, a int, e int) {
	if a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}

func TestCorrectFuelIsReturnedForModule(t *testing.T) {
	assertIsEqual(t, calculateFuelForModule(14, false), 2)
	assertIsEqual(t, calculateFuelForModule(1969, false), 654)
	assertIsEqual(t, calculateFuelForModule(100756, false), 33583)
}
func TestCalculateFuelShouldReturnTheTotalForAllModules(t *testing.T) {
	assertIsEqual(t, calculateFuelTotal([]float64{14, 12}, false), 4)
	assertIsEqual(t, calculateFuelTotal(getModules(), false), 3426455)
}

func TestCorrectFuelIsReturnedForModuleWhenFuelIncluded(t *testing.T) {
	assertIsEqual(t, calculateFuelForModule(14, true), 2)
	assertIsEqual(t, calculateFuelForModule(1969, true), 966)
	assertIsEqual(t, calculateFuelForModule(100756, true), 50346)
}

func TestSouldCalculateFuelForFuelmodules(t *testing.T) {
	assertIsEqual(t, calculateFuelTotal([]float64{14, 12}, true), 4)
	assertIsEqual(t, calculateFuelTotal([]float64{14, 1969}, true), 968)
	assertIsEqual(t, calculateFuelTotal(getModules(), true), 5136807)
}
