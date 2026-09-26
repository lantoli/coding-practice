package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const filepath15 = "inputs/15_input.txt"

func main() {
	fmt.Println(calc15(false)) // 513214
	fmt.Println(calc15(true))  // 258826
}

func calc15(part2 bool) int {
	f, _ := os.Open(filepath15)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	scanner.Scan()
	seqs := bytes.Split(scanner.Bytes(), []byte{','})
	if part2 {
		total = calc15Part2(seqs)
	} else {
		for _, seq := range seqs {
			total += calc15Hash(seq)
		}
	}
	return total
}

func calc15Hash(seq []byte) int {
	ret := 0
	for _, s := range seq {
		ret = (ret + int(s)) * 17 % 256
	}
	return ret
}

func calc15Part2(seqs [][]byte) int {
	boxes := make([][]stbox, 256)
	for _, seq := range seqs {
		length := 0
		s := seq[:len(seq)-1]
		delete := seq[len(seq)-1] == '-'
		if !delete {
			length = int(seq[len(seq)-1]) - '0'
			s = seq[:len(seq)-2]
		}
		tag := string(s)
		box := calc15Hash(s)
		found := false
		for i := range boxes[box] {
			if boxes[box][i].tag == tag {
				if delete {
					boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
				} else {
					boxes[box][i].length = length
				}
				found = true
				break
			}
		}
		if !found && !delete {
			boxes[box] = append(boxes[box], stbox{tag, length})
		}
	}
	total := 0
	for b := range boxes {
		for l := range boxes[b] {
			total += boxes[b][l].length * (b + 1) * (l + 1)
		}
	}
	return total
}

type stbox struct {
	tag    string
	length int
}
