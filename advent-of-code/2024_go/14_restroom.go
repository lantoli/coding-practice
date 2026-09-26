package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	cols14, rows14 = 101, 103
	//cols14, rows14 = 11, 7 // sample
	px14, py14, vx14, vy14 []int
)

// 215476074 6285 (sample 12)
func main() {
	solve("14_input.txt", line14, nil, solve14a, solve14b)
}

func solve14a() int {
	var x, y []int
	for i := range px14 {
		x = append(x, modpos(px14[i]+vx14[i]*100, cols14))
		y = append(y, modpos(py14[i]+vy14[i]*100, rows14))
	}
	q1, q2, q3, q4 := quad14(x, y)
	return q1 * q2 * q3 * q4
}

func solve14b() int {
	x := append([]int(nil), px14...)
	y := append([]int(nil), py14...)
	minNum, minSec := math.MaxInt64, 0
	for seconds := 1; seconds <= 1_000_000; seconds++ {
		for i := range x {
			x[i] = modpos(x[i]+vx14[i], cols14)
			y[i] = modpos(y[i]+vy14[i], rows14)
		}
		q1, q2, q3, q4 := quad14(x, y)
		if minNum > q1*q2*q3*q4 {
			minNum = q1 * q2 * q3 * q4
			minSec = seconds
		}
	}
	return minSec
}

func quad14(x, y []int) (q1, q2, q3, q4 int) {
	midcol, midrow := cols14/2, rows14/2
	for i := range x {
		if x[i] < midcol && y[i] < midrow {
			q1++
		} else if x[i] > midcol && y[i] < midrow {
			q2++
		} else if x[i] < midcol && y[i] > midrow {
			q3++
		} else if x[i] > midcol && y[i] > midrow {
			q4++
		}
	}
	return q1, q2, q3, q4
}

func line14(line string) {
	px, py, vx, vy := next14(line)
	px14 = append(px14, px)
	py14 = append(py14, py)
	vx14 = append(vx14, vx)
	vy14 = append(vy14, vy)
}

func next14(line string) (px, py, vx, vy int) {
	l := strings.Fields(line)
	p := strings.Split(strings.Split(l[0], "=")[1], ",")
	v := strings.Split(strings.Split(l[1], "=")[1], ",")
	px, _ = strconv.Atoi(p[0])
	py, _ = strconv.Atoi(p[1])
	vx, _ = strconv.Atoi(v[0])
	vy, _ = strconv.Atoi(v[1])
	return px, py, vx, vy
}

func writeGrid(filename string, grid [][]int) {
	file, err := os.Create(filename)
	if err != nil {
		panic("error creating file: " + filename)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, row := range grid {
		for _, val := range row {
			if val == 0 {
				fmt.Fprintf(writer, ".")
			} else {
				fmt.Fprintf(writer, "%d", val)
			}
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
