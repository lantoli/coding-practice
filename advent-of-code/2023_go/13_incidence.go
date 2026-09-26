package main

import (
	"bufio"
	"fmt"
	"os"
)

const filepath13 = "inputs/13_input.txt"

func main() {
	fmt.Println(calc13(false)) // 33728
	fmt.Println(calc13(true))  // 28235
}

func calc13(smudge bool) int {
	f, _ := os.Open(filepath13)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	m := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			total += calc13Main(m, smudge)
			m = make([][]byte, 0)
		} else {
			m = append(m, []byte(line))
		}
	}
	if len(m) > 0 {
		total += calc13Main(m, smudge)
	}
	return total
}

func calc13Main(m [][]byte, smudge bool) int {
	prev := calc13Map(m, 0)
	if !smudge {
		return prev
	}
	rows, cols := len(m), len(m[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			swap(m, y, x)
			if ret := calc13Map(m, prev); ret != 0 {
				return ret
			}
			swap(m, y, x)
		}
	}
	return 0
}

func calc13Map(m [][]byte, prev int) int {
	rows, cols := len(m), len(m[0])
outer1:
	for x := 0; x+1 < cols; x++ {
		for inc := 0; x-inc >= 0 && x+1+inc < cols; inc++ {
			for y := 0; y < rows; y++ {
				if m[y][x-inc] != m[y][x+inc+1] {
					continue outer1
				}
			}
		}
		ret := x + 1
		if ret != prev {
			return ret
		}
	}
outer2:
	for y := 0; y+1 < rows; y++ {
		for inc := 0; y-inc >= 0 && y+1+inc < rows; inc++ {
			for x := 0; x < cols; x++ {
				if m[y-inc][x] != m[y+inc+1][x] {
					continue outer2
				}
			}
		}
		ret := (y + 1) * 100
		if ret != prev {
			return ret
		}
	}
	return 0
}

func swap(m [][]byte, y, x int) {
	if m[y][x] == '.' {
		m[y][x] = '#'
	} else {
		m[y][x] = '.'
	}
}
