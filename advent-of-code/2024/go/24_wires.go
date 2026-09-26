package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 52728619468518 XX  (sample 2024 XX)
func main() {
	solve24a()
}

var (
	file24 = "24_input.txt"
)

func solve24a() {
	f, _ := os.Open("inputs/" + file24)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	vals := make(map[string]bool)
	remaining := make(map[string]struct{})
	gates := make([]st24, 0)
	readingVals := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingVals = false
			continue
		}
		if readingVals {
			const sep = ": "
			idx := strings.Index(line, sep)
			vals[line[:idx]] = line[idx+len(sep):] == "1"
		} else {
			parts := strings.Split(line, " -> ")
			part1 := strings.Split(parts[0], " ")
			gate := st24{part1[1], part1[0], part1[2], parts[1]}
			gates = append(gates, gate)
			remaining[gate.in1] = struct{}{}
			remaining[gate.in2] = struct{}{}
			remaining[gate.out] = struct{}{}
		}
	}
	for gate := range vals {
		delete(remaining, gate)
	}
	for (len(remaining)) > 0 {
		for _, gate := range gates {
			_, rin1 := remaining[gate.in1]
			_, rin2 := remaining[gate.in2]
			_, rout := remaining[gate.out]
			if rout && !rin1 && !rin2 {
				switch gate.op {
				case "AND":
					vals[gate.out] = vals[gate.in1] && vals[gate.in2]
				case "OR":
					vals[gate.out] = vals[gate.in1] || vals[gate.in2]
				case "XOR":
					vals[gate.out] = vals[gate.in1] != vals[gate.in2]
				default:
					panic("unknown op")
				}
				delete(remaining, gate.out)
			}
		}
	}
	z := make([]string, 0)
	for k := range vals {
		if k[0] == 'z' {
			z = append(z, k)
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(z)))
	num := ""
	for _, k := range z {
		if vals[k] {
			num += "1"
		} else {
			num += "0"
		}
	}
	ret, _ := strconv.ParseInt(num, 2, 64)
	fmt.Println(ret)
}

type st24 struct {
	op       string
	in1, in2 string
	out      string
}
