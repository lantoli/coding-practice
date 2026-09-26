package main

// 754 1609 (sample 36 81)
func main() {
	solve("10_input.txt", line10a, nil, solve10a, solve10b)
}

var (
	input10        = make([][]int, 0)
	rows10, cols10 int
)

func line10a(line string) {
	row := make([]int, len(line))
	for i, ch := range line {
		row[i] = int(ch) - '0'
	}
	input10 = append(input10, row)
	rows10 = len(input10)
	cols10 = len(input10[0])
}

func solve10a() int {
	total := 0
	for y := range rows10 {
		for x := range cols10 {
			if input10[y][x] == 0 {
				visited := make(map[int]interface{}, 0)
				paths10a(y, x, visited)
				total += len(visited)
			}
		}
	}
	return total
}

func paths10a(y, x int, visited map[int]interface{}) {
	if input10[y][x] == 9 {
		visited[y*cols10+x] = true
		return
	}
	if y-1 >= 0 && input10[y-1][x] == input10[y][x]+1 {
		paths10a(y-1, x, visited)
	}
	if y+1 < rows10 && input10[y+1][x] == input10[y][x]+1 {
		paths10a(y+1, x, visited)
	}
	if x-1 >= 0 && input10[y][x-1] == input10[y][x]+1 {
		paths10a(y, x-1, visited)
	}
	if x+1 < cols10 && input10[y][x+1] == input10[y][x]+1 {
		paths10a(y, x+1, visited)
	}
}

func solve10b() int {
	total := 0
	for y := range rows10 {
		for x := range cols10 {
			if input10[y][x] == 0 {
				total += paths10b(y, x, 0)
			}
		}
	}
	return total
}

func paths10b(y, x, acc int) int {
	if input10[y][x] == 9 {
		return acc + 1
	}
	total := 0
	if y-1 >= 0 && input10[y-1][x] == input10[y][x]+1 {
		total += paths10b(y-1, x, acc)
	}
	if y+1 < rows10 && input10[y+1][x] == input10[y][x]+1 {
		total += paths10b(y+1, x, acc)
	}
	if x-1 >= 0 && input10[y][x-1] == input10[y][x]+1 {
		total += paths10b(y, x-1, acc)
	}
	if x+1 < cols10 && input10[y][x+1] == input10[y][x]+1 {
		total += paths10b(y, x+1, acc)
	}
	return total
}
