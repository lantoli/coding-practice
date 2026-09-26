package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const filepath18 = "inputs/18_input.txt"

var grid18 []big.Int

func main() {
	fmt.Println(calc18(false)) // 48652
	fmt.Println(calc18(true))  // XXX
}

func calc18(swap bool) int {
	f, _ := os.Open(filepath18)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	insts := make([]inst18, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if swap {
			idx := strings.Index(line, "(#") + 2
			count, _ := strconv.ParseInt(line[idx:idx+5], 16, 64)
			var ch byte
			switch line[idx+5] {
			case '0':
				ch = 'R'
			case '1':
				ch = 'D'
			case '2':
				ch = 'L'
			case 3:
				ch = 'U'
			}
			insts = append(insts, inst18{ch, int(count)})
		} else {
			parts := strings.Fields(line)
			count, _ := strconv.Atoi(parts[1])
			insts = append(insts, inst18{parts[0][0], count})
		}
	}
	return calc18Main(insts)
}

func calc18Main(insts []inst18) int {
	var ymin, xmin, ymax, xmax, y, x int
	for _, inst := range insts {
		yinc, xinc := inst.inc()
		y, x = y+yinc*inst.count, x+xinc*inst.count
		xmax, xmin, ymin, ymax = max(xmax, x), min(xmin, x), min(ymin, y), max(ymax, y)
	}
	yrange, xrange := ymax-ymin+1, xmax-xmin+1
	grid18 = make([]big.Int, yrange)
	//for i := range grid18 {
	//	grid18[i] = make([]bool, xrange)
	//	}
	y, x = -ymin, -xmin
	grid18[y].SetBit(&grid18[y], x, 1)
	for _, inst := range insts {
		yinc, xinc := inst.inc()
		for i := 0; i < inst.count; i++ {
			y, x = y+yinc, x+xinc
			grid18[y].SetBit(&grid18[y], x, 1)
		}
	}
	out := 0
	for y := 0; y < yrange; y++ {
		out += color(y, 0, yrange, xrange)
		out += color(y, xrange-1, yrange, xrange)
	}
	for x := 0; x < xrange; x++ {
		out += color(0, x, yrange, xrange)
		out += color(yrange-1, x, yrange, xrange)
	}
	return xrange*yrange - out
}

func color(y, x, yrange, xrange int) int {
	ret := 0
	visited := append([]int{}, y*xrange+x)
	count := 0
	for len(visited) > 0 {
		newVisited := make([]int, 0, len(visited)*4)
		for _, v := range visited {
			y, x := v/xrange, v%xrange
			if y >= 0 && y < yrange && x >= 0 && x < xrange && grid18[y].Bit(x) == 0 {
				grid18[y].SetBit(&grid18[y], x, 1)
				ret++
				newVisited = append(newVisited, y*xrange+x-1, y*xrange+x+1, (y-1)*xrange+x, (y+1)*xrange+x)
				if len(newVisited) > 10 {
					newVisited = newVisited[:10] // DELETE THIS
				}
			}
			count++
			if (count % 10) == 0 {
				runtime.GC()
			}
		}
		visited = newVisited
	}
	return ret
}

type inst18 struct {
	dir   byte
	count int
}

func (i *inst18) inc() (yinc, xinc int) {
	switch i.dir {
	case 'R':
		xinc = 1
	case 'L':
		xinc = -1
	case 'U':
		yinc = -1
	case 'D':
		yinc = 1
	}
	return
}
