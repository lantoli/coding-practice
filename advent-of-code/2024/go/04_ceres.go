package main

// 2591 1880 (sample 18 9)
func main() {
	solve("04_input.txt", line04a, line04b, solve04a, solve04b)
}

var (
	grid04 = make([]string, 0)
)

func line04a(line string) {
	grid04 = append(grid04, line)
}

func solve04a() int {
	word := "XMAS"
	total := 0
	for y := range len(grid04) {
		for x := range len(grid04[y]) {
			for yinc := -1; yinc <= 1; yinc++ {
				for xinc := -1; xinc <= 1; xinc++ {
					for pos, yy, xx := 0, y, x; isGood04(yy, xx) && grid04[yy][xx] == word[pos]; pos, yy, xx = pos+1, yy+yinc, xx+xinc {
						if pos == len(word)-1 {
							total++
							break
						}
					}
				}
			}
		}
	}
	return total
}

func isGood04(y, x int) bool {
	return y >= 0 && y < len(grid04) && x >= 0 && x < len(grid04[y])
}

func line04b(line string) {
}

func solve04b() int {
	total := 0
	for y := 1; y < len(grid04)-1; y++ {
		for x := 1; x < len(grid04[y])-1; x++ {
			if grid04[y][x] == 'A' &&
				(grid04[y-1][x-1] == 'M' && grid04[y+1][x+1] == 'S' || grid04[y-1][x-1] == 'S' && grid04[y+1][x+1] == 'M') &&
				(grid04[y+1][x-1] == 'M' && grid04[y-1][x+1] == 'S' || grid04[y+1][x-1] == 'S' && grid04[y-1][x+1] == 'M') {
				total++
			}
		}
	}
	return total
}
