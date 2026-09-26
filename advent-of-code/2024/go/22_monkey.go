package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 19927218456 2189 (sample 37327623 23)
func main() {
	solve22a()
	solve22b()
}

var (
	file22 = "22_input.txt"
	n22    = 2000
)

func solve22a() {
	f, _ := os.Open("inputs/" + file22)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		for range n22 {
			num = next22(num)
		}
		total += num
	}
	fmt.Println(total)
}

func solve22b() {
	f, _ := os.Open("inputs/" + file22)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	changes := make([][]int, 0)
	prices := make([][]int, 0)
	for scanner.Scan() {
		change := make([]int, n22)
		price := make([]int, n22)
		num, _ := strconv.Atoi(scanner.Text())
		nump := num % 10
		for i := range n22 {
			next := next22(num)
			nextp := next % 10
			price[i], change[i] = nextp, nextp-nump
			num, nump = next, nextp
		}
		changes = append(changes, change)
		prices = append(prices, price)
	}
	fmt.Println(calc22b(changes, prices))
}

func next22(num int) int {
	num = ((num * (1 << 6)) ^ num) % 16777216
	num = ((num / (1 << 5)) ^ num) % 16777216
	num = ((num * (1 << 11)) ^ num) % 16777216
	return num
}

type st22 struct {
	a, b, c, d int
}

func calc22b(changes, prices [][]int) int {
	visited := make(map[st22]int)
	for n := range changes {
		change := changes[n]
		for i := 0; i < len(change)-3; i++ {
			st := st22{change[i], change[i+1], change[i+2], change[i+3]}
			if _, found := visited[st]; !found {
				visited[st] = count22(changes, prices, st)
			}
		}
	}
	total := 0
	for _, v := range visited {
		if v > total {
			total = v
		}
	}
	return total
}

func count22(changes, prices [][]int, st st22) int {
	total := 0
	for n := range changes {
		change := changes[n]
		price := prices[n]
		for i := 0; i < len(change)-3; i++ {
			if change[i] == st.a && change[i+1] == st.b && change[i+2] == st.c && change[i+3] == st.d {
				total += price[i+3]
				break
			}
		}
	}
	return total
}
