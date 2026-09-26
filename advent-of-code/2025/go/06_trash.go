package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Sample A", solve06a("inputs/06_sample.txt")) // 4277556
	fmt.Println("A", solve06a("inputs/06_input.txt"))         // 5524274308182
	fmt.Println("Sample B", solve06b("inputs/06_sample.txt")) // 3263827
	fmt.Println("B", solve06b("inputs/06_input.txt"))         // 8843673199391
}

func solve06a(filename string) int {
	var (
		f, _    = os.Open(filename)
		scanner = bufio.NewScanner(f)
		total   = 0
		numbers = make([][]string, 0)
		ops     = make([]string, 0)
	)
	defer f.Close()
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '*' || line[0] == '+' {
			ops = strings.Fields(line)
			break
		} else {
			numbers = append(numbers, strings.Fields(line))
		}
	}
	for x := range ops {
		sub := 0
		mul := ops[x] == "*"
		if mul {
			sub = 1
		}
		for y := range numbers {
			num, _ := strconv.Atoi(strings.TrimSpace(numbers[y][x]))
			if mul {
				sub *= num
			} else {
				sub += num
			}
		}
		total += sub
	}
	return total
}

func solve06b(filename string) int {
	var (
		f, _    = os.Open(filename)
		scanner = bufio.NewScanner(f)
		lines   = make([]string, 0)
		total   = 0
	)
	defer f.Close()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	nums := make([]int, 0)
	for x := len(lines[0]) - 1; x >= 0; x-- {
		num := 0
		for y := range len(lines) - 1 {
			if lines[y][x] != ' ' {
				num = num*10 + int(lines[y][x]-'0')
			}
		}
		nums = append(nums, num)
		op := lines[len(lines)-1][x]
		if op == ' ' {
			continue
		}
		x--
		sum, mul := 0, 1
		for _, num := range nums {
			sum += num
			mul *= num
		}
		if op == '+' {
			total += sum
		} else {
			total += mul
		}
		nums = nil
	}
	return total
}
