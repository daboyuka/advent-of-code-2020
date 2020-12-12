package aoc2020

import (
	"fmt"

	. "aoc2020/helpers"
)

type instr struct {
	Kind string
	Amt  int
}

type dir int

const (
	east = dir(iota)
	north
	west
	south
)

var xdelta = [...]int{
	east:  1,
	west:  -1,
	north: 0,
	south: 0,
}
var ydelta = [...]int{
	north: 1,
	south: -1,
	east:  0,
	west:  0,
}

func ParseInstructions(lines []string) (out []instr) {
	for _, line := range lines {
		kind, amt := line[:1], Atoi(line[1:])
		out = append(out, instr{Kind: kind, Amt: amt})
	}
	return out
}

func Problem12a(lines []string) {
	ins := ParseInstructions(lines)

	x, y := 0, 0
	d := east
	for _, in := range ins {
		switch in.Kind {
		case "N":
			y += in.Amt
		case "S":
			y -= in.Amt
		case "E":
			x += in.Amt
		case "W":
			x -= in.Amt

		case "L":
			d += dir(in.Amt / 90)
			d %= 4
		case "R":
			d -= dir(in.Amt / 90)
			for d < 0 {
				d += 4
			}
		case "F":
			x += in.Amt * xdelta[d]
			y += in.Amt * ydelta[d]
		}
	}

	fmt.Println(x, y, Abs(x)+Abs(y))
}

func Problem12b(lines []string) {
	ins := ParseInstructions(lines)

	x, y := 0, 0
	wx, wy := 10, 1
	for _, in := range ins {
		switch in.Kind {
		case "N":
			wy += in.Amt
		case "S":
			wy -= in.Amt
		case "E":
			wx += in.Amt
		case "W":
			wx -= in.Amt

		case "L":
			for i := 0; i < in.Amt/90; i++ {
				wx, wy = -wy, wx
			}
		case "R":
			for i := 0; i < in.Amt/90; i++ {
				wx, wy = wy, -wx
			}
		case "F":
			x += in.Amt * wx
			y += in.Amt * wy
		}

		fmt.Println(x, y, wx, wy)
	}

	fmt.Println(x, y, Abs(x)+Abs(y))
}
