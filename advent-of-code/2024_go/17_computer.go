package main

import (
	"fmt"
	"strings"
)

// 3,6,3,7,0,7,0,3,0 136904920099226 (sample 4,6,3,5,6,3,5,2,1,0 117440)
func main() {
	solve17a(true)
	solve17b()
}

func solve17a(isInput bool) {
	a, b, c := 729, 0, 0
	program := []int{0, 1, 5, 4, 3, 0}
	if isInput {
		a = 136904920099226 // 51064159
		program = []int{2, 4, 1, 5, 7, 5, 1, 6, 0, 3, 4, 6, 5, 5, 3, 0}
	}
	out := calc17(a, b, c, program)
	fmt.Println(strings.Trim(strings.Join(strings.Fields(strings.Trim(fmt.Sprint(out), "[]")), ","), " "))
}

func solve17b() {
	program := []int{2, 4, 1, 5, 7, 5, 1, 6, 0, 3, 4, 6, 5, 5, 3, 0}
	fmt.Println(calc17b(0, 0, program))
}

func calc17b(aa, matches int, program []int) int {
	aa *= 8
	for i := range 8 {
		a := aa + i
		out := calc17(a, 0, 0, program)
		if program[len(program)-1-matches] == out[len(out)-1-matches] {
			if len(program)-1-matches == 0 {
				return a
			}
			if ret := calc17b(a, matches+1, program); ret != -1 {
				return ret
			}
		}
	}
	return -1
}

// In each iteration an output is shown and a is divided by 8, values for b and c are reset.
func solve17c() {
	a := 136904920099226
	outs := make([]int, 0)
	count := 0
	for ; a > 0; count++ {
		b := (a % 8) ^ 5
		out := (b ^ 6 ^ (a / (1 << b))) % 8
		a /= 8
		outs = append(outs, out)
	}
	fmt.Println(outs)
}

func calc17(a, b, c int, program []int) []int {
	out := make([]int, 0)
	for ip := 0; ip < len(program); {
		opcode, literal := program[ip], program[ip+1]
		combo := literal
		if literal == 4 {
			combo = a
		} else if literal == 5 {
			combo = b
		} else if literal == 6 {
			combo = c
		}
		moveip := true
		switch opcode {
		case 0: // adv
			a /= 1 << combo
		case 1: // bxl
			b ^= literal
		case 2: // bst
			b = combo % 8
		case 3: // jnz
			if a != 0 {
				moveip = false
				ip = literal
			}
		case 4: // bxc
			b ^= c
		case 5: // out
			val := combo % 8
			out = append(out, val)
		case 6: // bdv
			b = a / (1 << combo)
		case 7: // cdv
			c = a / (1 << combo)
		default:
			panic(fmt.Sprintf("invalid opcode: %d", opcode))
		}
		if moveip {
			ip += 2
		}
	}
	return out
}
