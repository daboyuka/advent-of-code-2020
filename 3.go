package main

import "fmt"

type grid []string

func (g grid) at(row, col int) (tree bool) {
	return g[row][col%len(g[row])] == '#'
}

func countSlope(g grid, right, down int) (trees int) {
	row, col := 0, 0
	for {
		col += right
		row += down
		if row >= len(g) {
			return
		}

		if g.at(row, col) {
			trees++
		}
	}
}

func problem3a(lines []string) {
	fmt.Println(countSlope(grid(lines), 3, 1))
}

func problem3b(lines []string) {
	fmt.Println(1 *
		countSlope(grid(lines), 1, 1) *
		countSlope(grid(lines), 3, 1) *
		countSlope(grid(lines), 5, 1) *
		countSlope(grid(lines), 7, 1) *
		countSlope(grid(lines), 1, 2),
	)
}
