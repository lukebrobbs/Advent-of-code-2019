package main

import "testing"

func TestShouldCountNumbersWith6Digits(t *testing.T) {
	expected := secureContainer(1, 100000)
	expectedTwo := secureContainer(1, 100002)
	if expected != 1 {
		t.Fatalf("Expected 1, instead recieved %v", expected)
	}
	if expectedTwo != 3 {
		t.Fatalf("Expected 3, instead recieved %v", expected)
	}
}

func TestReturnedValueShouldBeWithinRangeOfGivenInputs(t *testing.T) {}

func TestReturnedIntsShouldContainTwoAdjacentDigits(t *testing.T) {

}

func TestReturnedIntsShouldNotDecreaseFromRightToLeft(t *testing.T) {}
