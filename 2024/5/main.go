package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("5/input.txt")
	if err != nil {
		panic(err)
	}
	output1, output2, err := solve(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d, %d\n", output1, output2)
}

func solve(input string) (int, int, error) {
	rows := strings.Split(input, "\n")

	rules := make(map[string][]string) // A -> B means B has to be before A.
	sumPass := 0
	sumFixed := 0
	rulesDone := false
	for _, row := range rows {
		if row == "" {
			fmt.Printf("parsed rules: %v\n", rules)
			rulesDone = true
			continue
		}

		if !rulesDone {
			parts := strings.Split(row, "|")
			if len(parts) != 2 {
				fmt.Printf("incorrect rule: %s\n", row)
				continue
			}
			if _, ok := rules[parts[1]]; !ok {
				rules[parts[1]] = make([]string, 0)
			}
			rules[parts[1]] = append(rules[parts[1]], parts[0])
		} else {
			ok, mid := check(row, rules, false)
			if ok {
				fmt.Printf("update %s PASSES, adding %d\n", row, mid)
				sumPass += mid
			} else {
				_, mid2 := check(row, rules, true)
				sumFixed += mid2
			}
		}
	}
	return sumPass, sumFixed, nil
}

func check(update string, rules map[string][]string, fix bool) (bool, int) {
	indeces := map[string]int{} // assume each page only appears once
	parts := strings.Split(update, ",")
	for j, page := range parts {
		indeces[page] = j
	}
	// check rules are respected
	for page, i := range indeces {
		if rs, ok := rules[page]; ok {
			for _, r := range rs {
				if j, ok := indeces[r]; ok {
					if j >= i { // broke rule!
						fmt.Printf("update %s BROKE rule: %s|%s\n", update, page, r)
						if fix {
							// swap them and try again.. like bubblesort!
							parts[i] = parts[j]
							parts[j] = page
							newUpdate := strings.Join(parts, ",")
							fmt.Printf("attempting reordered %s\n", newUpdate)
							return check(newUpdate, rules, fix)
						} else {
							return false, 0
						}
					}
				}
			}
		}
	}
	mid, _ := strconv.Atoi(parts[len(parts)/2]) // assume always odd number of pages
	return true, mid
}
