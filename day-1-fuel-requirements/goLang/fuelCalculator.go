package main

import "math"

type module int

func (m module) calculateTotal(includeFuel bool) int {
	moduleFuel := math.Floor(float64(m)/3 - 2)
	if includeFuel == false {
		return int(moduleFuel)
	}
	total := 0
	if moduleFuel > 0 {
		total += int(moduleFuel)
		total += module(moduleFuel).calculateTotal(true)
	}
	return int(total)
}

func calculateFuelTotal(m []module, includeFuel bool) int {
	total := 0
	for _, module := range m {
		if includeFuel {
			total += module.calculateTotal(true)
		} else {
			total += module.calculateTotal(false)
		}
	}
	return total
}
