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
	// fmt.Println(output)
	// output, err = solve2(string(input))
	// if err != nil {
	// 	panic(err)
	// }
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
	fmt.Printf("bank is %s\n", bank)
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

		// find the largest digit. if multiple, find the first instance
		// map[int] -> []int (positions)
		// then, go from 9 -> 0, find first position of first digit, use it.
		// we can be greedy in this way.
		// then, find next highest, but with strictly higher position
		// if we make it to 12 digits, that's it.
		// if we don't, then we have to start again, but with a lower first digit.

		// first, preprocess bank and build map
		bank := []int{}
		digits := map[int][]int{} // <- in increasing order
		maxSeen := 0
		for pos := range bankStr {
			digit, _ := strconv.Atoi(string(bankStr[pos]))
			digits[digit] = append(digits[digit], pos)
			bank = append(bank, digit)
			maxSeen = int(math.Max(float64(maxSeen), float64(digit)))
		}
		positions := []int{}
		for _, d := range []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0} {
			positions = append(positions, digits[d]...)
		}

		maxes := map[int]map[int]int{}
		for i, _ := range bank {
			fillMaxes(bank, maxes, len(bank)-i-1, 12)
		}

		fmt.Printf("map for %s is:\n%v\n", bankStr, digits)
		fmt.Printf("int bank is: %v; positions are: %v\n", bank, positions)

		mJolt1 := maxJolt1(bankStr)
		mJolt2 := maxes[0][12]
		fmt.Printf("max jolt for %s\n  method 1: %d\n  method 2: %d\n", bankStr, mJolt1, mJolt2)

		joltSum += mJolt2

		// now, attempt to build starting from top digits
		// firstDigit := maxSeen
		// for {
		// 	mJolt, found := maxJolt(digits, firstDigit, 12)
		// 	if found {
		// 		// mJolt1 := maxJolt1(bank)
		// 		// if mJolt != mJolt1 {
		// 		// 	fmt.Printf("max jolt for %s\n  method 1: %d\n  method 2: %d\n", bank, mJolt1, mJolt)
		// 		// }
		// 		joltSum += mJolt
		// 		break
		// 	} else {
		// 		firstDigit-- // should never get to 0
		// 	}
		// }

		// break
	}

	return joltSum, nil
}

func fillMaxes(bank []int, maxes map[int]map[int]int, pos int, upto int) {
	// build maxes from the right side to the left side
	// maxes[pos] -> max int of each length, starting from pos
	// for example, if maxes[13][3] = 253,
	// it means that 253 is the max jolt we've seen when using 3 batteries, starting from the 13th onwards
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

func maxJoltRecursive(bank []int, positions []int, minPos int, batts int) int {
	for _, pos := range positions {
		if pos >= minPos {
			if batts == 1 {
				fmt.Printf("picking %d\n", bank[pos])
				return bank[pos]
			} else {
				return (10^batts)*bank[pos] + maxJoltRecursive(bank, positions, pos+1, batts-1)
			}
		}
	}

	fmt.Println("NOT POSSIBRRRR")
	return 0 // not possible
}

func maxJolt(digits map[int][]int, firstDigit int, numBatteries int) (int, bool) {
	chosen := []int{}
	minPos := 0
	maxDigit := firstDigit
	for {
		if len(digits[maxDigit]) > 0 {
			fmt.Printf("chosen so far: %v; minPos: %d; maxDigit: %d\n", chosen, minPos, maxDigit)
			found := false
			for _, pos := range digits[maxDigit] {
				if pos >= minPos {
					chosen = append(chosen, maxDigit)
					minPos = pos + 1
					found = true
					break
				}
			}

			if !found {
				// no more maxDigits left to use, continue with maxDigits-1, unless we're at 0
				maxDigit--
				if maxDigit < 0 {
					return 0, false // no jolt found for numBatteries and firstDigit
				}
			} else { // we added a digit to the chosen ones
				if len(chosen) == numBatteries { // we're done
					joltStr := ""
					for _, d := range chosen {
						joltStr += strconv.Itoa(d)
					}
					maxJolt, _ := strconv.Atoi(joltStr)
					return maxJolt, true
				} else {
					// continue to add, but now maxDigit is 9 again
					maxDigit = 9
				}
			}
		} else { // no more options left
			// no more maxDigits left to use, continue with maxDigits-1, unless we're at 0
			maxDigit--
			if maxDigit < 0 {
				return 0, false // no jolt found for numBatteries and firstDigit
			}
		}
	}
}
