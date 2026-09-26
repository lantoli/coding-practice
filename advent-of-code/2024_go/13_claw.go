package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	file13 = "13_input.txt"
	re13   = regexp.MustCompile(`(?:[\+=])(\d+)`)
)

// 29023 96787395375634 (sample 480 875318608908)
// takes minutes using parallelism, algorithm can surely be optimized to run in seconds in serial
func main() {
	solve13a(false)
	solve13a(true)
}

func solve13a(partB bool) {
	f, _ := os.Open("inputs/" + file13)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	num := 0
	ch := make(chan int)
	for scanner.Scan() {
		ax, ay := next13(scanner)
		bx, by := next13(scanner)
		px, py := next13(scanner)
		if partB {
			px += 10000000000000
			py += 10000000000000
		}
		go calc13a(ax, ay, bx, by, px, py, ch)
		num++
	}
	total := 0
	for range num {
		total += <-ch
	}
	fmt.Println(total)
}

func next13(scanner *bufio.Scanner) (int, int) {
	coord := re13.FindAllStringSubmatch(scanner.Text(), -1)
	scanner.Scan()
	x, _ := strconv.Atoi(coord[0][1])
	y, _ := strconv.Atoi(coord[1][1])
	return x, y
}

func calc13a(ax, ay, bx, by, px, py int, ch chan int) {
	ret := 0
	for b := min(px/bx, py/by); b >= 0; b-- {
		a := (px - b*bx) / ax
		if a*ax+b*bx == px && a*ay+b*by == py {
			ret = a*3 + b
			break
		}
	}
	ch <- ret
}
