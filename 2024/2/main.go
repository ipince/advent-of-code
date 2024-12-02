package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("2/input.txt")
	if err != nil {
		panic(err)
	}
	output, err := solve(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

func solve(input string) (string, error) {
	numSafe := 0
	for _, line := range strings.Split(input, "\n") {
		levels := strings.Fields(line)

		isSafe := safe(levels, false, 0)
		if !isSafe {
			for i := 0; i < len(levels); i++ {
				saf := safe(levels, true, i)
				if saf {
					isSafe = true
					break
				}
			}
		}
		if isSafe {
			numSafe++
		}
	}
	return strconv.Itoa(numSafe), nil
}

func safe(levelsIn []string, skip bool, idx int) bool {
	levels := make([]string, len(levelsIn))
	copy(levels, levelsIn)
	if skip {
		levels = append(levels[:idx], levels[idx+1:]...)
	}
	dir := 0
	for i, level := range levels {
		if i == 0 {
			continue
		}

		l, _ := strconv.Atoi(level)
		p, _ := strconv.Atoi(levels[i-1])

		d := math.Abs(float64(l - p))
		if d < 1 || d > 3 {
			return false
		}
		if i == 1 {
			if l > p {
				dir = 1
			} else {
				dir = 0
			}
		}
		// i >= 1
		if l > p {
			if dir != 1 {
				return false
			}
		} else {
			if dir != 0 {
				return false
			}
		}
	}
	return true
}
