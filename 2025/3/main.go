package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("3/input.txt")
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

	joltSum := 0
	for _, bank := range lines {
		joltSum += maxJolt1(bank)
	}

	return joltSum, nil
}

func maxJolt1(bank string) int {
	maxJolt := 0
	// fmt.Printf("bank is %s\n", bank)
	for i := range bank {
		for j := i + 1; j < len(bank); j++ {
			// fmt.Printf("checking pos %d %c and pos %d %c\n", i, bank[i], j, bank[j])
			jolt, _ := strconv.Atoi(string(bank[i]) + string(bank[j]))
			if jolt > maxJolt {
				maxJolt = jolt
			}
		}
	}
	return maxJolt
}

func solve2(input string) (int, error) {
	lines := strings.Split(input, "\n")

	joltSum := 0
	for _, bankStr := range lines {

		// first, preprocess bank and build map
		bank := []int{}
		for pos := range bankStr {
			digit, _ := strconv.Atoi(string(bankStr[pos]))
			bank = append(bank, digit)
		}

		// build maxes from the right side to the left side
		// maxes[pos] -> max jolt of each length, starting from pos
		// for example, if maxes[13][3] = 253,
		// it means that 253 is the max jolt we've seen when using 3 batteries, starting from the 13th onwards
		maxes := map[int]map[int]int{}
		for i := range bank {
			fillMaxes(bank, maxes, len(bank)-i-1, 12)
		}

		mJolt1 := maxJolt1(bankStr)
		mJolt2 := maxes[0][12]
		fmt.Printf("max jolt for %s\n  method 1: %d\n  method 2: %d\n", bankStr, mJolt1, mJolt2)

		joltSum += mJolt2
	}

	return joltSum, nil
}

func fillMaxes(bank []int, maxes map[int]map[int]int, pos int, upto int) {
	// fmt.Printf("filling for pos %d, maxes are: %v\n", pos, maxes)

	if _, ok := maxes[pos]; !ok {
		maxes[pos] = map[int]int{}
	}
	// base case
	if pos == len(bank)-1 {
		maxes[pos][1] = bank[pos]
	} else {
		for l := range upto { // l+1 goes from 1 to upto (inclusive)
			if l+1 == 1 { // base case, length 1
				maxes[pos][l+1] = intMax(bank[pos], maxes[pos+1][l+1])
			} else { // length is 2 or more
				joltUsingPos, _ := strconv.Atoi(strconv.Itoa(bank[pos]) + strconv.Itoa(maxes[pos+1][l]))
				maxes[pos][l+1] = intMax(joltUsingPos, maxes[pos+1][l+1])
			}
		}
	}
}

func intMax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
