package main

import (
	"strconv"
	"strings"
)

// 564 604 (sample 2 4)
func main() {
	solve("02_input.txt", line02a, line02b, func() int { return total02a }, func() int { return total02b })
}

var (
	total02a = 0
	total02b = 0
)

func line02a(line string) {
	nums := getList02(line)
	if isGood02(nums) {
		total02a++
	}
}

func line02b(line string) {
	initialNums := getList02(line)
	for i := range initialNums {
		nums := append([]int(nil), initialNums[0:i]...)
		nums = append(nums, initialNums[i+1:]...)
		if isGood02(nums) {
			total02b++
			return
		}
	}
}

func getList02(line string) []int {
	nums := make([]int, 0)
	for _, numStr := range strings.Fields(line) {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}
	return nums
}

func isGood02(nums []int) bool {
	s := sign(nums[1] - nums[0])
	for i := 1; i < len(nums); i++ {
		if sign(nums[i]-nums[i-1]) != s || abs(nums[i]-nums[i-1]) > 3 || nums[i] == nums[i-1] {
			return false
		}
	}
	return true
}
