package main

import (
	"bytes"
	"fmt"
	"os"
)

const filepath03 = "inputs/03_input.txt"

func main() {
	fmt.Println(calc03a(filepath03)) // 560670
	fmt.Println(calc03b(filepath03)) // 91622824
}

func calc03a(filepath string) int {
	total := 0
	content, _ := os.ReadFile(filepath)
	engine := bytes.Split(content, []byte("\n"))
	rows := len(engine) - 1
	cols := len(engine[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			n := 0
			isNumber := false
			for col < cols && engine[row][col] >= '0' && engine[row][col] <= '9' {
				n = n*10 + int(engine[row][col]-'0')
				isNumber = isNumber || adjSymbol(engine, rows, cols, row, col)
				col++
			}
			if isNumber {
				total += n
			}
		}
	}
	return total
}

func calc03b(filepath string) int {
	total := 0
	content, _ := os.ReadFile(filepath)
	engine := bytes.Split(content, []byte("\n"))
	rows := len(engine) - 1
	cols := len(engine[0])
	for row := 0; row < rows; row++ {
	outer:
		for col := 0; col < cols; col++ {
			if engine[row][col] == '*' {
				var a, b int
				inc := []int{-1, 0, 1}
				for _, yinc := range inc {
					lastPos := 0
					for _, xinc := range inc {
						if (xinc != 0 || yinc != 0) && col+xinc > lastPos {
							var n int
							n, lastPos = getNumber(engine, rows, cols, row+yinc, col+xinc)
							if n > 0 {
								if a == 0 {
									a = n
								} else if b == 0 {
									b = n
								} else {
									continue outer
								}
							}
						}
					}
				}
				total += a * b
			}
		}
	}
	return total
}

func getNumber(engine [][]byte, rows, cols, row, col int) (int, int) {
	if engine[row][col] < '0' || engine[row][col] > '9' {
		return 0, col
	}
	n := 0
	for col > 0 && engine[row][col-1] >= '0' && engine[row][col-1] <= '9' {
		col--
	}
	for col < cols && engine[row][col] >= '0' && engine[row][col] <= '9' {
		n = n*10 + int(engine[row][col]-'0')
		col++
	}
	return n, col
}

func adjSymbol(engine [][]byte, rows, cols, row, col int) bool {
	inc := []int{-1, 0, 1}
	for _, yinc := range inc {
		for _, xinc := range inc {
			if xinc != 0 || yinc != 0 {
				y := row + yinc
				x := col + xinc
				if y >= 0 && y < rows && x >= 0 && x < cols && engine[y][x] != '.' && (engine[y][x] < '0' || engine[y][x] > '9') {
					return true
				}
			}
		}
	}
	return false
}
