package main

import "fmt"

type intCode []int

func (ic intCode) parameters(i int) (int, int) {
	return ic[ic[i+1]], ic[ic[i+2]]
}
func (ic intCode) handleOne(i int) {
	first, second := ic.parameters(i)
	ic[ic[i+3]] = first + second
}
func (ic intCode) handleTwo(i int) {
	first, second := ic.parameters(i)
	ic[ic[i+3]] = first * second
}

func (ic intCode) run() []int {
	for i := 0; i < len(ic); i += 4 {
		switch ic[i] {
		case 1:
			ic.handleOne(i)
		case 2:
			ic.handleTwo(i)
		case 99:
			break
		}
	}
	return ic
}

func partTwo(ic intCode, n int) {
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			p := append(ic[:0:0], ic...)
			p[1] = i
			p[2] = j
			for k := 0; k < len(p); k += 4 {
				switch p[k] {
				case 1:
					p.handleOne(k)
				case 2:
					p.handleTwo(k)
				case 99:
					break
				}
				if p[0] == n {
					fmt.Println(p[1], p[2])
				}
			}
		}
	}
}
