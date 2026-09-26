package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Sample A", solve05("inputs/05_sample.txt", false)) // 3
	fmt.Println("A", solve05("inputs/05_input.txt", false))         // 821
	fmt.Println("Sample B", solve05("inputs/05_sample.txt", true))  // 14
	fmt.Println("B", solve05("inputs/05_input.txt", true))          // 344771884978261
}

func solve05(filename string, b bool) int {
	var (
		f, _            = os.Open(filename)
		scanner         = bufio.NewScanner(f)
		min, max, total = math.MaxInt64, 0, 0
		ranges          = make([][2]int, 0)
	)
	defer f.Close()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		r := [2]int{}
		fmt.Sscanf(line, "%d-%d", &r[0], &r[1])
		ranges = append(ranges, r)
		if r[0] < min {
			min = r[0]
		}
		if r[1] > max {
			max = r[1]
		}
	}
	if !b {
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			num, _ := strconv.Atoi(line)
			for _, r := range ranges {
				if num >= r[0] && num <= r[1] {
					total++
					break
				}
			}
		}
		return total
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	merged := make([][2]int, 0)
	for _, r := range ranges {
		if len(merged) == 0 || merged[len(merged)-1][1] < r[0] {
			merged = append(merged, r)
		} else {
			if r[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = r[1]
			}
		}
	}
	for _, r := range merged {
		total += r[1] - r[0] + 1
	}
	return total
}
