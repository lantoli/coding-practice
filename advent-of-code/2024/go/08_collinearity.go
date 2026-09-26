package main

// 269 949 (sample 14 34)
func main() {
	solve("08_input.txt", line08a, nil, solve08a, solve08b)
}

var (
	grid08 = make([][]rune, 0)
)

func line08a(line string) {
	grid08 = append(grid08, []rune(line))
}

func solve08a() int {
	rows, cols := len(grid08), len(grid08[0])
	found := make(map[int]struct{}, rows*cols)
	for pos1 := 0; pos1 < rows*cols-1; pos1++ {
		y1, x1 := pos1/cols, pos1%cols
		if grid08[y1][x1] == '.' {
			continue
		}
		for pos2 := pos1 + 1; pos2 < rows*cols; pos2++ {
			y2, x2 := pos2/cols, pos2%cols
			if grid08[y1][x1] != grid08[y2][x2] {
				continue
			}
			ydif, xdif := y2-y1, x2-x1
			if nodey, nodex := y1-ydif, x1-xdif; isValid08(nodey, nodex) {
				found[nodey*cols+nodex] = struct{}{}
			}
			if nodey, nodex := y2+ydif, x2+xdif; isValid08(nodey, nodex) {
				found[nodey*cols+nodex] = struct{}{}
			}
		}
	}
	return len(found)
}

func solve08b() int {
	rows, cols := len(grid08), len(grid08[0])
	found := make(map[int]struct{}, rows*cols)
	for pos1 := 0; pos1 < rows*cols-1; pos1++ {
		y1, x1 := pos1/cols, pos1%cols
		if grid08[y1][x1] == '.' {
			continue
		}
		for pos2 := pos1 + 1; pos2 < rows*cols; pos2++ {
			y2, x2 := pos2/cols, pos2%cols
			if grid08[y1][x1] != grid08[y2][x2] {
				continue
			}
			ydif, xdif := y2-y1, x2-x1
			for nodey, nodex := y1, x1; isValid08(nodey, nodex); nodey, nodex = nodey-ydif, nodex-xdif {
				found[nodey*cols+nodex] = struct{}{}
			}
			for nodey, nodex := y2, x2; isValid08(nodey, nodex); nodey, nodex = nodey+ydif, nodex+xdif {
				found[nodey*cols+nodex] = struct{}{}
			}
		}
	}
	return len(found)
}

func isValid08(y, x int) bool {
	return y >= 0 && y < len(grid08) && x >= 0 && x < len(grid08[0])
}
