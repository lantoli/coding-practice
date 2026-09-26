package main

import (
	"bufio"
	"fmt"
	"os"
)

const filepath11 = "inputs/11_input.txt"

func main() {
	fmt.Println(calc11(2))       // 9605127
	fmt.Println(calc11(1000000)) // 458191688761
}

func calc11(expand int) int {
	f, _ := os.Open(filepath11)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	image := make([][]byte, 0)
	for scanner.Scan() {
		image = append(image, []byte(scanner.Text()))
	}
	n := len(image)
	rows := make(map[int]bool)
	cols := make(map[int]bool)
outer:
	for y := range image {
		for x := range image {
			if image[y][x] == '#' {
				continue outer
			}
		}
		rows[y] = true
	}
outer2:
	for x := range image {
		for y := range image {
			if image[y][x] == '#' {
				continue outer2
			}
		}
		cols[x] = true
	}

	total := 0
	for i := 0; i < n*n; i++ {
		y1 := i / n
		x1 := i % n
		if image[y1][x1] == '#' {
			for j := i + 1; j < n*n; j++ {
				y2 := j / n
				x2 := j % n
				if image[y2][x2] == '#' {
					x1, x2 := ordered(x1, x2)
					y1, y2 := ordered(y1, y2)
					total += x2 - x1 + y2 - y1
					for y := y1 + 1; y < y2; y++ {
						if rows[y] {
							total += expand - 1
						}
					}
					for x := x1 + 1; x < x2; x++ {
						if cols[x] {
							total += expand - 1
						}
					}
				}
			}
		}
	}
	return total
}

func ordered(a, b int) (int, int) {
	if a <= b {
		return a, b
	}
	return b, a
}
