package main

func secureContainer(s int, f int) (count int) {
	for i := s; i <= f; i++ {
		if countDigits(i) == 6 {
			count++
		}
	}
	return count
}

func countDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
