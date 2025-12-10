package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("8/input.txt")
	if err != nil {
		panic(err)
	}
	output, err := solve1(string(input), 1000000000, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

type Pair struct {
	a Pos
	b Pos
}

type Pos struct {
	x, y, z int64
}

func (p *Pos) dist(p2 Pos) int64 {
	return (p.x-p2.x)*(p.x-p2.x) + (p.y-p2.y)*(p.y-p2.y) + (p.z-p2.z)*(p.z-p2.z)
}

func solve1(input string, conns int, numCircuits int) (int, error) {
	lines := strings.Split(input, "\n")

	boxes := []Pos{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		boxes = append(boxes, Pos{x: int64(x), y: int64(y), z: int64(z)})
	}
	fmt.Printf("%d boxes\n", len(boxes)) //boxes)

	// calculate pairwise distances (n^2)
	mins := []Pair{}
	for i, b1 := range boxes {
		for j, b2 := range boxes {
			if i < j {
				fmt.Printf("dist from %v to %v is %d\n", b1, b2, b1.dist(b2))
				mins = append(mins, Pair{a: b1, b: b2})
			}
		}
	}
	slices.SortFunc(mins, func(pair1 Pair, pair2 Pair) int {
		return int(pair1.a.dist(pair1.b) - pair2.a.dist(pair2.b))
	})

	fmt.Printf("sorted pairwise distances: %v\n", mins)

	// groups can be a set of Pos's
	circuits := map[Pos]map[Pos]bool{} // from one box to circuit
	for _, box := range boxes {
		circuits[box] = map[Pos]bool{box: true} // single-box circuits to start
	}

	// connect!
	for i, pair := range mins {
		if i >= conns { // num connections
			break
		}
		// join circuits
		fmt.Printf("joining %v and %v\n", circuits[pair.a], circuits[pair.b])
		for box, _ := range circuits[pair.b] {
			circuits[pair.a][box] = true
		}
		if len(circuits[pair.a]) == len(boxes) {
			// one single circuit! exit now!
			return int(pair.a.x * pair.b.x), nil
		}
		// update pointers
		for box, _ := range circuits[pair.a] {
			circuits[box] = circuits[pair.a]
		}
	}

	fmt.Printf("circuits:\n")
	for _, c := range circuits {
		fmt.Printf("%v\n", c)
	}

	sizes := []int{}
	seenHashes := map[int64]bool{}
	for _, circuit := range circuits {
		hash := int64(1)
		for box, _ := range circuit {
			hash *= 3*box.x + 5*box.y + 7*box.z
		}
		if _, ok := seenHashes[hash]; ok {
			continue
		}
		sizes = append(sizes, len(circuit))
		seenHashes[hash] = true
	}
	slices.Sort(sizes)
	fmt.Printf("sorted circuit sizes: %v\n", sizes)

	total := 1
	for i, _ := range sizes {
		if i >= numCircuits { // 3 largest
			break
		}
		total *= sizes[len(sizes)-i-1]
	}

	return total, nil
}

func solve2(input string) (int, error) {
	lines := strings.Split(input, "\n")

	numPaths := map[int]int{} // number of paths to end up at each column position
	for i, line := range lines {
		fmt.Printf("processing line %d out of %d\n", i, len(lines))
		for j, cell := range line {
			if string(cell) == "S" {
				numPaths[j] = 1
			}

			if string(cell) == "^" {
				numPaths[j-1] += numPaths[j]
				numPaths[j+1] += numPaths[j]
				numPaths[j] = 0
			}
		}
	}

	total := 0
	for _, num := range numPaths {
		total += num
	}

	return total, nil
}
