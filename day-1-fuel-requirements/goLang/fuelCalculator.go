package main

import "math"

func calculateFuelForModule(n float64, includeFuel bool) int {
	moduleFuel := math.Floor(n/3 - 2)
	if includeFuel == false {
		return int(moduleFuel)
	}
	total := 0
	if moduleFuel > 0 {
		total += int(moduleFuel)
		total += calculateFuelForModule(moduleFuel, true)
	}
	return int(total)
}

func calculateFuelTotal(m []float64, includeFuel bool) int {
	total := 0
	for _, module := range m {
		if includeFuel {
			total += int(calculateFuelForModule(module, true))
		} else {
			total += int(calculateFuelForModule(module, false))
		}
	}
	return total
}
