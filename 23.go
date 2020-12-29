package aoc2020

import (
	"fmt"
	"strings"

	. "aoc2020/helpers"
)

type cup struct {
	idZB int
	next *cup
}

type game struct {
	idToCup []*cup
	curCup  *cup
}

func (g *game) RemoveN(after *cup, n int) (firstRemoved, lastRemoved *cup) {
	firstRemoved, lastRemoved = after.next, after
	for i := 0; i < n; i++ {
		lastRemoved = lastRemoved.next
	}

	after.next, lastRemoved.next = lastRemoved.next, nil
	return firstRemoved, lastRemoved
}

// first/last inclusive
func (g *game) Insert(after, first, last *cup) {
	after.next, last.next = first, after.next
}

// first/last inclusive
func (c *cup) InRange(first, last *cup) bool {
	for it := first; ; it = it.next {
		if it == c {
			return true
		} else if it == last {
			return false
		}
	}
}

func (g *game) Move() {
	// Step 1: remove next few cups
	const removeN = 3
	firstRemoved, lastRemoved := g.RemoveN(g.curCup, removeN)

	// Step 2: find destination cup
	var destC *cup

	decId := func(id int) int { return (id + len(g.idToCup) - 1) % len(g.idToCup) }
	for destId := decId(g.curCup.idZB); ; destId = decId(destId) {
		destC = g.idToCup[destId]
		if !destC.InRange(firstRemoved, lastRemoved) {
			break
		}
	}

	// Step 3: reinsert cups after dest cup
	g.Insert(destC, firstRemoved, lastRemoved)

	// Step 4: new current cup is next clockwise
	g.curCup = g.curCup.next
}

func Problem23a(lines []string) {
	const nCups = 9

	var idToCup [nCups]*cup
	var firstCup *cup

	var prevC *cup
	for _, char := range lines[0] {
		newC := &cup{idZB: Atoi(string(char)) - 1}

		idToCup[newC.idZB] = newC
		if firstCup == nil {
			firstCup = newC
		} else {
			prevC.next = newC
		}

		prevC = newC
	}
	prevC.next = firstCup // complete the loop

	g := game{idToCup: idToCup[:], curCup: firstCup}

	const nRounds = 100
	for i := 0; i < nRounds; i++ {
		g.Move()
	}

	buf := strings.Builder{}
	cup1 := g.idToCup[0]
	for c := cup1.next; c != cup1; c = c.next {
		buf.WriteRune('1' + rune(c.idZB))
	}
	fmt.Println(buf.String())
}

func Problem23b(lines []string) {
	const nCups = 1e6
	g := game{idToCup: make([]*cup, nCups)}

	var maxIdZB int
	var prevC *cup
	for _, char := range lines[0] {
		newC := &cup{idZB: Atoi(string(char)) - 1}
		if maxIdZB < newC.idZB {
			maxIdZB = newC.idZB
		}

		g.idToCup[newC.idZB] = newC
		if g.curCup == nil {
			g.curCup = newC
		} else {
			prevC.next = newC
		}
		prevC = newC
	}

	for idZB := maxIdZB + 1; idZB < nCups; idZB++ {
		newC := &cup{idZB: idZB}

		g.idToCup[newC.idZB] = newC
		prevC.next = newC
		prevC = newC
	}

	prevC.next = g.curCup // complete the loop

	const nRounds = 1e7
	for i := 0; i < nRounds; i++ {
		g.Move()
	}

	cup1 := g.idToCup[0]
	fmt.Println(cup1.next.idZB+1, cup1.next.next.idZB+1, (cup1.next.idZB+1)*(cup1.next.next.idZB+1))
}
