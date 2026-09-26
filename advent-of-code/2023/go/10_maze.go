package main

import (
	"bufio"
	"fmt"
	"os"
)

const filepath10 = "inputs/10_input.txt"

func main() {
	fmt.Println(calc10(false)) // 7093
	fmt.Println(calc10(true))  // 407 (ret: 3844, 3844, 407, 407)
}

func calc10(enclosed bool) []int {
	f, _ := os.Open(filepath10)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	maze := make([][]byte, 0)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		maze = append(maze, line)
	}
	rows := len(maze)
	cols := len(maze[0])
	var source state
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if maze[row][col] == 'S' {
				source = state{row: row, col: col}
				break
			}
		}
	}
	total := 0
	states := append([]state{}, source)
	var sinfo info
	visited := make(map[state]bool)
	visited[source] = true
	if source.row > 0 && infos[maze[source.row-1][source.col]].s {
		states = append(states, state{row: source.row - 1, col: source.col})
		sinfo.n = true
	}
	if infos[maze[source.row+1][source.col]].n {
		states = append(states, state{row: source.row + 1, col: source.col})
		sinfo.s = true
	}
	if infos[maze[source.row][source.col+1]].w {
		states = append(states, state{row: source.row, col: source.col + 1})
		sinfo.e = true
	}
	if source.col > 0 && infos[maze[source.row][source.col-1]].e {
		states = append(states, state{row: source.row, col: source.col - 1})
		sinfo.w = true
	}

	for k, v := range infos {
		if v == sinfo {
			maze[source.row][source.col] = k
		}
	}
	for {
		newStates := make([]state, 0)
		for _, s := range states {
			if !visited[s] {
				visited[s] = true
				i := infos[maze[s.row][s.col]]
				if i.n {
					newStates = append(newStates, state{row: s.row - 1, col: s.col})
				}
				if i.s {
					newStates = append(newStates, state{row: s.row + 1, col: s.col})
				}
				if i.e {
					newStates = append(newStates, state{row: s.row, col: s.col + 1})
				}
				if i.w {
					newStates = append(newStates, state{row: s.row, col: s.col - 1})
				}
			}
		}
		if len(newStates) == 0 {
			break
		}
		states = newStates
		total++
	}
	if enclosed {
		return []int{
			countCluster(visited, maze, state{source.row, source.col, false, true, true, false}),
			countCluster(visited, maze, state{source.row, source.col, true, false, true, false}),
			countCluster(visited, maze, state{source.row, source.col, false, true, false, true}),
			countCluster(visited, maze, state{source.row, source.col, true, false, false, true}),
		}
	} else {
		return []int{total}
	}
}

func countCluster(walls map[state]bool, maze [][]byte, ini state) int {
	rows := len(maze)
	cols := len(maze[0])
	visited := make(map[state]bool)
	count := make(map[state]bool)
	states := append([]state{}, ini)
	for len(states) > 0 {
		s := states[len(states)-1]
		states = states[:len(states)-1]
		if visited[s] {
			continue
		}
		visited[s] = true
		basicState := state{row: s.row, col: s.col}
		if !walls[basicState] {
			count[basicState] = true
		}

		if s.s && s.row+1 < rows {
			i := infos[maze[s.row+1][s.col]]
			states = append(states, state{s.row + 1, s.col, !s.n, !s.s, s.e, s.w})
			if (s.w && !i.w) || (s.e && !i.e) || !walls[state{row: s.row + 1, col: s.col}] {
				states = append(states, state{s.row + 1, s.col, s.n, s.s, s.e, s.w})
			}
		}

		if s.n && s.row-1 >= 0 {
			i := infos[maze[s.row-1][s.col]]
			states = append(states, state{s.row - 1, s.col, !s.n, !s.s, s.e, s.w})
			if (s.w && !i.w) || (s.e && !i.e) || !walls[state{row: s.row - 1, col: s.col}] {
				states = append(states, state{s.row - 1, s.col, s.n, s.s, s.e, s.w})
			}
		}

		if s.e && s.col+1 < cols {
			i := infos[maze[s.row][s.col+1]]
			states = append(states, state{s.row, s.col + 1, s.n, s.s, !s.e, !s.w})
			if (s.n && !i.n) || (s.s && !i.s) || !walls[state{row: s.row, col: s.col + 1}] {
				states = append(states, state{s.row, s.col + 1, s.n, s.s, s.e, s.w})
			}
		}

		if s.w && s.col-1 >= 0 {
			i := infos[maze[s.row][s.col-1]]
			states = append(states, state{s.row, s.col - 1, s.n, s.s, !s.e, !s.w})
			if (s.n && !i.n) || (s.s && !i.s) || !walls[state{row: s.row, col: s.col - 1}] {
				states = append(states, state{s.row, s.col - 1, s.n, s.s, s.e, s.w})
			}
		}
	}
	return len(count)
}

var infos = map[byte]info{
	'|': {true, true, false, false},
	'-': {false, false, true, true},
	'L': {true, false, true, false},
	'J': {true, false, false, true},
	'7': {false, true, false, true},
	'F': {false, true, true, false},
}

type info struct {
	n, s, e, w bool
}

type state struct {
	row, col   int
	n, s, e, w bool
}
