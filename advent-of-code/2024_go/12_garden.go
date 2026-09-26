package main

// 1550156 946084 (sample 1930 1206)
func main() {
	solve("12_input.txt", line12, nil, func() int { return solve12(false) }, func() int { return solve12(true) })
}

var (
	input12        = make([][]rune, 0)
	rows12, cols12 int
)

func line12(line string) {
	row := make([]rune, 0, len(line))
	for _, ch := range line {
		row = append(row, ch)
	}
	input12 = append(input12, row)
	rows12 = len(input12)
	cols12 = len(input12[0])
}

func solve12(sides bool) int {
	visited := make(map[int]bool)
	total := 0
	for y := range rows12 {
		for x := range cols12 {
			total += calc12(y, x, visited, sides)
		}
	}
	return total
}

func calc12(yini, xini int, visited map[int]bool, sides bool) int {
	a := 0
	ch := input12[yini][xini]
	list := []int{yini*cols12 + xini}
	borders := make(map[int]interface{})
	for len(list) > 0 {
		pos := list[len(list)-1]
		y, x := pos/cols12, pos%cols12
		list = list[:len(list)-1]
		if input12[y][x] != ch || visited[pos] {
			continue
		}
		visited[pos] = true
		a++
		if y == 0 || input12[y-1][x] != ch {
			borders[0*cols12*rows12+y*cols12+x] = nil
		}
		if y == rows12-1 || input12[y+1][x] != ch {
			borders[1*cols12*rows12+y*cols12+x] = nil
		}
		if x == 0 || input12[y][x-1] != ch {
			borders[2*cols12*rows12+y*cols12+x] = nil
		}
		if x == cols12-1 || input12[y][x+1] != ch {
			borders[3*cols12*rows12+y*cols12+x] = nil
		}
		if y > 0 {
			list = append(list, pos-cols12)
		}
		if y < rows12-1 {
			list = append(list, pos+cols12)
		}
		if x > 0 {
			list = append(list, pos-1)
		}
		if x < cols12-1 {
			list = append(list, pos+1)
		}
	}
	if sides {
		for yini := range rows12 {
			for xini := range cols12 {
				if _, found := borders[0*cols12*rows12+yini*cols12+xini]; found {
					for x := xini + 1; x < cols12; x++ {
						if _, found := borders[0*cols12*rows12+yini*cols12+x]; found {
							delete(borders, 0*cols12*rows12+yini*cols12+x)
						} else {
							break
						}
					}
				}
				if _, found := borders[1*cols12*rows12+yini*cols12+xini]; found {
					for x := xini + 1; x < cols12; x++ {
						if _, found := borders[1*cols12*rows12+yini*cols12+x]; found {
							delete(borders, 1*cols12*rows12+yini*cols12+x)
						} else {
							break
						}
					}
				}
				if _, found := borders[2*cols12*rows12+yini*cols12+xini]; found {
					for y := yini + 1; y < rows12; y++ {
						if _, found := borders[2*cols12*rows12+y*cols12+xini]; found {
							delete(borders, 2*cols12*rows12+y*cols12+xini)
						} else {
							break
						}
					}
				}
				if _, found := borders[3*cols12*rows12+yini*cols12+xini]; found {
					for y := yini + 1; y < rows12; y++ {
						if _, found := borders[3*cols12*rows12+y*cols12+xini]; found {
							delete(borders, 3*cols12*rows12+y*cols12+xini)
						} else {
							break
						}
					}
				}
			}
		}
	}
	return a * len(borders)
}
