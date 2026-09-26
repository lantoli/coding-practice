package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filepath19 = "inputs/19_input.txt"

func main() {
	fmt.Println(calc19(false)) // 432434
	fmt.Println(calc19(true))  // 132557544578569
}

func calc19(part2 bool) int {
	f, _ := os.Open(filepath19)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	workflows := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		index := strings.Index(line, "{")
		workflows[line[:index]] = line[index+1 : len(line)-1]
	}
	if part2 {
		piece := map[string]int{"xmin": 1, "xmax": 4000, "mmin": 1, "mmax": 4000, "amin": 1, "amax": 4000, "smin": 1, "smax": 4000}
		return calc19Part2(piece, workflows, "in", 0)
	}
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		piece := make(map[string]int)
		for _, assig := range strings.Split(line[1:len(line)-1], ",") {
			parts := strings.Split(assig, "=")
			num, _ := strconv.Atoi(parts[1])
			piece[parts[0]] = num
		}
		total += calc19Part1(piece, workflows, "in")
	}
	return total
}

func calc19Part1(piece map[string]int, ws map[string]string, in string) int {
	for _, rule := range strings.Split(ws[in], ",") {
		parts := strings.Split(rule, ":")
		next := parts[0]
		if len(parts) == 2 {
			letter := parts[0][0:1]
			cond := parts[0][1]
			num, _ := strconv.Atoi(parts[0][2:])
			if cond == '<' && piece[letter] >= num {
				continue
			} else if cond == '>' && piece[letter] <= num {
				continue
			}
			next = parts[1]
		}
		if next == "A" {
			ret := 0
			for _, v := range piece {
				ret += v
			}
			return ret
		} else if next == "R" {
			return 0
		} else {
			return calc19Part1(piece, ws, next)
		}
	}
	return 0
}

func calc19Part2(piece map[string]int, ws map[string]string, in string, pos int) int {
	var ret int
	rule := strings.Split(ws[in], ",")[pos]
	parts := strings.Split(rule, ":")
	next := parts[0]
	if len(parts) == 2 {
		next = parts[1]
		letter := parts[0][0:1]
		cond := parts[0][1]
		num, _ := strconv.Atoi(parts[0][2:])

		if cond == '<' {
			if piece[letter+"min"] >= num {
				return calc19Part2(piece, ws, in, pos+1)
			}
			if piece[letter+"max"] >= num {
				newPiece := duplicatePiece(piece)
				newPiece[letter+"min"] = num
				ret = calc19Part2(newPiece, ws, in, pos+1)
				piece = duplicatePiece(piece)
				piece[letter+"max"] = num - 1
			}
		}

		if cond == '>' {
			if piece[letter+"max"] <= num {
				return calc19Part2(piece, ws, in, pos+1)
			}
			if piece[letter+"min"] <= num {
				newPiece := duplicatePiece(piece)
				newPiece[letter+"max"] = num
				ret = calc19Part2(newPiece, ws, in, pos+1)
				piece = duplicatePiece(piece)
				piece[letter+"min"] = num + 1
			}
		}
	}
	if next == "A" {
		xmin, xmax := piece["xmin"], piece["xmax"]
		mmin, mmax := piece["mmin"], piece["mmax"]
		amin, amax := piece["amin"], piece["amax"]
		smin, smax := piece["smin"], piece["smax"]
		return ret + (xmax-xmin+1)*(mmax-mmin+1)*(amax-amin+1)*(smax-smin+1)
	}
	if next == "R" {
		return ret
	}
	return ret + calc19Part2(piece, ws, next, 0)
}

func duplicatePiece(piece map[string]int) map[string]int {
	newPiece := make(map[string]int)
	for k, v := range piece {
		newPiece[k] = v
	}
	return newPiece
}
