package aoc2020

import (
	"fmt"
	"strings"

	. "aoc2020/helpers"
)

type tile struct {
	Size      int
	CornerIds [4]int // NE NW SW SE

	Sides    [4]tileEdge // E N W S; string order is ccwise
	RevSides [4]tileEdge //

	Inner [4][]tileEdge // string order points E, N, W, S; slice order 90deg right of string order
}

func revLines(ssIn []tileEdge) []tileEdge {
	ss := append([]tileEdge(nil), ssIn...)
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[len(ss)-1-i] = ss[len(ss)-1-i], ss[i]
	}
	return ss
}
func revStrs(ssIn []tileEdge) []tileEdge {
	ss := append([]tileEdge(nil), ssIn...)
	for i, s := range ss {
		ss[i] = s.Rev()
	}
	return ss
}
func fuseLines(a, b []tileEdge, inbetween ...tileEdge) []tileEdge {
	c := append([]tileEdge(nil), a...)
	c = append(c, inbetween...)
	c = append(c, b...)
	return c
}
func fuseStrs(a, b []tileEdge, inbetween ...tileEdge) []tileEdge {
	c := make([]tileEdge, len(a))
	for i := range a {
		c[i] = a[i]
		for _, x := range inbetween {
			c[i] += tileEdge(x[i])
		}
		c[i] += b[i]
	}
	return c
}
func dropEnds(x tileEdge) tileEdge { return x[1 : len(x)-1] }

type tileEdge string

func (te tileEdge) Rev() tileEdge {
	b := strings.Builder{}
	for i := len(te) - 1; i >= 0; i-- {
		b.WriteByte(te[i])
	}
	return tileEdge(b.String())
}

func (te tileEdge) Match(mask tileEdge, offset int) bool {
	if offset+len(mask) > len(te) {
		return false
	}
	for i, c := range mask {
		if c != '#' {
			continue
		} else if te[offset+i] != '#' {
			return false
		}
	}
	return true
}

func (t tile) String() string {
	width := len(t.Sides[1])
	b := strings.Builder{}
	neStr, nwStr, swStr, seStr := fmt.Sprintf("%d", t.CornerIds[0]), fmt.Sprintf("%d", t.CornerIds[1]), fmt.Sprintf("%d", t.CornerIds[2]), fmt.Sprintf("%d", t.CornerIds[3])
	b.WriteString(nwStr)
	b.WriteString(strings.Repeat(" ", width-len(nwStr)-len(neStr)))
	b.WriteString(neStr)
	b.WriteByte('\n')

	inset := (len(t.Sides[0]) - len(t.Inner[0])) / 2
	centerStr := strings.Repeat(" ", len(t.Sides[0])-2)
	insetStr := strings.Repeat(" ", inset-1)

	b.WriteString(string(t.RevSides[1]))
	b.WriteByte('\n')
	for i, cw := range t.Sides[2] {
		if i == 0 || i == len(t.Sides[2])-1 {
			continue
		}
		ce := t.RevSides[0][i]
		b.WriteRune(cw)
		if i >= inset && i < len(t.Sides[2])-inset {
			b.WriteString(insetStr + string(t.Inner[0][i-inset]) + insetStr)
		} else {
			b.WriteString(centerStr)
		}
		b.WriteByte(ce)
		b.WriteByte('\n')
	}
	b.WriteString(string(t.Sides[3]))
	b.WriteByte('\n')

	b.WriteString(swStr)
	b.WriteString(strings.Repeat(" ", width-len(swStr)-len(seStr)))
	b.WriteString(seStr)
	b.WriteByte('\n')

	return b.String()
}

func ParseTile(lines []string) (t tile) {
	idLine := lines[0]
	id := Atoi(idLine[len("Tile ") : len(idLine)-len(":")])
	lines = lines[1:]

	nEdge, sEdge := tileEdge(lines[0]).Rev(), tileEdge(lines[len(lines)-1])
	wEdge, eEdge := tileEdge(""), tileEdge("")
	eInner := make([]tileEdge, len(lines)-2)
	sInner := make([]tileEdge, len(lines)-2)
	for lineIdx, line := range lines {
		wEdge += tileEdge(line[0])
		eEdge += tileEdge(line[len(line)-1])

		if lineIdx != 0 && lineIdx != len(lines)-1 {
			eInner[lineIdx-1] = dropEnds(tileEdge(line))
			for col, c := range line[1 : len(line)-1] {
				sInner[col] += tileEdge(c)
			}
		}
	}
	eEdge = eEdge.Rev()
	sInner = revLines(sInner)
	nInner, wInner := revStrs(revLines(sInner)), revStrs(revLines(eInner))

	return tile{
		Size:      1,
		CornerIds: [4]int{id, id, id, id},
		Sides:     [4]tileEdge{eEdge, nEdge, wEdge, sEdge},
		RevSides:  [4]tileEdge{eEdge.Rev(), nEdge.Rev(), wEdge.Rev(), sEdge.Rev()},
		Inner:     [4][]tileEdge{eInner, nInner, wInner, sInner},
	}
}

func ParseTiles(linegroups [][]string) (tiles []tile) {
	for _, lines := range linegroups {
		tiles = append(tiles, ParseTile(lines))
	}
	return tiles
}

func (t tile) RotateLeft() tile {
	return tile{
		Size:      t.Size,
		CornerIds: [4]int{t.CornerIds[3], t.CornerIds[0], t.CornerIds[1], t.CornerIds[2]},
		Sides:     [4]tileEdge{t.Sides[3], t.Sides[0], t.Sides[1], t.Sides[2]},
		RevSides:  [4]tileEdge{t.RevSides[3], t.RevSides[0], t.RevSides[1], t.RevSides[2]},
		Inner:     [4][]tileEdge{t.Inner[3], t.Inner[0], t.Inner[1], t.Inner[2]},
	}
}

func (t tile) FlipEW() tile {
	return tile{
		Size:      t.Size,
		CornerIds: [4]int{t.CornerIds[1], t.CornerIds[0], t.CornerIds[3], t.CornerIds[2]},
		Sides:     [4]tileEdge{t.RevSides[2], t.RevSides[1], t.RevSides[0], t.RevSides[3]},
		RevSides:  [4]tileEdge{t.Sides[2], t.Sides[1], t.Sides[0], t.Sides[3]},
		Inner:     [4][]tileEdge{revStrs(t.Inner[0]), revLines(t.Inner[1]), revStrs(t.Inner[2]), revLines(t.Inner[3])},
	}
}

func (t tile) Match(other tile) (sideIdx, otherSideIdx int, rev bool, ok bool) {
	for sideIdx, side := range t.Sides {
		for otherSideIdx, otherSide := range other.Sides {
			if side == otherSide {
				return sideIdx, otherSideIdx, false, true
			}
		}
		for otherSideIdx, otherSide := range other.RevSides {
			if side == otherSide {
				return sideIdx, otherSideIdx, true, true
			}
		}
	}
	return 0, 0, false, false
}

func (t tile) Join(other tile) (tile, bool) {
	sideIdx, otherSideIdx, rev, ok := t.Match(other)
	if !ok {
		return tile{}, false
	}

	for sideIdx != 1 {
		t = t.RotateLeft()
		sideIdx = (sideIdx + 1) % 4
	}
	for otherSideIdx != 3 {
		other = other.RotateLeft()
		otherSideIdx = (otherSideIdx + 1) % 4
	}
	if !rev {
		other = other.FlipEW()
	}

	if t.Sides[1] != other.RevSides[3] {
		panic(fmt.Errorf("%d %d %t -> %+v %+v", sideIdx, otherSideIdx, rev, t.Sides[1], other.RevSides[3]))
	}

	// t goes to the south of other; t's north side matches other's south side
	return tile{
		Size:      t.Size + other.Size,
		CornerIds: [4]int{other.CornerIds[0], other.CornerIds[1], t.CornerIds[2], t.CornerIds[3]},
		Sides:     [4]tileEdge{t.Sides[0] + other.Sides[0], other.Sides[1], other.Sides[2] + t.Sides[2], t.Sides[3]},
		RevSides:  [4]tileEdge{other.RevSides[0] + t.RevSides[0], other.RevSides[1], t.RevSides[2] + other.RevSides[2], t.RevSides[3]},
		Inner: [4][]tileEdge{
			fuseLines(other.Inner[0], t.Inner[0]), //, dropEnds(other.Sides[3]), dropEnds(t.RevSides[1])),
			fuseStrs(t.Inner[1], other.Inner[1]),  // dropEnds(t.RevSides[1]), dropEnds(other.Sides[3])),
			fuseLines(t.Inner[2], other.Inner[2]), // dropEnds(t.Sides[1]), dropEnds(other.RevSides[3])),
			fuseStrs(other.Inner[3], t.Inner[3]),  // dropEnds(other.RevSides[3]), dropEnds(t.Sides[1]))},
		},
	}, true
}

func Problem20a(linegroups [][]string) {
	tiles := ParseTiles(linegroups)

	allTiles := map[*tile]bool{}
	for _, t := range tiles {
		tCopy := t
		allTiles[&tCopy] = true
	}

	for len(allTiles) > 1 {
		fmt.Println("ITER", len(allTiles))
		for t := range allTiles {
			for {
				found := false
				for other := range allTiles {
					if t == other {
						continue
					}

					if joined, ok := t.Join(*other); ok {
						//fmt.Println(t.CornerIds[0], other.CornerIds[0])
						delete(allTiles, t)
						delete(allTiles, other)
						allTiles[&joined] = true
						t = &joined
						found = true
					}
				}
				if !found {
					break
				}
			}
		}
	}

	var last *tile
	for last = range allTiles {
	}
	fmt.Println(last.String(), last.CornerIds[0]*last.CornerIds[1]*last.CornerIds[2]*last.CornerIds[3])
}

func MatchAt(lines []tileEdge, matchMask []tileEdge, row, col int) bool {
	if row+len(matchMask) > len(lines) {
		return false
	}
	for mRow, mLine := range matchMask {
		if !lines[row+mRow].Match(mLine, col) {
			return false
		}
	}
	return true
}

func CountLineHashes(lines []tileEdge) (n int) {
	for _, line := range lines {
		n += strings.Count(string(line), "#")
	}
	return n
}

func MaskToCoords(row, col int, matchMask []tileEdge, m map[[2]int]bool) {
	for mRow, mask := range matchMask {
		for mCol, c := range mask {
			if c == '#' {
				m[[2]int{row + mRow, col + mCol}] = true
			}
		}
	}
}

func Problem20b(linegroups [][]string) {
	tiles := ParseTiles(linegroups)

	allTiles := map[*tile]bool{}
	for _, t := range tiles {
		tCopy := t
		allTiles[&tCopy] = true
	}

	for len(allTiles) > 1 {
		fmt.Println("ITER", len(allTiles))
		for t := range allTiles {
			for {
				found := false
				for other := range allTiles {
					if t == other {
						continue
					}

					if joined, ok := t.Join(*other); ok {
						//fmt.Println(t.CornerIds[0], other.CornerIds[0])
						delete(allTiles, t)
						delete(allTiles, other)
						allTiles[&joined] = true
						t = &joined
						found = true
					}
				}
				if !found {
					break
				}
			}
		}
	}

	var last *tile
	for last = range allTiles {
	}

	last2 := last.FlipEW().RotateLeft() // .FlipEW()
	last = &last2
	fmt.Println(last)

	var monster = [3]tileEdge{
		"                   #",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}
	var monsterInv = [3]tileEdge{monster[2], monster[1], monster[0]}

	monsters := 0
	monsterPixels := map[[2]int]bool{}
	for dir, lines := range last.Inner {
		for rowIdx := 0; rowIdx < len(lines); rowIdx++ {
			for colIdx := 0; colIdx < len(lines[rowIdx]); colIdx++ {
				if MatchAt(lines, monster[:], rowIdx, colIdx) {
					fmt.Println("monster at", dir, rowIdx, colIdx)
					monsters++
					MaskToCoords(rowIdx, colIdx, monster[:], monsterPixels)
				}
				if MatchAt(lines, monsterInv[:], rowIdx, colIdx) {
					fmt.Println("inv monster at", dir, rowIdx, colIdx)
					monsters++
					MaskToCoords(rowIdx, colIdx, monsterInv[:], monsterPixels)
				}
			}
		}
	}

	fmt.Println(monsters, CountLineHashes(last.Inner[0]), CountLineHashes(monster[:]),
		CountLineHashes(last.Inner[0])-monsters*CountLineHashes(monster[:]),
		CountLineHashes(last.Inner[0])-len(monsterPixels))
}
