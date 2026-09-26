package main

import (
	"strconv"
	"strings"
)

// 220999 261936432123724 (sample 55312 65601038650482)
func main() {
	solve("11_input.txt", line11a, nil, solve11a, solve11b)
}

var (
	input11a = make([]int, 0)
)

func line11a(line string) {
	for _, numStr := range strings.Fields(line) {
		num, _ := strconv.Atoi(numStr)
		input11a = append(input11a, num)
	}
}

func solve11a() int {
	oldList := input11a
	for range 25 {
		newList := make([]int, 0, len(oldList))
		for _, num := range oldList {
			if num == 0 {
				newList = append(newList, 1)
				continue
			}
			a, b := splitEven(num)
			if a >= 0 {
				newList = append(newList, a, b)
				continue
			}
			newList = append(newList, num*2024)
		}
		oldList = newList
	}
	return len(oldList)
}

func solve11b() int {
	oldMap := make(map[int]int)
	for _, num := range input11a {
		oldMap[num]++
	}
	for range 75 {
		newMap := make(map[int]int)
		for num, count := range oldMap {
			if num == 0 {
				newMap[1] += count
				continue
			}
			a, b := splitEven(num)
			if a >= 0 {
				newMap[a] += count
				newMap[b] += count
				continue
			}
			newMap[num*2024] += count
		}
		oldMap = newMap
	}
	total := 0
	for _, count := range oldMap {
		total += count
	}
	return total
}

func splitEven(num int) (int, int) {
	digits := 0
	for n := num; n > 0; n /= 10 {
		digits++
	}
	if digits%2 == 1 {
		return -1, -1
	}
	n := 0
	for num > 0 {
		n = n*10 + num%10
		num /= 10
	}
	a, b := 0, 0
	for range digits / 2 {
		a = a*10 + n%10
		n /= 10
	}
	for range digits / 2 {
		b = b*10 + n%10
		n /= 10
	}
	return a, b
}
