package main

import "math"

func calculateFuel(n float64) float64 {
	moduleFuel := math.Floor(n/3 - 2)
	return moduleFuel
}
