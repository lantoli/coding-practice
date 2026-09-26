package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filepath09 = "inputs/09_input.txt"

func main() {
	fmt.Println(calc09(false)) // 1898776583
	fmt.Println(calc09(true))  // 1100
}

func calc09(backwards bool) int {
	f, _ := os.Open(filepath09)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		strNum := strings.Split(scanner.Text(), " ")
		nums := make([]int, len(strNum))
		for i := range strNum {
			nums[i], _ = strconv.Atoi(strNum[i])
		}
		total += getNums09(nums, backwards)
	}
	return total
}

func getNums09(nums []int, backwards bool) int {
	seqs := make([][]int, len(nums))
	seqs[0] = nums
	row := 0
	more := true
	for more {
		more = false
		row++
		seqs[row] = make([]int, len(nums)-row)
		for i := range seqs[row] {
			seqs[row][i] = seqs[row-1][i+1] - seqs[row-1][i]
			if seqs[row][i] != 0 {
				more = true
			}
		}
	}
	total := 0
	for row >= 0 {
		if backwards {
			total = seqs[row][0] - total
		} else {
			total += seqs[row][len(seqs[row])-1]
		}
		row--
	}
	return total
}
