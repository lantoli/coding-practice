package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

// 1083 as,bu,cp,dj,ez,fd,hu,it,kj,nx,pp,xh,yu (sample 7 co,de,ka,ta)
func main() {
	solve23a()
	solve23b()
}

var (
	file23 = "23_input.txt"
)

func solve23a() {
	n, list, matrix := get23()
	total := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if i == j || i == k || j == k || !matrix[i*n+j] || !matrix[j*n+k] || !matrix[i*n+k] {
					continue
				}
				if strings.HasPrefix(list[i], "t") || strings.HasPrefix(list[j], "t") || strings.HasPrefix(list[k], "t") {
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

func solve23b() {
	n, computers, matrix := get23()
	states := make([][]int, n)
	var longest []int
	for i := range n {
		states[i] = []int{i}
	}
	for len(states) > 0 {
		state := states[0]
		states = states[1:]
		if len(state) > len(longest) {
			longest = state
		}
		last := state[len(state)-1]
		for j := last + 1; j < n; j++ {
			together := true
			for i := 0; i < len(state); i++ {
				if !matrix[state[i]*n+j] {
					together = false
					break
				}
			}
			if together {
				newState := make([]int, len(state)+1)
				copy(newState, state)
				newState[len(newState)-1] = j
				states = append(states, newState)
			}
		}
	}
	var ret []string
	for _, i := range longest {
		ret = append(ret, computers[i])
	}
	fmt.Println(strings.Join(ret, ","))
}

func get23() (n int, list []string, matrix []bool) {
	f, _ := os.Open("inputs/" + file23)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	m := make(map[string][]string)
	unsorted := make(map[string]struct{})
	for scanner.Scan() {
		line := scanner.Text()
		a, b := line[0:2], line[3:5]
		unsorted[a] = struct{}{}
		unsorted[b] = struct{}{}
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}
	list = make([]string, 0, len(unsorted))
	for k := range unsorted {
		list = append(list, k)
	}
	sort.Strings(list)
	n = len(list)
	matrix = make([]bool, n*n)
	for i := range list {
		for _, val := range m[list[i]] {
			j := slices.Index(list, val)
			matrix[i*n+j] = true
			matrix[j*n+i] = true
		}
	}
	return
}
