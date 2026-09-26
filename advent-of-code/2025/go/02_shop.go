package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Sample A", solve02("inputs/02_sample.txt", false)) // 1227775554
	fmt.Println("A", solve02("inputs/02_input.txt", false))         // 23534117921
	fmt.Println("Sample B", solve02("inputs/02_sample.txt", true))  // 4174379265
	fmt.Println("B", solve02("inputs/02_input.txt", true))          // 31755323497
}

func solve02(filename string, b bool) int {
	var (
		data, _ = os.ReadFile(filename)
		total   = 0
	)
	for _, pair := range strings.Split(strings.TrimSpace(string(data)), ",") {
		parts := strings.Split(pair, "-")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		for num := num1; num <= num2; num++ {
			if (!b && isRepeated2(num)) || (b && isRepeatedN(num)) {
				total += num
			}
		}
	}
	return total
}

func isRepeated2(num int) bool {
	str := strconv.Itoa(num)
	if len(str)%2 != 0 {
		return false
	}
	return str[0:len(str)/2] == str[len(str)/2:]
}

func isRepeatedN(num int) bool {
	str := strconv.Itoa(num)
outer:
	for n := 1; n <= len(str)/2; n++ {
		if len(str)%n != 0 {
			continue
		}
		part := str[0:n]
		for i := n; i < len(str); i += n {
			if str[i:i+n] != part {
				continue outer
			}
		}
		return true
	}
	return false
}
