package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lukebrobbs/Advent-of-code-2019/utils"
)

func makeDigits(n int) []int {
	digits := make([]int, 5)
	for i := 0; i < 5; i++ {
		digits[4-i] = int(n % 10)
		n /= 10
	}
	return digits
}
func runCode(origCode []int, input int) []int {

	code := make([]int, len(origCode))
	copy(code, origCode)
	result := make([]int, 0)
	var opcode, m1, m2 int
	ptr := 0
	running := true
	for running {
		op := makeDigits(code[ptr])
		m1, m2 = op[2], op[1]
		p1 := 0
		p2 := 0
		opcode = 10*op[3] + op[4]
		if opcode == 1 || opcode == 2 || (opcode >= 4 && opcode <= 8) {
			if m1 == 0 {
				p1 = code[code[ptr+1]]
			} else {
				p1 = code[ptr+1]
			}
		}
		if opcode == 1 || opcode == 2 || (opcode > 4 && opcode <= 8) {
			if m2 == 0 {
				p2 = code[code[ptr+2]]
			} else {
				p2 = code[ptr+2]
			}

		}
		switch opcode {
		case 1:

			code[code[ptr+3]] = p1 + p2
			ptr += 4
		case 2:

			code[code[ptr+3]] = p1 * p2
			ptr += 4
		case 3:
			code[code[ptr+1]] = input
			ptr += 2
		case 4:
			result = append(result, p1)
			ptr += 2
		case 5:
			if p1 != 0 {
				ptr = p2
			} else {
				ptr += 3
			}
		case 6:
			if p1 == 0 {
				ptr = p2
			} else {
				ptr += 3
			}
		case 7:
			if p1 < p2 {
				code[code[ptr+3]] = 1
			} else {
				code[code[ptr+3]] = 0
			}
			ptr += 4
		case 8:
			if p1 == p2 {
				code[code[ptr+3]] = 1
			} else {
				code[code[ptr+3]] = 0
			}
			ptr += 4
		case 99:
			running = false
		default:
			fmt.Println("ERROR:", opcode)
			break
		}

	}
	return result
}

func main() {

	const (
		year = 2019
		day  = 5
		goal = 19690720
	)

	filePath := fmt.Sprintf("input.txt")
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := utils.ReadInput(filePath)

	var (
		solution1, solution2 int64
	)

	ll := strings.Split(lines[0], ",")

	code := make([]int, len(ll))
	for i := range ll {
		code[i], _ = strconv.Atoi(ll[i])
	}

	part1 := runCode(code, 1)
	part2 := runCode(code, 5)
	for _, i := range part1 {
		if i != 0 {
			solution1 = int64(i)
			break
		}
	}
	for _, i := range part2 {
		if i != 0 {
			solution2 = int64(i)
			break
		}
	}
	fmt.Printf("%s\nSolution:\nPart1:\t%v\nPart2:\t%v",
		header, solution1, solution2)
}
