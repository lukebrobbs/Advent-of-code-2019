package main

type intCode []int

func (ic intCode) posOneAndTwo(i int) (int, int) {
	return ic[ic[i+1]], ic[ic[i+2]]
}
func (ic intCode) handleOne(i int) {
	first, second := ic.posOneAndTwo(i)
	ic[ic[i+3]] = first + second
}
func (ic intCode) handleTwo(i int) {
	first, second := ic.posOneAndTwo(i)
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
