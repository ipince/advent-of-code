package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("7/input.txt")
	if err != nil {
		panic(err)
	}
	output, err := solve2(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

func solve1(input string) (int, error) {
	lines := strings.Split(input, "\n")

	cols := map[int]bool{}
	// splitters := [][]int{} // for each row, the splitter's position
	splits := 0
	for _, line := range lines {
		for j, cell := range line {
			if string(cell) == "S" {
				cols[j] = true
			}

			if string(cell) == "^" && cols[j] {
				cols[j] = false
				cols[j-1] = true
				cols[j+1] = true
				splits++
			}
		}
	}

	return splits, nil
}

func solve2(input string) (int, error) {
	lines := strings.Split(input, "\n")

	numPaths := map[int]int{} // number of paths to end up at each column position
	for i, line := range lines {
		fmt.Printf("processing line %d out of %d\n", i, len(lines))
		for j, cell := range line {
			if string(cell) == "S" {
				numPaths[j] = 1
			}

			if string(cell) == "^" {
				numPaths[j-1] += numPaths[j]
				numPaths[j+1] += numPaths[j]
				numPaths[j] = 0
			}
		}
	}

	total := 0
	for _, num := range numPaths {
		total += num
	}

	return total, nil
}
