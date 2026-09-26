package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// 188384 XX (sample 126384 XX)
func main() {
	solve("21_input.txt", line21, nil, func() int { return solve21(false) }, func() int { return solve21(true) })
}

var (
	input21 = make([]string, 0)
	cols21  = 3
	dirs21  = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
)

func line21(line string) {
	input21 = append(input21, line)

}

func solve21(b bool) int {
	if b {
		return -1
	}
	total := 0
	for _, codes := range input21 {
		total += calc21(codes)
	}
	return total
}

func calc21(code string) int {
	numpad := []rune{'7', '8', '9', '4', '5', '6', '1', '2', '3', 0, '0', 'A'}
	dirpad := []rune{0, '^', 'A', '<', 'v', '>'}

	nummap := initPaths21(numpad)
	dirmap := initPaths21(dirpad)

	ret := math.MaxInt

	dirs := nextdirs21(code, nummap)
	for i := range 2 {
		dirs = apply21(dirs, dirmap)
		fmt.Print(i, ", ", len(dirs), ", ", len21(dirs), "---")
	}
	ret = min(ret, code21(code)*len21(dirs))
	fmt.Println()
	return ret
}

func apply21(ins []string, pad map[rune]map[rune][]stp21) []string {
	ret := make([]string, 0)
	for _, in := range ins {
		ret = append(ret, nextdirs21(in, pad)...)
	}
	return ret
}

type st21 struct {
	pos  int
	path string
}

type stp21 struct {
	path string
	dirs string
}

func code21(code string) int {
	numStr := ""
	for _, elm := range code {
		if elm != 'A' {
			numStr += string(elm)
		}
	}
	num, _ := strconv.Atoi(numStr)
	return num
}

func len21(dirs []string) int {
	ret := math.MaxInt
	for _, dir := range dirs {
		ret = min(ret, len(dir))
	}
	return ret
}

func nextdirs21(code string, pad map[rune]map[rune][]stp21) []string {
	dirs := []string{""}
	prev := 'A'
	for _, next := range code {
		nextdirs := []string{}
		for _, path := range pad[prev][next] {
			for _, dir := range dirs {
				nextdirs = append(nextdirs, dir+path.dirs+"A")
			}
		}
		prev = next
		dirs = nextdirs
	}
	return dirs
}

func simplify21(dirs []string) []string {
	ends := make(map[string]string)
	for _, dir := range dirs {
		key := dir[0:1] + dir[len(dir)-1:len(dir)]
		val, found := ends[key]
		if !found || len(val) > len(dir) {
			ends[key] = dir
		}
	}
	ret := make([]string, 0)
	for _, end := range ends {
		ret = append(ret, end)
	}
	return ret
}

func initPaths21(grid []rune) map[rune]map[rune][]stp21 {
	rets := make(map[rune]map[rune][]stp21)
	rows := len(grid) / cols21
	for pos, button := range grid {
		dist := make(map[rune][]string)
		list := []st21{{pos, ""}}
		for len(list) > 0 {
			cur := list[0]
			list = list[1:]
			paths, _ := dist[grid[cur.pos]]
			if slices.Contains(paths, cur.path) {
				continue
			}
			dist[grid[cur.pos]] = append(paths, cur.path)
			x, y := cur.pos/cols21, cur.pos%cols21
			for _, dir := range dirs21 {
				xnew, ynew := x+dir[0], y+dir[1]
				posnew := xnew*cols21 + ynew
				if xnew >= 0 && xnew < rows && ynew >= 0 && ynew < cols21 && grid[posnew] != 0 && !strings.ContainsRune(cur.path, grid[posnew]) {
					list = append(list, st21{xnew*cols21 + ynew, cur.path + string(grid[posnew])})
				}
			}
		}
		ret := make(map[rune][]stp21)
		rets[button] = ret
		for k, v := range dist {
			minlen := math.MaxInt
			for _, path := range v {
				minlen = min(minlen, len(path))
			}
			for _, path := range v {
				if len(path) == minlen {
					dirs := ""
					rold := grid[pos]
					for _, rnew := range path {
						posold := slices.Index(grid, rold)
						posnew := slices.Index(grid, rnew)
						if posnew == posold+1 {
							dirs += ">"
						} else if posnew == posold-1 {
							dirs += "<"
						} else if posnew == posold+cols21 {
							dirs += "v"
						} else if posnew == posold-cols21 {
							dirs += "^"
						} else {
							panic("bad path")
						}
						rold = rnew
					}
					ret[k] = append(ret[k], stp21{path, dirs})
				}
			}
		}
	}
	return rets
}

// DELETE

func solve(inputFile string, processLine1, processLine2 func(string), ret1, ret2 func() int) {
	f1, _ := os.Open("inputs/" + inputFile)
	defer f1.Close()
	scanner1 := bufio.NewScanner(f1)
	for scanner1.Scan() {
		line := scanner1.Text()
		if processLine1 != nil {
			processLine1(line)
		}
	}
	fmt.Println(ret1())

	f2, _ := os.Open("inputs/" + inputFile)
	defer f2.Close()
	scanner2 := bufio.NewScanner(f2)
	for scanner2.Scan() {
		line := scanner2.Text()
		if processLine2 != nil {
			processLine2(line)
		}
	}
	fmt.Println(ret2())
}

func abs[T int](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func sign[T int](x T) T {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
}

func min[T int](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func max[T int](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func modpos[T int](a, b T) T {
	return (a%b + b) % b
}
