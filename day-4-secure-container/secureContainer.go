package main

import (
	"strconv"
)

func secureContainer(s int, f int) (count int) {
	for i := s; i <= f; i++ {
		if countDigits(i) == 6 && containsDoubles(i) && doesntDecrease(i) {
			count++
		}
	}
	return count
}

func containsDoubles(i int) (d bool) {
	m := ""
	s := strconv.Itoa(i)
	for i := 0; i < len(s)-1; i++ {
		if d == true && (string(s[i]) != string(s[i+1])) {
			return true
		}
		if string(s[i]) == string(s[i+1]) {
			m += string(s[i])
			d = len(m) == 1
		} else if len(m) > 1 {
			m = ""
		}
	}
	return d
}

func doesntDecrease(i int) (d bool) {
	s := strconv.Itoa(i)
	for i := 0; i < len(s)-1; i++ {
		if string(s[i]) <= string(s[i+1]) {
			d = true
		} else if string(s[i]) > string(s[i+1]) {
			d = false
			return d
		}
	}
	return d
}

func countDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
