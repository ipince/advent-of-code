package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("6/input.txt")
	if err != nil {
		panic(err)
	}
	output, err := solve(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

type coords struct {
	x int
	y int
}
type coordsDir struct {
	x   int
	y   int
	dir int
}

func solve(input string) (int, error) {
	rows := strings.Split(input, "\n")

	grid := [][]string{}
	var initial coords
	for i, row := range rows {
		cells := []string{}
		for j, cell := range row {
			if string(cell) == "^" {
				initial = coords{x: i, y: j}
			}
			cells = append(cells, string(cell))
		}
		grid = append(grid, cells)
	}

	visited, _, _ := simulate(grid, initial)

	// simulate obstacle along visited path
	obstacleOptions := 0
	for v, _ := range visited {
		if v != initial {
			grid[v.x][v.y] = "#"
			fmt.Printf("simulating grid with obstacle in (%d, %d)\n", v.x, v.y)
			_, looped, _ := simulate(grid, initial)
			if looped {
				fmt.Printf("    LOOPED!!\n")
				obstacleOptions++
			}
			grid[v.x][v.y] = "."
		}
	}

	return obstacleOptions, nil
}

func simulate(grid [][]string, initial coords) (map[coords]bool, bool, error) {
	visited := map[coords]bool{}
	visitedDir := map[coordsDir]bool{}
	position := initial
	dir := 0 // 0 1 2 3 4
	for {
		//fmt.Printf("visiting (%d, %d)\n", position.x, position.y)
		visited[position] = true

		if _, ok := visitedDir[coordsDir{
			x:   position.x,
			y:   position.y,
			dir: dir,
		}]; ok { // we've been here before...
			return nil, true, nil
		}
		visitedDir[coordsDir{
			x:   position.x,
			y:   position.y,
			dir: dir,
		}] = true

		next := newPos(position, dir)
		if inside(grid, next) {
			if grid[next.x][next.y] == "#" {
				//fmt.Printf("turning right on (%d, %d)\n", position.x, position.y)
				dir = (dir + 1) % 4
				next = newPos(position, dir)
				if !inside(grid, next) {
					fmt.Printf("exited at (%d, %d)\n", position.x, position.y)
					break
				}
			}
		} else { // exited!
			fmt.Printf("exited at (%d, %d)\n", position.x, position.y)
			break
		}
		position = next
	}
	return visited, false, nil
}

func inside(grid [][]string, next coords) bool {
	return next.x >= 0 && next.x < len(grid) && next.y >= 0 && next.y < len(grid[0])
}

func newPos(pos coords, dir int) coords {
	var next coords
	switch dir {
	case 0: // up
		next = coords{x: pos.x - 1, y: pos.y}
	case 1: // right
		next = coords{x: pos.x, y: pos.y + 1}
	case 2: // down
		next = coords{x: pos.x + 1, y: pos.y}
	case 3: // left
		next = coords{x: pos.x, y: pos.y - 1}
	}
	return next
}
