package main

import (
	"slices"
	"strconv"
	"strings"
)

// 5964 4719 (sample 143 123)
func main() {
	solve("05_input.txt", line05a, line05b, solve05a, solve05b)
}

var (
	order05         = make(map[int][]int)
	list05          = make([][]int, 0)
	readOrder05     = true
	listIncorrect05 = make([][]int, 0)
)

func line05a(line string) {
	if line == "" {
		readOrder05 = false
		return
	}
	if readOrder05 {
		nums := strings.Split(line, "|")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		order05[num1] = append(order05[num1], num2)
	} else {
		list := make([]int, 0)
		for _, numStr := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(numStr)
			list = append(list, num)
		}
		list05 = append(list05, list)
	}
}

func solve05a() int {
	total := 0
loop:
	for _, list := range list05 {
		pos := make(map[int]int)
		for i, num := range list {
			pos[num] = i
		}
		for prev, nexts := range order05 {
			for _, next := range nexts {
				pprev, ok1 := pos[prev]
				pnext, ok2 := pos[next]
				if ok1 && ok2 && pprev > pnext {
					listIncorrect05 = append(listIncorrect05, list)
					continue loop
				}
			}
		}
		total += list[len(list)/2]
	}
	return total
}

func line05b(line string) {
}

func solve05b() int {
	total := 0
	for _, list := range listIncorrect05 {
		for changed := true; changed; {
			changed = false
			for i := 1; i < len(list); i++ {
				if nexts, ok := order05[list[i]]; ok {
					for j := 0; j < i; j++ {
						if slices.Contains(nexts, list[j]) {
							list[i], list[j] = list[j], list[i]
							changed = true
							break
						}
					}
				}
			}
		}
		total += list[len(list)/2]
	}
	return total
}
