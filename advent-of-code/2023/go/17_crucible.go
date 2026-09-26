package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const filepath17 = "inputs/17_input.txt"

func main() {
	ultra = false
	fmt.Println(calc17()) // 861
	ultra = true
	fmt.Println(calc17()) // 1037
}

var (
	grid     []string
	n        byte
	ret17    int
	global17 map[st17]int
	ultra    bool
)

func calc17() int {
	f, _ := os.Open(filepath17)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	grid = make([]string, 0)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if len(grid) != len(grid[0]) {
		panic("grid is not square")
	}
	n = byte(len(grid))
	ret17 = math.MaxInt
	global17 = make(map[st17]int)
	calc17Main(grid)
	return ret17
}

func calc17Main(grid []string) {
	calc17Rec(grid, 0, &st17{0, 0, EAST, 0})
	calc17Rec(grid, 0, &st17{0, 0, SOUTH, 0})
}

func calc17Rec(grid []string, heat int, last *st17) {
	if value, exist := global17[*last]; heat >= ret17 || (exist && value <= heat) {
		return
	}
	global17[*last] = heat
	if last.y == n-1 && last.x == n-1 && (!ultra || last.single >= 4) {
		ret17 = min(ret17, heat)
		return
	}
	candidates := make([]*st17, 0)
	copy := &st17{last.y, last.x, last.dir, last.single}
	if copy.move() {
		candidates = append(candidates, copy)

	}
	copy = &st17{last.y, last.x, last.dir, last.single}
	if copy.moveLeft() {
		candidates = append(candidates, copy)
	}
	copy = &st17{last.y, last.x, last.dir, last.single}
	if copy.moveRight() {
		candidates = append(candidates, copy)
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].x+candidates[i].y >= candidates[j].x+candidates[j].y
	})
	for _, c := range candidates {
		calc17Rec(grid, heat+int(grid[c.y][c.x])-'0', c)
	}
}

type st17 struct {
	y      byte
	x      byte
	dir    byte
	single byte
}

func (s *st17) move() bool {
	var max byte = 3
	if ultra {
		max = 10
	}
	switch s.dir {
	case NORTH:
		s.y--
	case EAST:
		s.x++
	case SOUTH:
		s.y++
	case WEST:
		s.x--
	}
	s.single++
	return s.single <= max && s.y >= 0 && s.y < n && s.x >= 0 && s.x < n
}

func (s *st17) moveLeft() bool {
	if ultra && s.single < 4 {
		return false
	}
	s.dir = (s.dir - 1 + 4) % 4
	s.single = 0
	return s.move()
}

func (s *st17) moveRight() bool {
	if ultra && s.single < 4 {
		return false
	}
	s.dir = (s.dir + 1) % 4
	s.single = 0
	return s.move()
}

const (
	NORTH = 0
	EAST  = 1
	SOUTH = 2
	WEST  = 3
)
