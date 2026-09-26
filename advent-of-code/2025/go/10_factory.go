package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func main() {
	sa, sb := solve10("inputs/10_sample.txt")
	fmt.Println("Sample A", sa) // 7
	fmt.Println("Sample B", sb) // 33
	ia, ib := solve10("inputs/10_input.txt")
	fmt.Println("A", ia) // 524
	fmt.Println("B", ib) // 21696, solved with AI
}

func solve10(filename string) (int, int) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	totalA, totalB := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, " ")

		// Parse lights
		lights := make([]bool, len(parts[0])-2)
		for i := range lights {
			lights[i] = parts[0][i+1] != '#'
		}

		// Parse joltages
		var joltages []int
		joltStr := parts[len(parts)-1]
		for _, s := range strings.Split(joltStr[1:len(joltStr)-1], ",") {
			n, _ := strconv.Atoi(s)
			joltages = append(joltages, n)
		}

		// Parse buttons
		var buttons [][]int
		for i := 1; i < len(parts)-1; i++ {
			var btn []int
			btnStr := parts[i]
			for _, s := range strings.Split(btnStr[1:len(btnStr)-1], ",") {
				n, _ := strconv.Atoi(s)
				btn = append(btn, n)
			}
			buttons = append(buttons, btn)
		}

		totalA += solve10a(lights, buttons, 0)
		totalB += solve10b(joltages, buttons)
	}
	return totalA, totalB
}

func solve10a(lights []bool, buttons [][]int, pos int) int {
	// Check if all lights are on
	allOn := true
	for _, l := range lights {
		if !l {
			allOn = false
			break
		}
	}
	if allOn {
		return 0
	}
	if pos == len(buttons) {
		return math.MaxInt64
	}

	// Try not pressing this button
	r1 := solve10a(lights, buttons, pos+1)

	// Try pressing this button
	for _, b := range buttons[pos] {
		lights[b] = !lights[b]
	}
	r2 := solve10a(lights, buttons, pos+1)
	for _, b := range buttons[pos] {
		lights[b] = !lights[b]
	}
	if r2 != math.MaxInt64 {
		r2++
	}

	if r1 <= r2 {
		return r1
	}
	return r2
}

// solve10b solves the ILP using gonum for matrix ops + search for free variables
func solve10b(joltages []int, buttons [][]int) int {
	m, n := len(joltages), len(buttons)
	if n == 0 {
		for _, j := range joltages {
			if j != 0 {
				return math.MaxInt64
			}
		}
		return 0
	}

	// Build augmented matrix [A|b] for Gaussian elimination
	aug := mat.NewDense(m, n+1, nil)
	for j, btn := range buttons {
		for _, pos := range btn {
			aug.Set(pos, j, 1)
		}
	}
	for i, j := range joltages {
		aug.Set(i, n, float64(j))
	}

	// RREF using gonum operations
	pivotCols := rref(aug, m, n)

	// Check consistency
	for row := len(pivotCols); row < m; row++ {
		if math.Abs(aug.At(row, n)) > 1e-9 {
			return math.MaxInt64
		}
	}

	// No free variables - return sum
	numFree := n - len(pivotCols)
	if numFree == 0 {
		sum := 0
		for i := range pivotCols {
			v := int(math.Round(aug.At(i, n)))
			if v < 0 {
				return math.MaxInt64
			}
			sum += v
		}
		return sum
	}

	// Identify free variables
	isPivot := make([]bool, n)
	for _, c := range pivotCols {
		isPivot[c] = true
	}
	var freeVars []int
	for j := 0; j < n; j++ {
		if !isPivot[j] {
			freeVars = append(freeVars, j)
		}
	}

	maxJ := 0
	for _, j := range joltages {
		if j > maxJ {
			maxJ = j
		}
	}

	// Search for minimum sum solution
	minSum := math.MaxInt64
	var search func(idx int, vals []int, sum int)
	search = func(idx int, vals []int, sum int) {
		if sum >= minSum {
			return
		}
		if idx == numFree {
			total := sum
			for i := range pivotCols {
				v := aug.At(i, n)
				for fi, fc := range freeVars {
					v -= float64(vals[fi]) * aug.At(i, fc)
				}
				r := int(math.Round(v))
				if math.Abs(v-float64(r)) > 1e-9 || r < 0 {
					return
				}
				total += r
			}
			if total < minSum {
				minSum = total
			}
			return
		}

		// Compute upper bound
		upper := maxJ
		for i := range pivotCols {
			c := aug.At(i, freeVars[idx])
			if c > 1e-9 {
				rhs := aug.At(i, n)
				for fi := 0; fi < idx; fi++ {
					rhs -= float64(vals[fi]) * aug.At(i, freeVars[fi])
				}
				for fi := idx + 1; fi < numFree; fi++ {
					if lc := aug.At(i, freeVars[fi]); lc < -1e-9 {
						rhs -= float64(maxJ) * lc
					}
				}
				if rhs < -1e-9 {
					return
				}
				if b := int(math.Floor(rhs/c + 1e-9)); b < upper {
					upper = b
				}
			}
		}

		for v := 0; v <= upper; v++ {
			vals[idx] = v
			search(idx+1, vals, sum+v)
		}
	}
	search(0, make([]int, numFree), 0)
	return minSum
}

// rref performs Gaussian elimination to reduced row echelon form
func rref(aug *mat.Dense, m, n int) []int {
	var pivotCols []int
	pivotRow := 0

	for col := 0; col < n && pivotRow < m; col++ {
		// Find pivot
		found := -1
		for row := pivotRow; row < m; row++ {
			if math.Abs(aug.At(row, col)) > 1e-9 {
				found = row
				break
			}
		}
		if found == -1 {
			continue
		}

		// Swap rows using gonum
		if found != pivotRow {
			for c := 0; c <= n; c++ {
				tmp := aug.At(pivotRow, c)
				aug.Set(pivotRow, c, aug.At(found, c))
				aug.Set(found, c, tmp)
			}
		}
		pivotCols = append(pivotCols, col)

		// Scale pivot row
		piv := aug.At(pivotRow, col)
		for c := col; c <= n; c++ {
			aug.Set(pivotRow, c, aug.At(pivotRow, c)/piv)
		}

		// Eliminate column
		for row := 0; row < m; row++ {
			if row != pivotRow {
				if f := aug.At(row, col); math.Abs(f) > 1e-9 {
					for c := col; c <= n; c++ {
						aug.Set(row, c, aug.At(row, c)-f*aug.At(pivotRow, c))
					}
				}
			}
		}
		pivotRow++
	}
	return pivotCols
}
