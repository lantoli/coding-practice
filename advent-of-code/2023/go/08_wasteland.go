package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const filepath08 = "inputs/08_input.txt"

func main() {
	fmt.Println(calc08(false)) // 21883
	fmt.Println(calc08(true))  // 12833235391111
}

type node struct {
	l string
	r string
}

func calc08(ghost bool) int {
	f, _ := os.Open(filepath08)
	var pattern = regexp.MustCompile(`([A-Z0-9]+)`)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	inst := scanner.Text()
	scanner.Scan()
	nodes := make(map[string]node)
	for scanner.Scan() {
		read := pattern.FindAllString(scanner.Text(), -1)
		nodes[read[0]] = node{read[1], read[2]}
	}
	curs := make([]string, 0)
	if ghost {
		for cur := range nodes {
			if cur[2] == 'A' {
				curs = append(curs, cur)
			}
		}
	} else {
		curs = append(curs, "AAA")
	}
	nums := make([]int, 0)
	for i := range curs {
		total := 0
		for ghost && curs[i][2] != 'Z' || !ghost && curs[i] != "ZZZ" {
			if inst[total%len(inst)] == 'R' {
				curs[i] = nodes[curs[i]].r
			} else {
				curs[i] = nodes[curs[i]].l
			}
			total++
		}
		nums = append(nums, total)
	}
	return lcm(nums)
}

func lcm(nums []int) int {
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = lcm2(lcm, nums[i])
	}
	return lcm
}

func lcm2(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
