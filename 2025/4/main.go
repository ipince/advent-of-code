package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("4/input.txt")
	if err != nil {
		panic(err)
	}
	output, err := solve(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

func solve(input string) (int, error) {
	lines := strings.Split(input, "\n")
	h := len(lines)
	w := len(lines[0]) // assume at least 1

	grid := make([][]bool, h)
	for i, l := range strings.Split(input, "\n") {
		grid[i] = make([]bool, w)
		for j, r := range l {
			grid[i][j] = string(r) == "@"
		}
	}

	fmt.Printf("grid: %v\n", grid)

	return removeAll(grid), nil
}

func removeAll(grid [][]bool) int {
	total := 0
	for {
		ct := removePass(grid)
		if ct == 0 {
			break
		}
		fmt.Printf("removed %d rolls... continuing\n", ct)
		total += ct
	}
	return total
}

func removePass(grid [][]bool) int {
	ct := 0

	// make one pass
	is := []int{}
	js := []int{}
	for i, row := range grid {
		for j, cell := range row {
			if !cell {
				continue // no roll here
			}
			if accessible(grid, i, j) {
				is = append(is, i) // remove later
				js = append(js, j) // remove later
				ct++
			}
		}
	}
	for k := range is {
		grid[is[k]][js[k]] = false
	}

	return ct
}

func count1(grid [][]bool) int {
	ct := 0
	// grid is 138x138 => 138*138*8 ~= 150k ops
	for i, row := range grid {
		for j, cell := range row {
			if !cell {
				continue // no roll here
			}
			if accessible(grid, i, j) {
				ct++
			}
		}
	}
	return ct
}

func accessible(grid [][]bool, i int, j int) bool {
	// count neighbors
	n := 0

	if i-1 >= 0 && j-1 >= 0 && grid[i-1][j-1] { // top left
		n++
	}
	if i-1 >= 0 && grid[i-1][j] { // top mid
		n++
	}
	if i-1 >= 0 && j+1 < len(grid[i]) && grid[i-1][j+1] { // top right
		n++
	}

	if j-1 >= 0 && grid[i][j-1] { // left
		n++
	}
	if j+1 < len(grid[i]) && grid[i][j+1] { // right
		n++
	}

	if i+1 < len(grid) && j-1 >= 0 && grid[i+1][j-1] { // bottom left
		n++
	}
	if i+1 < len(grid) && grid[i+1][j] { // bottom mid
		n++
	}
	if i+1 < len(grid) && j+1 < len(grid[i]) && grid[i+1][j+1] { // bottom right
		n++
	}

	return n < 4
}
