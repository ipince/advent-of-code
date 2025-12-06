package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("6/input.txt")
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

	parts := strings.Fields(lines[0])
	operands := make([][]int, len(parts))
	total := 0
	for i, l := range lines {

		fmt.Printf("%v\n", operands)
		fmt.Println(l)

		parts := strings.Fields(l)

		if i == len(lines)-1 {
			// operations
			for j, op := range parts {
				res := 0
				if op == "+" {
					res = fold(operands[j], func(a, acc int) int { return a + acc }, 0)
				} else {
					res = fold(operands[j], func(a, acc int) int { return a * acc }, 1)
				}
				fmt.Printf("applied %s to %v = %d\n", op, operands[j], res)
				total += res
			}
			break
		}

		for j, p := range parts {
			if operands[j] == nil {
				operands[j] = make([]int, len(lines)-1)
			}
			val, _ := strconv.Atoi(p)
			operands[j][i] = val
		}
	}

	fmt.Printf("%v\n", operands)

	return total, nil
}

func fold(s []int, f func(a int, acc int) int, init int) int {
	acc := init
	for _, a := range s {
		acc = f(a, acc)
	}
	return acc
}

func solve2(input string) (int, error) {
	lines := strings.Split(input, "\n")

	// convert to grid, char per char
	grid := make([][]string, len(lines))
	for i, l := range lines {
		grid[i] = make([]string, len(l))
		for j, r := range l {
			grid[i][j] = string(r)
		}
	}

	fmt.Printf("%v\n", grid)

	probs := [][]int{}
	var prob []int
	for col, _ := range grid[0] {
		numberStr := ""
		for row := range len(grid) { // fixed col
			digitStr := grid[row][col]

			if row == len(grid)-1 {
				// final row, string is just spaces, move to new problem
				if strings.ReplaceAll(numberStr, " ", "") == "" {
					// go to next problem
					fmt.Printf("adding prob %v to problems\n", prob)
					probs = append(probs, prob)
					prob = []int{}
				} else {
					fmt.Printf("adding: %s to prob %v\n", numberStr, prob)
					num, _ := strconv.Atoi(strings.Trim(numberStr, " "))
					prob = append(prob, num)
				}
				if digitStr == "+" || digitStr == "*" {
					// do this later
				}
			} else {
				numberStr += digitStr
			}
		}
	}
	probs = append(probs, prob) // add last problem

	fmt.Printf("%v\n", probs)
	ops := strings.Fields(lines[len(lines)-1])
	fmt.Printf("%v\n", ops)
	total := 0
	for i, op := range ops {
		res := 0
		if op == "+" {
			res = fold(probs[i], func(a, acc int) int { return a + acc }, 0)
		} else {
			res = fold(probs[i], func(a, acc int) int { return a * acc }, 1)
		}
		total += res
	}

	return total, nil
}
