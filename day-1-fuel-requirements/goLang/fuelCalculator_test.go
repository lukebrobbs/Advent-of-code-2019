package main

import "testing"

func assertIsEqual(t *testing.T, a int, e int) {
	if a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}

func TestCorrectFuelIsReturnedForModule(t *testing.T) {
	assertIsEqual(t, module(14).calculateTotal(false), 2)
	assertIsEqual(t, module(1969).calculateTotal(false), 654)
	assertIsEqual(t, module(100756).calculateTotal(false), 33583)
}
func TestCalculateFuelShouldReturnTheTotalForAllModules(t *testing.T) {
	assertIsEqual(t, calculateFuelTotal([]module{14, 12}, false), 4)
	assertIsEqual(t, calculateFuelTotal([]module{-14, 12}, false), 2)
	assertIsEqual(t, calculateFuelTotal(getModules(), false), 3426455)
}

func TestCorrectFuelIsReturnedForModuleWhenFuelIncluded(t *testing.T) {
	assertIsEqual(t, module(14).calculateTotal(true), 2)
	assertIsEqual(t, module(1969).calculateTotal(true), 966)
	assertIsEqual(t, module(100756).calculateTotal(true), 50346)
}

func TestSouldCalculateTotalWhenFuelIncluded(t *testing.T) {
	assertIsEqual(t, calculateFuelTotal([]module{14, 12}, true), 4)
	assertIsEqual(t, calculateFuelTotal([]module{-14, 12}, true), 2)
	assertIsEqual(t, calculateFuelTotal([]module{14, 1969}, true), 968)
	assertIsEqual(t, calculateFuelTotal(getModules(), true), 5136807)
}
