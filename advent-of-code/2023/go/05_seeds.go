package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const filepath05 = "inputs/05_input.txt"

func main() {
	fmt.Println(calc05(false)) // 175622908
	fmt.Println(calc05(true))  // 5200543
}

type elm struct {
	pos int
	len int
}

func calc05(hasRange bool) int {
	f, _ := os.Open(filepath05)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	seeds := make([]elm, 0)
	nums := getNums(scanner.Text())
	if hasRange {
		for i := 0; i < len(nums); i += 2 {
			seeds = append(seeds, elm{nums[i], nums[i+1]})
		}
	} else {
		for i := range nums {
			seeds = append(seeds, elm{nums[i], 1})
		}
	}
	next := make([]elm, 0)
	for scanner.Scan() {
		to := getNums(scanner.Text())
		if len(to) == 0 {
			seeds = append(next, seeds...)
			next = make([]elm, 0)
			continue
		}
		dest := to[0]
		pos := to[1]
		length := to[2]
		for i := 0; i < len(seeds); i++ {
			left := max(seeds[i].pos, pos)
			right := min(seeds[i].pos+seeds[i].len-1, pos+length-1)
			newLen := right - left + 1
			if newLen > 0 {
				next = append(next, elm{dest - pos + left, newLen})
				seeds[i].len = 0
				if left > seeds[i].pos {
					next = append(next, elm{seeds[i].pos, left - seeds[i].pos})
				}
				if right < seeds[i].pos+seeds[i].len-1 {
					next = append(next, elm{right + 1, seeds[i].pos + seeds[i].len - right - 1})
				}
			}
		}
	}
	min := math.MaxInt32
	for _, seed := range append(next, seeds...) {
		if seed.len > 0 && seed.pos < min {
			min = seed.pos
		}
	}
	return min
}

var pattern = regexp.MustCompile(`(\d+)`)

func getNums(line string) []int {
	matches := pattern.FindAllStringSubmatch(line, -1)
	nums := make([]int, len(matches))
	for i, match := range matches {
		nums[i], _ = strconv.Atoi(match[0])
	}
	return nums
}
