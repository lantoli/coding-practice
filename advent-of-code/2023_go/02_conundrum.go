package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const filepath02 = "inputs/02_input.txt"

func main() {
	fmt.Println(calc02(filepath02)) // 1734 70387
}

func calc02(filepath string) (total1, total2 int) {
	pattern := regexp.MustCompile(`^Game (\d+):|([^;]+)`)
	f, _ := os.Open(filepath)
	defer f.Close()
	for scanner := bufio.NewScanner(f); scanner.Scan(); {
		line := scanner.Text()
		matches := pattern.FindAllStringSubmatch(line, -1)
		id, _ := strconv.Atoi(matches[0][1])
		good := true
		var green, red, blue int
		for _, values := range matches[1:] {
			for _, subset := range strings.Split(values[0], ",") {
				num := new(int)
				color := new(string)
				fmt.Sscanf(strings.TrimSpace(subset), "%d %s", num, color)
				if *color == "red" {
					red = max(red, *num)
					good = good && *num <= 12
				}
				if *color == "green" {
					green = max(green, *num)
					good = good && *num <= 13
				}
				if *color == "blue" {
					blue = max(blue, *num)
					good = good && *num <= 14
				}
			}
		}
		if good {
			total1 += id
		}
		total2 += green * red * blue
	}
	return
}
