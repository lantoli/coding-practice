package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filepath07 = "inputs/07_input.txt"

func main() {
	fmt.Println(calc07(false)) // 248812215
	fmt.Println(calc07(true))  // 250057090
}

type round struct {
	hand [5]byte
	bid  int
	t    int
}

func calc07(joker bool) int {
	f, _ := os.Open(filepath07)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	r := make([]round, 0)
	for scanner.Scan() {
		r = append(r, getRound(scanner.Text(), joker))
	}
	sort.Slice(r, func(i, j int) bool {
		if r[i].t != r[j].t {
			return r[i].t < r[j].t
		}
		for x := range r[i].hand {
			if r[i].hand[x] != r[j].hand[x] {
				return r[i].hand[x] < r[j].hand[x]
			}
		}
		return false
	})
	total := 0
	for i := range r {
		total += r[i].bid * (i + 1)
	}
	return total
}

func getRound(line string, joker bool) round {
	cards := "23456789TJQKA"
	if joker {
		cards = "J23456789TQKA"
	}
	parts := strings.Split(line, " ")
	bid, _ := strconv.Atoi(parts[1])
	var hand [5]byte
	group := make(map[byte]int)
	freqs := make([]int, 0)
	var j int
	for i := range hand {
		hand[i] = byte(strings.Index(cards, string(parts[0][i])))
		if joker && hand[i] == 0 {
			j++
		} else {
			group[hand[i]]++
		}
	}
	for _, v := range group {
		freqs = append(freqs, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freqs)))
	var t int
	if j == 5 || freqs[0]+j == 5 {
		t = 6
	} else if freqs[0]+j == 4 {
		t = 5
	} else if freqs[0]+j == 3 && freqs[1] == 2 {
		t = 4
	} else if freqs[0]+j == 3 {
		t = 3
	} else if freqs[0]+j == 2 && freqs[1] == 2 {
		t = 2
	} else if freqs[0]+j == 2 {
		t = 1
	}
	return round{hand, bid, t}
}
