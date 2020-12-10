package aoc2020

import (
	"fmt"

	. "aoc2020/helpers"
)

type thingy string

func ParseThingies(lines []string) (out []thingy) {
	for _, line := range lines {
		out = append(out, thingy(line))
	}
	return out
}

func Problem11a(lines []string) {
	x := ParseThingies(lines)
	_ = x
}

func Problem11b(lines []string) {
	x := ParseThingies(lines)
	_ = x
}
