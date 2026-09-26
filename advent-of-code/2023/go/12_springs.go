package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const filepath12 = "inputs/12_input.txt"

func main() {
	fmt.Println(calc12(1)) // 7025
	fmt.Println(calc12(5)) // XXX (it takes a few hours with goroutines, there is surely a faster algorithm)
}

func calc12(repeat int) int {
	f, _ := os.Open(filepath12)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	lines := 0
	ch := make(chan int)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		nums := make([]int, 0)
		for _, numStr := range strings.Split(line[1], ",") {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		go calc12Line(ch, line[0], nums, repeat)
		lines++
	}
	for lines > 0 {
		total += <-ch
		lines--
	}
	return total
}

func calc12Line(ch chan int, field string, nums []int, repeat int) {
	f := ""
	n := make([]int, 0)
	for i := 0; i < repeat; i++ {
		f += field
		n = append(n, nums...)
		if i != repeat-1 {
			f += "?"
		}
	}
	f = regexp.MustCompile(`\.+`).ReplaceAllString(f+".", ".")
	ch <- calc12Rec([]byte(f), n, 0, 0, 0)
}

func calc12Rec(field []byte, nums []int, posField, posNums int, failing int) int {
	if posNums == len(nums) || posField == len(field) {
		for posField != len(field) && field[posField] != '#' {
			posField++
		}
		if posNums == len(nums) && posField == len(field) {
			return 1
		} else {
			return 0
		}
	}
	if field[posField] == '.' {
		if failing > 0 {
			if failing != nums[posNums] {
				return 0
			}
			posNums++
		}
		ret := calc12Rec(field, nums, posField+1, posNums, 0)
		return ret
	} else if field[posField] == '#' {
		ret := calc12Rec(field, nums, posField+1, posNums, failing+1)
		return ret
	}
	field[posField] = '#'
	total := calc12Rec(field, nums, posField, posNums, failing)
	field[posField] = '.'
	total += calc12Rec(field, nums, posField, posNums, failing)
	field[posField] = '?'
	return total
}
