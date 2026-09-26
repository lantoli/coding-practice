package main

import (
	"fmt"
	"math"
	"strings"
)

// 102504 535 (sample 7036 11048, 45 64)
func main() {
	solve("16_input.txt", line16, nil, func() int { return solve16(false) }, func() int { return solve16(true) })
}

var (
	input16                        = make([][]rune, 0)
	rows16, cols16, yini16, xini16 int
	dirs16                         = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func line16(line string) {
	row := make([]rune, 0, len(line))
	for _, ch := range line {
		row = append(row, ch)
	}
	input16 = append(input16, row)
	rows16 = len(input16)
	cols16 = len(input16[0])
	if idx := strings.Index(line, "S"); idx != -1 {
		yini16, xini16 = rows16-1, idx
	}
}

func solve16(b bool) int {
	yend, xend := -1, -1
	dir := 0
	rets := []int{math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}
	pos := dir*cols16*rows16 + yini16*cols16 + xini16
	list := append([]int(nil), pos)
	scores := make(map[int]int)
	scores[pos] = 0
	prev := make(map[int][]int)
	for len(list) > 0 {
		pos = list[0]
		list = list[1:]
		dir, y, x := pos/(cols16*rows16), pos/cols16%rows16, pos%cols16
		score, ok := scores[pos]
		if !ok {
			panic(fmt.Sprintf("not found: %d %d %d %d", pos, dir, y, x))
		}
		if ynext, xnext := y+dirs16[dir][0], x+dirs16[dir][1]; input16[ynext][xnext] != '#' {
			if input16[ynext][xnext] == 'E' {
				yend, xend = ynext, xnext
				rets[dir] = min(rets[dir], score+1)
			}
			posnext := dir*cols16*rows16 + ynext*cols16 + xnext
			nextscore := score + 1
			curscore := scores[posnext]
			if curscore == 0 || nextscore < curscore {
				prev[posnext] = []int{pos}
				scores[posnext] = nextscore
				list = append(list, posnext)
			} else if nextscore == curscore {
				prev[posnext] = append(prev[posnext], pos)
			}
		}
		nextdir := (dir + 1) % 4
		posnext := nextdir*cols16*rows16 + y*cols16 + x
		nextscore := score + 1000
		curscore := scores[posnext]
		if curscore == 0 || nextscore < curscore {
			prev[posnext] = []int{pos}
			scores[posnext] = nextscore
			list = append(list, posnext)
		} else if nextscore == curscore {
			prev[posnext] = append(prev[posnext], pos)
		}
		nextdir = (dir - 1 + 4) % 4
		posnext = nextdir*cols16*rows16 + y*cols16 + x
		curscore = scores[posnext]
		if curscore == 0 || nextscore < curscore {
			prev[posnext] = []int{pos}
			scores[posnext] = nextscore
			list = append(list, posnext)
		} else if nextscore == curscore {
			prev[posnext] = append(prev[posnext], pos)
		}
	}
	ret := min(rets[0], min(rets[1], min(rets[2], rets[3])))
	if !b {
		return ret
	}
	path3 := make(map[int]interface{})
	list = make([]int, 0)
	for dir := range 4 {
		if rets[dir] == ret {
			list = append(list, dir*cols16*rows16+yend*cols16+xend)
		}
	}
	for len(list) > 0 {
		pos = list[0]
		list = list[1:]
		if _, found := path3[pos]; found {
			continue
		}
		path3[pos] = nil
		for _, prevpos := range prev[pos] {
			list = append(list, prevpos)
		}
	}
	path2 := make(map[int]interface{})
	for pos := range path3 {
		_, y, x := pos/(cols16*rows16), pos/cols16%rows16, pos%cols16
		path2[y*cols16+x] = nil
	}
	for pos := range path2 {
		y, x := pos/cols16, pos%cols16
		input16[y][x] = 'O'
	}
	return len(path2)
}

func drawMaze() {
	for j := range input16 {
		for i := range input16[0] {
			fmt.Print(string(input16[j][i]))
		}
		fmt.Println()
	}
	fmt.Println()
}
