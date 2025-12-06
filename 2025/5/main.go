package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("5/input.txt")
	if err != nil {
		panic(err)
	}
	output, err := solve(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

type Range struct {
	min int
	max int
}

func solve(input string) (int, error) {
	lines := strings.Split(input, "\n")

	ranges := []Range{}
	pastDB := false
	count := 0
	for _, line := range lines {
		if line == "" {
			pastDB = true
			return allFresh(ranges), nil
			// continue
		}

		if !pastDB {
			parts := strings.Split(line, "-")
			min, _ := strconv.Atoi(parts[0])
			max, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, Range{min: min, max: max})
		} else {
			ingredient, _ := strconv.Atoi(line)

			if isFresh(ranges, ingredient) {
				count++
			}
		}
	}
	return count, nil
}

func isFresh(ranges []Range, ingredient int) bool {
	for _, r := range ranges {
		if ingredient >= r.min && ingredient <= r.max {
			return true
		}
	}
	return false
}

func print(ranges []Range) {
	fmt.Println()
	for _, r := range ranges {
		fmt.Printf("%d-%d\n", r.min, r.max)
	}
}

type Switch struct {
	mark int
	on   bool
}

func allFresh(ranges []Range) int {
	// sort the ranges by min, then max (just to debug)
	slices.SortFunc(ranges, func(a Range, b Range) int {
		if a.min == b.min {
			return a.max - b.max
		}
		return a.min - b.min
	})
	print(ranges)

	switches := []Switch{}
	// idea: make "switches" to switch from "OUT" to "IN", flatten and sort
	for _, r := range ranges {
		switches = append(switches, Switch{mark: r.min, on: true}, Switch{mark: r.max, on: false})
	}
	slices.SortFunc(switches, func(a Switch, b Switch) int {
		if a.mark == b.mark {
			if a.on != b.on { // put ON first
				if a.on {
					return -1
				} else {
					return 1
				}
			}
			return 0
		}
		return a.mark - b.mark
	})

	// go through switches and count when we're outside (or inside, but outside is easier)
	count := 0
	lastOff := 0
	depth := 0
	for _, s := range switches {
		fmt.Printf("count so far: %d; lastOff: %d; depth: %d. switch: (%d, %t)\n", count, lastOff, depth, s.mark, s.on)
		if s.on {
			depth++
			if depth == 1 { // we just entered a new range
				count += s.mark - lastOff - 1
			} // else, do nothing
		} else if !s.on {
			depth--
			if depth == 0 { // we're outside
				lastOff = s.mark
			}
		}
	}

	return switches[len(switches)-1].mark - count
}
