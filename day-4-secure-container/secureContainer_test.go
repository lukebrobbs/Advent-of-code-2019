package main

import "testing"

func TestShouldCountNumbersWith6Digits(t *testing.T) {
	expected := secureContainer(1, 100000)
	expectedTwo := secureContainer(111111, 111122)
	if expected != 0 {
		t.Fatalf("Expected 1, instead recieved %v", expected)
	}
	if expectedTwo != 1 {
		t.Fatalf("Expected 3, instead recieved %v", expectedTwo)
	}
}

func TestShouldOnlyCountIntsWithTwoAdjacentDigitsTheSame(t *testing.T) {
	expected := secureContainer(111111, 111122)
	if expected != 1 {
		t.Fatalf("Expected 3, instead recieved %v", expected)
	}
}

func TestReturnedIntsShouldNotDecreaseFromRightToLeft(t *testing.T) {
	expected := secureContainer(223450, 223451)
	if expected != 0 {
		t.Fatalf("Expected 2, instead recieved %v", expected)
	}
}
