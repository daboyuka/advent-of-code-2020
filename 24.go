package aoc2020

import (
	"fmt"
	//. "aoc2020/helpers"
)

type path []string

func ParsePaths(lines []string) (paths []path) {
	for _, line := range lines {
		var p path
		for i := 0; i < len(line); i++ {
			if line[i] == 'n' || line[i] == 's' {
				p = append(p, line[i:i+2])
				i++
			} else {
				p = append(p, line[i:i+1])
			}
		}
		paths = append(paths, p)
	}
	return paths
}

type offset struct {
	E  int
	NW int
}

func (p path) ToOffset() (off offset) {
	for _, x := range p {
		switch x {
		case "e":
			off.E++
		case "w":
			off.E--
		case "ne":
			off.NW++
			off.E++
		case "nw":
			off.NW++
		case "se":
			off.NW--
		case "sw":
			off.NW--
			off.E--
		}
	}
	return off
}

func (off offset) Neighbors() []offset {
	return []offset{
		{off.E, off.NW + 1},
		{off.E, off.NW - 1},
		{off.E + 1, off.NW + 1},
		{off.E - 1, off.NW - 1},
		{off.E - 1, off.NW},
		{off.E + 1, off.NW},
	}
}

func Problem24a(lines []string) {
	paths := ParsePaths(lines)

	blacks := map[offset]bool{}
	for _, p := range paths {
		off := p.ToOffset()
		blacks[off] = !blacks[off]
	}

	blackCount := 0
	for _, isBlack := range blacks {
		if isBlack {
			blackCount++
		}
	}

	fmt.Println(blackCount)
}

func BlackNeighbors(off offset, m map[offset]bool) (n int) {
	for _, offN := range off.Neighbors() {
		if m[offN] {
			n++
		}
	}
	return n
}

func NextColor(wasBlack bool, n int) bool {
	if wasBlack && (n == 0 || n > 2) {
		return false
	}
	if !wasBlack && n == 2 {
		return true
	}
	return wasBlack
}

func CountBlack(m map[offset]bool) (blackCount int) {
	for _, isBlack := range m {
		if isBlack {
			blackCount++
		}
	}
	return blackCount
}

func Problem24b(lines []string) {
	paths := ParsePaths(lines)

	cur := map[offset]bool{}
	for _, p := range paths {
		off := p.ToOffset()
		cur[off] = !cur[off]
	}

	fmt.Println(CountBlack(cur))

	const days = 100
	for i := 0; i < days; i++ {
		next := map[offset]bool{}
		for off := range cur {
			next[off] = NextColor(cur[off], BlackNeighbors(off, cur))
			for _, off2 := range off.Neighbors() {
				next[off2] = NextColor(cur[off2], BlackNeighbors(off2, cur))
			}
		}
		cur = next

		fmt.Println(i+1, CountBlack(cur))

	}

	blackCount := 0
	for _, isBlack := range cur {
		if isBlack {
			blackCount++
		}
	}

	fmt.Println(CountBlack(cur))
}
