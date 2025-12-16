package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("9/input.txt")
	if err != nil {
		panic(err)
	}
	output, err := solve2(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

type Point struct {
	x, y int
}

type Pair struct {
	a, b Point
}

func (p *Pair) inline(pt Point) bool {
	if p.a.x == p.b.x {
		return pt.x == p.a.x && (pt.y <= p.a.y && pt.y >= p.b.y || pt.y >= p.a.y && pt.y >= p.b.y)
	} else if p.a.y == p.b.y {
		return pt.y == p.a.y && (pt.x <= p.a.x && pt.x >= p.b.x || pt.x >= p.a.x && pt.x >= p.b.x)
	}
	return false
}

func solve1(input string) (int, error) {
	lines := strings.Split(input, "\n")

	tiles := []Point{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		tiles = append(tiles, Point{x: x, y: y})
	}

	// find pairwise max product
	max := 0
	for _, p1 := range tiles {
		for _, p2 := range tiles {
			area := int((math.Abs(float64(p1.x-p2.x)) + 1) * math.Abs(float64(p1.y-p2.y)+1))
			if area > max {
				max = area
			}
		}
	}

	return max, nil
}

func area(p1, p2 Point) int {
	return int((math.Abs(float64(p1.x-p2.x)) + 1) * math.Abs(float64(p1.y-p2.y)+1))
}

func solve2(input string) (int, error) {
	lines := strings.Split(input, "\n")

	reds := []Point{}
	redOrGreen := map[Point]bool{}
	minX := 9999999999
	minY := 9999999999
	maxX := 0
	maxY := 0
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}

		pt := Point{x: x, y: y}
		reds = append(reds, pt)
		redOrGreen[pt] = true

		nextParts := strings.Split(lines[(i+1)%len(lines)], ",")
		x2, _ := strconv.Atoi(nextParts[0])
		y2, _ := strconv.Atoi(nextParts[1])
		next := Point{x: x2, y: y2}

		if pt.x == next.x {
			min := int(math.Min(float64(pt.y), float64(next.y)))
			max := int(math.Max(float64(pt.y), float64(next.y)))
			for i := range max - min - 1 {
				redOrGreen[Point{x: pt.x, y: min + i + 1}] = true
			}
		} else if pt.y == next.y {
			min := int(math.Min(float64(pt.x), float64(next.x)))
			max := int(math.Max(float64(pt.x), float64(next.x)))
			for i := range max - min - 1 {
				redOrGreen[Point{x: min + i + 1, y: pt.y}] = true
			}
		}
	}

	// fill green tiles inside loop, using vertical lines
	// pairs := []Pair{}
	// insideGreens := map[Point]bool{}
	// for i := range maxX - minX + 1 {
	// 	x := minX + i
	// 	var prev *Point
	// 	for j := range maxY - minY + 1 {
	// 		y := minY + j
	// 		pt := Point{x: x, y: y}
	// 		// fmt.Printf("checking %v\n", pt)
	// 		if redOrGreen[pt] {
	// 			if prev == nil { // start a match
	// 				prev = &pt
	// 			} else {
	// 				// fmt.Printf(" filling in between %v and %v\n", prev, pt)
	// 				pairs = append(pairs, Pair{a: *prev, b: pt})
	// 				// found matching pair; fill in between and reset
	// 				// for k := range pt.y - prev.y {
	// 				// 	insideGreens[Point{x: x, y: prev.y + k}] = true
	// 				// }
	// 				prev = nil
	// 			}
	// 		}
	// 	}
	// }

	// idea: flood the outside from (0,0). whatever is not flooded is inside

	// queue := []Point{{x: minX - 1, y: minY - 1}}
	// seen := map[Point]bool{}
	// plain := map[Point]bool{} // assume 0,0 is plain
	// for {
	// 	// fmt.Printf("queue: %v\nseen: %v\nplain:%v\n\n", queue, seen, plain)
	// 	if len(queue) == 0 {
	// 		break
	// 	}
	// 	curr := queue[0]
	// 	queue = slices.Delete(queue, 0, 1)
	// 	if !redOrGreen[curr] {
	// 		plain[curr] = true
	// 		// only expand if we're plain
	// 		neighbors := []Point{
	// 			Point{x: curr.x + 1, y: curr.y},     // right
	// 			Point{x: curr.x + 1, y: curr.y + 1}, // diag
	// 			Point{x: curr.x, y: curr.y + 1},     // down
	// 		}
	// 		for _, n := range neighbors {
	// 			if !seen[n] && n.x <= maxX && n.y <= maxY {
	// 				queue = append(queue, n)
	// 			}
	// 		}
	// 	}
	// 	seen[curr] = true
	// }

	// idea2: if there is a RED inside the rectangle (not perimeter), then it's invalid

	fmt.Printf("outline: %v\n", redOrGreen)
	// fmt.Printf("inside greens: %v\n", insideGreens)
	// fmt.Printf("pairs: %v\n", pairs)

	// for pt := range insideGreens {
	// 	redOrGreen[pt] = true
	// }
	// fmt.Printf("full: %v\n", redOrGreen)
	pairs := []Pair{}
	for _, p1 := range reds {
		for _, p2 := range reds {
			pairs = append(pairs, Pair{a: p1, b: p2})
		}
	}
	slices.SortFunc(pairs, func(pair1, pair2 Pair) int {
		return area(pair2.a, pair2.b) - area(pair1.a, pair1.b)
	})

	limit := 1703306226 // from previous run
	for _, pair := range pairs {
		area := area(pair.a, pair.b)
		if area > limit {
			continue
		}
		fmt.Printf("checking %v and %v with area %d\n", pair.a, pair.b, area)
		if insideGreen(pair.a, pair.b, redOrGreen) { // return first
			return area, nil
		}
	}

	// find pairwise max product, if all tiles within are green
	// for _, p1 := range reds {
	// 	for _, p2 := range reds {
	// 		area := int((math.Abs(float64(p1.x-p2.x)) + 1) * math.Abs(float64(p1.y-p2.y)+1))
	// 		fmt.Printf("max so far: %d. checking %v and %v with area %d\n", max, p1, p2, area)
	// 		if area > max && insideGreen(p1, p2, redOrGreen) {
	// 			max = area
	// 			// check if all non-corners are green
	// 			// minX := int(math.Min(float64(p1.x), float64(p2.x)))
	// 			// maxX := int(math.Max(float64(p1.x), float64(p2.x)))
	// 			// minY := int(math.Min(float64(p1.y), float64(p2.y)))
	// 			// maxY := int(math.Max(float64(p1.y), float64(p2.y)))

	// 			// check perimeter is red or green
	// 			// perimGood := true
	// 			// for i := range maxX - minX {
	// 			// 	if !redOrGreen[Point{x: minX + i, y: p1.y}] || !redOrGreen[Point{x: minX + i, y: p2.y}] {
	// 			// 		perimGood = false
	// 			// 		break
	// 			// 	}
	// 			// }
	// 			// for j := range maxY - minY {
	// 			// 	perim1 := Point{x: p1.x, y: minY + j}
	// 			// 	perim2 := Point{x: p2.x, y: minY + j}
	// 			// 	if !redOrGreen[perim1] || !redOrGreen[perim2] {
	// 			// 		fmt.Printf("found point in perimeter that is not green or red: %v ot %v\n", perim1, perim2)
	// 			// 		perimGood = false
	// 			// 		break
	// 			// 	}
	// 			// }
	// 			// if !perimGood {
	// 			// 	continue
	// 			// }

	// 			// noPlains := true
	// 			// insideRed := false
	// 			// see if there's a red tile fully inside
	// 			// for _, red := range reds {
	// 			// 	if red.x > minX && red.x < maxX && red.y > minY && red.y < maxY {
	// 			// 		insideRed = true
	// 			// 		break
	// 			// 	}
	// 			// }
	// 			// 	allFilled := true
	// 			// outer:
	// 			// 	for i := range maxX - minX + 1 {
	// 			// 		for j := range maxY - minY + 1 {
	// 			// 			pt := Point{x: minX + i, y: minY + j}
	// 			// 			// fmt.Printf("checking %v, red/green? %t\n", pt, redOrGreen[pt])
	// 			// 			if !redOrGreen[pt] {
	// 			// 				found := false
	// 			// 				for _, pair := range pairs {
	// 			// 					if pair.inline(pt) {
	// 			// 						found = true
	// 			// 					}
	// 			// 				}
	// 			// 				if !found {
	// 			// 					allFilled = false
	// 			// 					break outer
	// 			// 				}
	// 			// 				// noPlains = false
	// 			// 			}

	// 			// 			// if slices.Contains(reds, pt) && pt.x != p1.x && pt.x != p2.x && pt.y != p1.y && pt.y != p2.y {
	// 			// 			// 	insideRed = true
	// 			// 			// 	break outer
	// 			// 			// }
	// 			// 		}
	// 			// 	}

	// 			// 	if allFilled {
	// 			// 		max = area
	// 			// 	}
	// 		}
	// 	}
	// }

	return 0, nil
}

func insideGreen(p1 Point, p2 Point, redOrGreen map[Point]bool) bool {
	minX := int(math.Min(float64(p1.x), float64(p2.x)))
	maxX := int(math.Max(float64(p1.x), float64(p2.x)))
	minY := int(math.Min(float64(p1.y), float64(p2.y)))
	maxY := int(math.Max(float64(p1.y), float64(p2.y)))

	// fill green tiles inside loop, using vertical lines
	// pairs := []Pair{}
	// insideGreens := map[Point]bool{}
	for j := range maxY - minY + 1 {
		y := minY + j
		var prev *Point
		isOk := false
		for i := range maxX - minX + 1 {
			x := minX + i
			pt := Point{x: x, y: y}
			// fmt.Printf("checking %v\n", pt)
			if redOrGreen[pt] {
				if prev == nil { // start a match
					prev = &pt
					isOk = true
				} else {
					// fmt.Printf(" filling in between %v and %v\n", prev, pt)
					// pairs = append(pairs, Pair{a: *prev, b: pt})
					// found matching pair; fill in between and reset
					// for k := range pt.y - prev.y {
					// 	insideGreens[Point{x: x, y: prev.y + k}] = true
					// }
					prev = nil
					isOk = false
				}
			} else {
				if !isOk {
					return false
				}
			}
		}
	}
	return true
}
