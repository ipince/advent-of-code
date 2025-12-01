package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("1/input.txt")
	if err != nil {
		panic(err)
	}
	output, err := solve1(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
	output, err = solve2(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

func solve1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	start := 50
	zeros := 0
	for _, line := range lines {
		count, _ := strconv.Atoi(line[1:])
		// fmt.Printf("start/current: %d; line: %s; moving %d steps; zeros: %d\n", start, line, count, zeros)
		if strings.HasPrefix(line, "L") {
			start -= count
		} else { // assume R
			start += count
		}
		start = start % 100
		if start == 0 {
			zeros++
		}
	}
	return zeros, nil
}

func solve2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	start := 50
	zeros := 0
	for _, line := range lines {
		count, _ := strconv.Atoi(line[1:])
		// fmt.Printf("start/current: %d; line: %s; moving %d steps; zeros: %d\n", start, line, count, zeros)
		if strings.HasPrefix(line, "L") {
			if start > 0 && (start-count) <= 0 { // crossed
				zeros++
			}
			start -= count
		} else { // assume R
			start += count
		}
		zeros += int(math.Abs(float64(start / 100)))
		start = start % 100
		if start < 0 {
			start += 100 // ensure positive
		}
	}
	return zeros, nil
}
