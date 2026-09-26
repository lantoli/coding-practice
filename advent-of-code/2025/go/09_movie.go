package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type coord struct{ x, y int }
type seg struct{ fixed, lo, hi int }

func main() {
	fmt.Println("Sample A", solve09("inputs/09_sample.txt", false)) // 50
	fmt.Println("A", solve09("inputs/09_input.txt", false))         // 4776100539
	fmt.Println("Sample B", solve09("inputs/09_sample.txt", true))  // 24
	fmt.Println("B", solve09("inputs/09_input.txt", true))          // 1476550548
}

func solve09(filename string, checkInside bool) int {
	coords := parseCoords(filename)
	hSegs, vSegs := buildSegments(coords)

	type rect struct{ x1, y1, x2, y2, area int }
	var rects []rect
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			x1, x2 := minmax(c1.x, c2.x)
			y1, y2 := minmax(c1.y, c2.y)
			rects = append(rects, rect{x1, y1, x2, y2, (x2 - x1 + 1) * (y2 - y1 + 1)})
		}
	}
	sort.Slice(rects, func(i, j int) bool { return rects[i].area > rects[j].area })

	for _, r := range rects {
		if !checkInside || isRectInside(r.x1, r.y1, r.x2, r.y2, hSegs, vSegs) {
			return r.area
		}
	}
	panic("no solution found")
}

func parseCoords(filename string) []coord {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var coords []coord
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			var c coord
			fmt.Sscanf(line, "%d,%d", &c.x, &c.y)
			coords = append(coords, c)
		}
	}
	return coords
}

func buildSegments(coords []coord) (hSegs, vSegs []seg) {
	for i, c1 := range coords {
		c2 := coords[(i+1)%len(coords)]
		if c1.y == c2.y {
			lo, hi := minmax(c1.x, c2.x)
			hSegs = append(hSegs, seg{c1.y, lo, hi})
		} else {
			lo, hi := minmax(c1.y, c2.y)
			vSegs = append(vSegs, seg{c1.x, lo, hi})
		}
	}
	return
}

func isRectInside(x1, y1, x2, y2 int, hSegs, vSegs []seg) bool {
	for _, h := range hSegs {
		if y1 < h.fixed && h.fixed < y2 && h.lo < x2 && x1 < h.hi {
			return false
		}
	}
	for _, v := range vSegs {
		if x1 < v.fixed && v.fixed < x2 && v.lo < y2 && y1 < v.hi {
			return false
		}
	}
	return true
}

func minmax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
