package aoc2020

import (
	"fmt"
	"strings"

	. "aoc2020/helpers"
)

type deck []int

func (d *deck) Pop() (c int) {
	c, *d = (*d)[0], (*d)[1:]
	return
}

func (d *deck) Push(cs ...int) {
	*d = append(*d, cs...)
}

func (d deck) Copy() deck {
	return append(deck(nil), d...)
}

func (d deck) Score() (s int) {
	for i, c := range d {
		s += (len(d) - i) * c
	}
	return s
}

func (d deck) Stamp() string {
	b := strings.Builder{}
	for i, c := range d {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(fmt.Sprintf("%d", c))
	}
	return b.String()
}

func doCombatRound(p1, p2 *deck) (done bool) {
	c1, c2 := p1.Pop(), p2.Pop()
	if c1 > c2 {
		p1.Push(c1, c2)
		return len(*p2) == 0
	} else {
		p2.Push(c2, c1)
		return len(*p1) == 0
	}
}

func runCombat(p1, p2 deck) (p1Wins bool, winningDeck deck) {
	for !doCombatRound(&p1, &p2) {
	}
	if len(p1) != 0 {
		return true, p1
	} else {
		return false, p2
	}
}

func ParseDeck(lines []string) (d deck) {
	return deck(IntLines(lines[1:])) // skip player #
}

func ParseDecks(linegroups [][]string) (p1, p2 deck) {
	return ParseDeck(linegroups[0]), ParseDeck(linegroups[1])
}

func Problem22a(linegroups [][]string) {
	p1, p2 := ParseDecks(linegroups)
	_, winningDeck := runCombat(p1, p2)

	fmt.Println(winningDeck.Score())
}

func doRCombat(p1, p2 deck, prevStates map[string]bool, subgameOutcomes map[string]bool) (p1Wins bool, winningDeck deck) {
	stamp := p1.Stamp() + "|" + p2.Stamp()
	//fmt.Println(stamp)
	if prevStates[stamp] {
		fmt.Println("repeat")
		return true, p1
	}
	prevStates[stamp] = true

	defer func() { subgameOutcomes[stamp] = p1Wins }()

	c1, c2 := p1.Pop(), p2.Pop()

	p1WinsRound := false
	if c1 <= len(p1) && c2 <= len(p2) {
		//fmt.Println("recurse", c1, c2, len(p1), len(p2))
		substamp := p1.Stamp() + "|" + p2.Stamp()
		if p1WinsSubgame, ok := subgameOutcomes[substamp]; ok {
			//fmt.Println("skipped known", substamp)
			p1WinsRound = p1WinsSubgame
		} else {
			p1WinsRound, _ = doRCombat(p1[:c1].Copy(), p2[:c2].Copy(), make(map[string]bool), subgameOutcomes)
		}
	} else {
		//fmt.Println("normal", c1, c2, len(p1), len(p2))
		p1WinsRound = c1 > c2
	}

	if p1WinsRound {
		p1.Push(c1, c2)
	} else {
		p2.Push(c2, c1)
	}

	if len(p2) == 0 {
		return true, p1
	} else if len(p1) == 0 {
		return false, p2
	} else {
		return doRCombat(p1, p2, prevStates, subgameOutcomes)
	}
}

func Problem22b(linegroups [][]string) {
	p1, p2 := ParseDecks(linegroups)

	_, winningDeck := doRCombat(p1, p2, make(map[string]bool), make(map[string]bool))

	fmt.Println(winningDeck.Score())
}
