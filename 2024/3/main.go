package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("3/input.txt")
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
	total, sum := 0, 0
	mode := true
	for _, line := range strings.Split(input, "\n") {
		sum, mode = mulsumDoDont(line, mode)
		total += sum
	}
	return total, nil
}

var mulRegex = regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

func mulsum(line string) int {
	if !mulRegex.MatchString(line) {
		return 0
	}
	total := 0
	matches := mulRegex.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		l, _ := strconv.Atoi(match[1])
		r, _ := strconv.Atoi(match[2])
		total += l * r
	}
	return total
}

var doRegex = regexp.MustCompile("do\\(\\)")
var dontRegex = regexp.MustCompile("don't\\(\\)")

func mulsumDoDont(line string, start bool) (int, bool) {
	if !mulRegex.MatchString(line) {
		return 0, start
	}
	dos := doRegex.FindAllStringSubmatchIndex(line, -1)
	donts := dontRegex.FindAllStringSubmatchIndex(line, -1)
	muls := mulRegex.FindAllStringSubmatchIndex(line, -1)

	at := map[int]bool{}
	for _, d := range dos {
		at[d[0]] = true
	}
	for _, d := range donts {
		at[d[0]] = false
	}
	mulmap := map[int][]int{}
	for _, m := range muls {
		mulmap[m[0]] = m
	}

	indeces := []int{}
	indeces = append(indeces, first(dos)...)
	indeces = append(indeces, first(donts)...)
	indeces = append(indeces, first(muls)...)
	sort.Ints(indeces)

	mode := start
	total := 0
	fmt.Println(dos)
	fmt.Println(donts)
	fmt.Println(muls)
	fmt.Println(indeces)
	for _, idx := range indeces {
		if m, ok := at[idx]; ok {
			fmt.Printf("switching mode from %t to %t\n", mode, m)
			mode = m
			continue
		}
		mul := mulmap[idx]
		if mode {
			l, _ := strconv.Atoi(line[mul[2]:mul[3]])
			r, _ := strconv.Atoi(line[mul[4]:mul[5]])
			total += l * r
			fmt.Printf("added %d * %d\n", l, r)
		} else {
			fmt.Printf("skipping %s\n", line[mul[0]:mul[1]])
		}
	}
	return total, mode
}

func first(ll [][]int) []int {
	firsts := make([]int, len(ll))
	for i, l := range ll {
		firsts[i] = l[0]
	}
	return firsts
}
