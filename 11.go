package aoc2020

import (
	"fmt"
	"reflect"
)

type airport [][]int // 0 -> floor, 1 -> empty, 2 -> occupied

func (a airport) occupied(row, col int) (occ, unocc, inbounds bool) {
	if row < 0 || row >= len(a) {
		return false, false, false
	}
	ar := a[row]
	if col < 0 || col >= len(ar) {
		return false, false, false
	}
	return ar[col] == 2, ar[col] == 1, true
}

func (a airport) neighbors(row, col int) (n int) {
	for _, rd := range [...]int{-1, 0, 1} {
		for _, cd := range [...]int{-1, 0, 1} {
			if rd == 0 && cd == 0 {
				continue
			} else if occ, _, _ := a.occupied(row+rd, col+cd); occ {
				n++
			}
		}
	}
	return n
}

func (a airport) neighborsLong(row, col int) (n int) {
	for _, rd := range [...]int{-1, 0, 1} {
		for _, cd := range [...]int{-1, 0, 1} {
			if rd == 0 && cd == 0 {
				continue
			}

			for x := 1; ; x++ {
				occ, unocc, inbounds := a.occupied(row+rd*x, col+cd*x)
				if !inbounds || unocc {
					break
				} else if occ {
					n++
					break
				}
			}
		}
	}
	return n
}

func (a airport) copy() airport {
	a2 := make(airport, len(a))
	for i, r := range a {
		a2[i] = append([]int(nil), r...)
	}
	return a2
}

func (a airport) countOcc() (occ int) {
	for _, ar := range a {
		for _, seat := range ar {
			if seat == 2 {
				occ++
			}
		}
	}
	return occ
}

func (a airport) String() (out string) {
	for _, ar := range a {
		for _, seat := range ar {
			switch seat {
			case 0:
				out = out + string('.')
			case 1:
				out = out + string('L')
			case 2:
				out = out + string('#')
			}
		}
		out += "\n"
	}
	return out
}

func evolve(from airport) airport {
	to := make(airport, len(from))
	for i, r := range from {
		to[i] = make([]int, len(r))
	}

	for row := range to {
		for col := range to[row] {
			n := from.neighbors(row, col)
			if from[row][col] == 0 {
				to[row][col] = 0
			} else if n == 0 {
				to[row][col] = 2
			} else if n >= 4 {
				to[row][col] = 1
			} else {
				to[row][col] = from[row][col]
			}
		}
	}

	return to
}

func ParseAirport(lines []string) (out airport) {
	for _, line := range lines {
		outrow := make([]int, len(line))
		for i, c := range line {
			if c == 'L' {
				outrow[i] = 1
			} else {
				outrow[i] = 0
			}
		}
		out = append(out, outrow)
	}

	return out
}

func Problem11a(lines []string) {
	ap := ParseAirport(lines)

	cur := ap
	for {
		fmt.Println(cur)
		next := evolve(cur)
		if reflect.DeepEqual(cur, next) {
			fmt.Println(cur.countOcc())
			return
		}
		cur = next
	}
}

func evolve2(from airport) airport {
	to := make(airport, len(from))
	for i, r := range from {
		to[i] = make([]int, len(r))
	}

	for row := range to {
		for col := range to[row] {
			n := from.neighborsLong(row, col)
			if from[row][col] == 0 {
				to[row][col] = 0
			} else if from[row][col] == 1 && n == 0 {
				to[row][col] = 2
			} else if from[row][col] == 2 && n >= 5 {
				to[row][col] = 1
			} else {
				to[row][col] = from[row][col]
			}
		}
	}

	return to
}

func Problem11b(lines []string) {
	ap := ParseAirport(lines)

	cur := ap
	for {
		fmt.Println(cur)
		next := evolve2(cur)
		if reflect.DeepEqual(cur, next) {
			fmt.Println(cur.countOcc())
			return
		}
		cur = next
	}
}
