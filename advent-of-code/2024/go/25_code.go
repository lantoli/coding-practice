package main

import (
	"bufio"
	"fmt"
	"os"
)

// 2815 XX  (sample 3 XX)
func main() {
	solve25a()
}

var (
	file25 = "25_input.txt"
)

func solve25a() {
	f, _ := os.Open("inputs/" + file25)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	locks := make([][]int, 0)
	keys := make([][]int, 0)
	for scanner.Scan() {
		var sc []string
		for range 7 {
			sc = append(sc, scanner.Text())
			scanner.Scan()
		}
		isLock, pins := key25a(sc)
		if isLock {
			locks = append(locks, pins)
		} else {
			keys = append(keys, pins)
		}
	}
	total := 0
	for _, key := range keys {
	loop:
		for _, lock := range locks {
			for i := range 5 {
				if key[i]+lock[i] > 5 {
					continue loop
				}
			}
			total++
		}
	}
	fmt.Println(total)
}

func key25a(sc []string) (bool, []int) {
	ret := make([]int, 5)
	if sc[0][0] == '#' {
		for j := range 5 {
			for i := range 7 {
				if sc[i][j] == '.' {
					ret[j] = i - 1
					break
				}
			}
		}
		return true, ret
	}

	for j := range 5 {
		for i := range 7 {
			if sc[6-i][j] == '.' {
				ret[j] = i - 1
				break
			}
		}
	}
	return false, ret
}
