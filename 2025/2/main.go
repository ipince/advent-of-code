package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("2/input.txt")
	if err != nil {
		panic(err)
	}
	output, err := solve1(string(input))
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
	line := lines[0]
	ranges := strings.Split(line, ",")
	sum := 0
	for _, r := range ranges {
		ends := strings.Split(r, "-")
		start, _ := strconv.Atoi(ends[0])
		end, _ := strconv.Atoi(ends[1])

		curr := start
		for curr < end {
			if isInvalid2(curr) {
				fmt.Printf("isInvalid(%d)? true\n", curr)
				sum += curr
			}

			curr++
		}
	}
	return sum, nil
}

func isInvalid1(id int) bool {
	idStr := strconv.Itoa(id)
	if len(idStr)%2 == 0 {
		return idStr[:len(idStr)/2] == idStr[len(idStr)/2:]
	}
	return false
}

func isInvalid2(id int) bool {
	idStr := strconv.Itoa(id)

	prefixLength := 1
	for prefixLength <= len(idStr)/2 {
		if len(idStr)%prefixLength != 0 {
			prefixLength++
			continue
		}
		reps := len(idStr) / prefixLength
		prefix := idStr[:prefixLength]

		repeated := ""
		for range reps {
			repeated += prefix
		}
		if repeated == idStr {
			return true
		}

		prefixLength++
	}
	return false
}
