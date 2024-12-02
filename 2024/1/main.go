package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("1/input.txt")
	if err != nil {
		panic(err)
	}
	output, err := solveSecondPart(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

func solveFirstPart(input string) (string, error) {
	left := []int{}
	right := []int{}
	for _, line := range strings.Split(input, "\n") {
		points := strings.Fields(line)
		if len(points) != 2 {
			fmt.Println(points)
			return "", fmt.Errorf("invalid input: %s", line)
		}
		l, err := strconv.Atoi(points[0])
		if err != nil {
			return "", err
		}
		r, err := strconv.Atoi(points[1])
		if err != nil {
			return "", err
		}
		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	var sum int64 = 0
	for i, _ := range left {
		sum += int64(math.Abs(float64(right[i] - left[i])))
	}

	return fmt.Sprintf("%d", sum), nil
}

func solveSecondPart(input string) (string, error) {
	left := []int{}
	right := []int{}
	for _, line := range strings.Split(input, "\n") {
		points := strings.Fields(line)
		if len(points) != 2 {
			fmt.Println(points)
			return "", fmt.Errorf("invalid input: %s", line)
		}
		l, err := strconv.Atoi(points[0])
		if err != nil {
			return "", err
		}
		r, err := strconv.Atoi(points[1])
		if err != nil {
			return "", err
		}
		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	counts := map[int]int{}
	for _, r := range right {
		if _, ok := counts[r]; !ok {
			counts[r] = 0
		}
		counts[r]++
	}

	var sum int64 = 0
	for _, l := range left {
		if c, ok := counts[l]; ok {
			sum += int64(l * c)
		}
	}

	return fmt.Sprintf("%d", sum), nil
}
