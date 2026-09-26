package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filepath20 = "inputs/20_input.txt"

func main() {
	fmt.Println(calc20()) // 763500168
}

func calc20() int {
	f, _ := os.Open(filepath20)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	m := make(map[string]*st20mod)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		name := parts[0]
		var symbol byte
		if name[0] == '%' || name[0] == '&' {
			symbol = name[0]
			name = name[1:]
		}
		m[name] = &st20mod{symbol, false, strings.Split(parts[1], ", "), 0, make(map[string]any)}
	}
	for _, st := range m {
		for _, d := range st.dest {
			if m[d] != nil {
				m[d].srcCount++
			}
		}
	}
	lo, hi := 0, 0
	for i := 0; i < 1000; i++ {
		l, h := calc20Main(m)
		lo += l
		hi += h
	}
	return lo * hi
}

func calc20Main(m map[string]*st20mod) (int, int) {
	var lo, hi int
	ps := []st20pul{{"broadcaster", "", false}}
	for len(ps) > 0 {
		newps := []st20pul{}
		for _, p := range ps {
			if p.pulse {
				hi++
			} else {
				lo++
			}
			st := m[p.mod]
			if st == nil {
				continue
			}
			if p.mod == "broadcaster" {
				newPulse := p.pulse
				for _, d := range st.dest {
					newps = append(newps, st20pul{d, p.mod, newPulse})
				}
			} else if st.symbol == '%' {
				if !p.pulse {
					st.on = !st.on
					newPulse := st.on
					for _, d := range st.dest {
						newps = append(newps, st20pul{d, p.mod, newPulse})
					}
				}
			} else if st.symbol == '&' {
				if p.pulse {
					st.srcHigh[p.src] = nil
				} else {
					delete(st.srcHigh, p.src)
				}
				newPulse := len(st.srcHigh) != st.srcCount
				for _, d := range st.dest {
					newps = append(newps, st20pul{d, p.mod, newPulse})
				}
			}
		}
		ps = newps
	}
	return lo, hi
}

type st20mod struct {
	symbol   byte
	on       bool
	dest     []string
	srcCount int
	srcHigh  map[string]any
}

type st20pul struct {
	mod   string
	src   string
	pulse bool
}
