package main

import (
	"fmt"
)

func main() {
	total := calculateFuelTotal(getModules(), false)
	totalWithFuel := calculateFuelTotal(getModules(), true)
	fmt.Printf("The total is %v, and including fuel modules is %v", total, totalWithFuel)
}
