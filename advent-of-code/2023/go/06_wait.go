package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const filepath06 = "inputs/06_input.txt"

func main() {
	fmt.Println(calc06(false)) // 625968
	fmt.Println(calc06(true))  // 43663323
}

func calc06(oneRace bool) int {
	total := 1
	var pattern = regexp.MustCompile(`(\d+)`)
	f, _ := os.Open(filepath06)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	times := pattern.FindAllStringSubmatch(scanner.Text(), -1)
	scanner.Scan()
	distances := pattern.FindAllStringSubmatch(scanner.Text(), -1)
	if oneRace {
		var strTime, strDistance string
		for i := range times {
			strTime += times[i][0]
			strDistance += distances[i][0]
		}
		times = [][]string{{strTime}}
		distances = [][]string{{strDistance}}
	}
	for i := range times {
		time, _ := strconv.Atoi(times[i][0])
		distance, _ := strconv.Atoi(distances[i][0])
		subtotal := 0
		for j := 1; j < time; j++ {
			if (time-j)*j > distance {
				subtotal++
			}
		}
		total *= subtotal
	}
	return total
}
