package main

import (
	"bytes"
	"fmt"
	"os"

	"go.openly.dev/pointy"
)

const filepath01 = "inputs/01_input.txt"

func main() {
	fmt.Println(calc01(filepath01, false)) // 55208
	fmt.Println(calc01(filepath01, true))  // 54578
}

func calc01(filepath string, includeLetters bool) int {
	var total int
	content, _ := os.ReadFile(filepath)
	lines := bytes.Split(content, []byte("\n"))
	for i := 0; i < len(lines)-1; i++ {
		var first *int
		var last int
		for pos := range lines[i] {
			if digit := getDigit(pos, lines[i], includeLetters); digit != nil {
				if first == nil {
					first = digit
				}
				last = *digit
			}
		}
		total += *first*10 + last
	}
	return total
}

func getDigit(pos int, line []byte, includeLetter bool) *int {
	char := line[pos]
	if char >= '0' && char <= '9' {
		return pointy.Int(int(char - '0'))
	}
	if !includeLetter {
		return nil
	}
	for n, word := range []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		if bytes.HasPrefix(line[pos:], []byte(word)) {
			return pointy.Int(n)
		}
	}
	return nil
}
