package aoc2020

import (
	"fmt"
)

type pos struct{ W, Z, R, C int }

func (p pos) DoForNeighbors(inclSelf, inclW bool, f func(p pos)) {
	wlb, wub := 0, 0
	if inclW {
		wlb, wub = -1, 1
	}

	for dw := wlb; dw <= wub; dw++ {
		for dz := -1; dz <= 1; dz++ {
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if !inclSelf && dw == 0 && dz == 0 && dr == 0 && dc == 0 {
						continue
					}
					f(pos{p.W + dw, p.Z + dz, p.R + dr, p.C + dc})
				}
			}
		}
	}
}

type cubegrid map[pos]bool

func ParseFlatgrid(lines []string) (f cubegrid) {
	f = make(cubegrid, len(lines))
	for row, line := range lines {
		for col, c := range line {
			f[pos{0, 0, row, col}] = c == '#'
		}
	}
	return f
}

func (g cubegrid) NeighborCount(inclW bool, p pos) (n int) {
	p.DoForNeighbors(false, inclW, func(p pos) {
		if g[p] {
			n++
		}
	})
	return n
}

func (g cubegrid) Step(inclW bool) (out cubegrid) {
	out = make(cubegrid, len(g))
	for p, pActive := range g {
		if !pActive {
			continue
		}
		p.DoForNeighbors(true, inclW, func(p pos) {
			if _, done := out[p]; done {
				return
			}

			active := g[p]
			n := g.NeighborCount(inclW, p)

			if active && n != 2 && n != 3 {
				active = false
			} else if !active && n == 3 {
				active = true
			}

			out[p] = active
		})
	}
	return out
}

func (g cubegrid) ActiveCount() (a int) {
	for _, isActive := range g {
		if isActive {
			a++
		}
	}
	return a
}

func Problem17a(lines []string) {
	g := ParseFlatgrid(lines)

	const iterations = 6
	for i := 0; i < iterations; i++ {
		g = g.Step(false)
	}

	fmt.Println(g.ActiveCount())
}

func Problem17b(lines []string) {
	g := ParseFlatgrid(lines)

	const iterations = 6
	for i := 0; i < iterations; i++ {
		g = g.Step(true)
	}

	fmt.Println(g.ActiveCount())
}
