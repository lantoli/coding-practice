package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type stdist struct {
	pos1, pos2, sqdist int
}

func main() {
	fmt.Println("Sample A", solve08("inputs/08_sample.txt", 10, false)) // 40
	fmt.Println("A", solve08("inputs/08_input.txt", 1000, false))       // 123930
	fmt.Println("Sample B", solve08("inputs/08_sample.txt", 10, true))  // 25272
	fmt.Println("B", solve08("inputs/08_input.txt", 1000, true))        // 27338688
}

func solve08(filename string, connections int, b bool) int {
	var (
		f, _    = os.Open(filename)
		scanner = bufio.NewScanner(f)
		n       = 0
		x       []int
		y       []int
		z       []int
		dist    []stdist
	)
	defer f.Close()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		x, y, z = append(x, 0), append(y, 0), append(z, 0)
		n++
		fmt.Sscanf(line, "%d,%d,%d", &x[len(x)-1], &y[len(y)-1], &z[len(z)-1])
	}
	parent := make([]int, n)
	for i := range n {
		parent[i] = i
		for j := i + 1; j < n; j++ {
			dist = append(dist, stdist{i, j, (x[i]-x[j])*(x[i]-x[j]) + (y[i]-y[j])*(y[i]-y[j]) + (z[i]-z[j])*(z[i]-z[j])})
		}
	}
	if b {
		connections = len(dist) + 1
	}
	sort.Slice(dist, func(i, j int) bool { return dist[i].sqdist < dist[j].sqdist })
	for _, d := range dist[:connections] {
		count1, count2, pos1, pos2 := 0, 0, d.pos1, d.pos2
		for pos1 != parent[pos1] {
			count1++
			pos1 = parent[pos1]
		}
		for pos2 != parent[pos2] {
			count2++
			pos2 = parent[pos2]
		}
		if pos1 != pos2 {
			parent[pos2] = pos1
			if b && allTogether(parent) {
				return x[d.pos1] * x[d.pos2]
			}
		}
	}
	counts := getCounts(parent)
	return counts[len(counts)-1] * counts[len(counts)-2] * counts[len(counts)-3]
}

func getCounts(parent []int) []int {
	n := len(parent)
	counts := make([]int, n)
	for i := range n {
		for i != parent[i] {
			i = parent[i]
		}
		counts[i]++
	}
	sort.Ints(counts)
	return counts
}

func allTogether(parent []int) bool {
	counts := getCounts(parent)
	return counts[len(counts)-1] == len(parent)
}
