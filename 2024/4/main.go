package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, err := os.ReadFile("4/input.txt")
	if err != nil {
		panic(err)
	}
	output, err := solveXmas(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

func solve(input string) (int, error) {
	rows := strings.Split(input, "\n")
	words := []string{"XMAS", "SAMX"}
	count := 0
	for i, row := range rows {
		for j, _ := range row {
			if row[j] != words[0][0] && row[j] != words[1][0] {
				continue
			}
			// find horizontally
			if j+3 < len(row) {
				candidate := row[j : j+4]
				//fmt.Printf("horizontal candidate: %s\n", candidate)
				if slices.Contains(words, candidate) {
					fmt.Printf("found %s horizontally at (%d, %d)\n", candidate, i, j)
					count++
				}
			}
			// vertically
			if i+3 < len(rows) {
				candidate := string(row[j]) + string(rows[i+1][j]) + string(rows[i+2][j]) + string(rows[i+3][j])
				if slices.Contains(words, candidate) {
					fmt.Printf("found %s vertically at (%d, %d)\n", candidate, i, j)
					count++
				}
			}
			// find in 2 diagonals. the other two are searched by searching backwards.
			if i+3 < len(rows) && j+3 < len(row) {
				candidate := string(row[j]) + string(rows[i+1][j+1]) + string(rows[i+2][j+2]) + string(rows[i+3][j+3])
				if slices.Contains(words, candidate) {
					fmt.Printf("found %s diagonally down at (%d, %d)\n", candidate, i, j)
					count++
				}
			}
			if i-3 >= 0 && j+3 < len(row) {
				candidate := string(row[j]) + string(rows[i-1][j+1]) + string(rows[i-2][j+2]) + string(rows[i-3][j+3])
				if slices.Contains(words, candidate) {
					fmt.Printf("found %s diagonally up at (%d, %d)\n", candidate, i, j)
					count++
				}
			}
		}
	}
	return count, nil
}

func solveXmas(input string) (int, error) {
	rows := strings.Split(input, "\n")
	count := 0
	around := []string{"MMSS", "MSSM", "SSMM", "SMMS"}
	for i, row := range rows {
		for j, _ := range row {
			if string(row[j]) != "A" {
				continue
			}
			// find surrounding square
			if j+1 >= len(row) || j-1 < 0 || i+1 >= len(rows) || i-1 < 0 {
				fmt.Printf("skipping (%d, %d)\n", i, j)
				continue
			}
			candidate := string(rows[i-1][j-1]) + string(rows[i-1][j+1]) + string(rows[i+1][j+1]) + string(rows[i+1][j-1])
			fmt.Printf("candidate %s at (%d, %d)\n", candidate, i, j)
			if slices.Contains(around, candidate) {
				fmt.Printf("found %s vertically at (%d, %d)\n", candidate, i, j)
				count++
			}
		}
	}
	return count, nil
}
