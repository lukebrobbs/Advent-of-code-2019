package main

type intCode []int

func (ic intCode) posOneAndTwo(i int) (int, int) {
	return ic[ic[i+1]], ic[ic[i+2]]
}

func (ic intCode) run() []int {
	for i := 0; i < len(ic); i += 4 {
		switch ic[i] {
		case 1:
			first, second := ic.posOneAndTwo(i)
			ic[ic[i+3]] = first + second
		case 2:
			first, second := ic.posOneAndTwo(i)
			ic[ic[i+3]] = first * second
		case 99:
			break
		default:
		}
	}
	return ic
}
