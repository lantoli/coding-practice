package main

import (
	"strconv"
	"strings"
)

// 3119088655389 264184041398847 (sample 3749 11387)
func main() {
	solve("07_input.txt", line07a, line07b, func() int { return total07a }, func() int { return total07b })
}

var (
	total07a = 0
	total07b = 0
)

func line07a(line string) {
	value, nums := getLineData07(line)
	if match07(value, nums, false) {
		total07a += value
	}
}

func line07b(line string) {
	value, nums := getLineData07(line)
	if match07(value, nums, true) {
		total07b += value
	}
}

func match07(value int, nums []int, extraOp bool) bool {
	if value < 0 {
		return false
	}
	if len(nums) == 1 {
		return nums[0] == value
	}
	copy1 := make([]int, len(nums)-1)
	copy(copy1, nums[1:])
	copy1[0] = nums[0] + nums[1]

	copy2 := make([]int, len(nums)-1)
	copy(copy2, nums[1:])
	copy2[0] = nums[0] * nums[1]

	ret := match07(value, copy1, extraOp) || match07(value, copy2, extraOp)
	if extraOp {
		copy3 := make([]int, len(nums)-1)
		copy(copy3, nums[1:])
		join := strconv.Itoa(nums[0]) + strconv.Itoa(nums[1])
		copy3[0], _ = strconv.Atoi(join)
		ret = ret || match07(value, copy3, extraOp)
	}
	return ret
}

func getLineData07(line string) (int, []int) {
	idx := strings.Index(line, ":")
	value, _ := strconv.Atoi(line[0:idx])
	strs := strings.Split(line[idx+2:], " ")
	nums := make([]int, 0)
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	return value, nums
}
