package main

import "math"

type module int

func (m module) calculateTotal(includeFuel bool) int {
	moduleFuel := int(math.Floor(float64(m)/3 - 2))
	if moduleFuel < 0 {
		moduleFuel = 0
	}
	if includeFuel == false {
		return moduleFuel
	}
	if moduleFuel > 0 {
		moduleFuel += module(moduleFuel).calculateTotal(true)
	}
	return moduleFuel
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
