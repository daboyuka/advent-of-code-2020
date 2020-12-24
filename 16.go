package aoc2020

import (
	"fmt"
	"strings"

	. "aoc2020/helpers"
)

type fieldRule struct {
	Name        string
	ValidRanges [][2]int
}

func (r fieldRule) Valid(v int) bool {
	for _, b := range r.ValidRanges {
		if v >= b[0] && v <= b[1] {
			return true
		}
	}
	return false
}

type ticket []int

type notes struct {
	rules  []fieldRule
	mine   ticket
	nearby []ticket
}

func ParseNotes(lines []string) (note notes) {
	i := 0
	for ; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}

		name, allRulesStr := Split2(line, ": ")
		rule := fieldRule{Name: name}

		ruleStrs := strings.Split(allRulesStr, " or ")
		for _, ruleStr := range ruleStrs {
			lb, ub := Split2(ruleStr, "-")
			rule.ValidRanges = append(rule.ValidRanges, [2]int{Atoi(lb), Atoi(ub)})
		}

		note.rules = append(note.rules, rule)
	}
	i++

	if lines[i] != "your ticket:" {
		panic(lines[i])
	}
	i++

	note.mine = IntLines(strings.Split(lines[i], ","))
	i++
	i++

	if lines[i] != "nearby tickets:" {
		panic(lines[i])
	}
	i++

	for _, line := range lines[i:] {
		note.nearby = append(note.nearby, IntLines(strings.Split(line, ",")))
	}

	return note
}

func Problem16a(lines []string) {
	notes := ParseNotes(lines)
	fmt.Printf("%+v\n", notes)

	badValSum := 0
	for _, t := range notes.nearby {
	vloop:
		for _, v := range t {
			for _, r := range notes.rules {
				if r.Valid(v) {
					fmt.Println(v, "valid in", r)
					continue vloop
				}
			}
			badValSum += v
		}
	}

	fmt.Println(badValSum)
}

func Problem16b(lines []string) {
	notes := ParseNotes(lines)

	type fieldpos struct {
		Name string
		Pos  int
	}

	valid := map[fieldpos]bool{}
	for _, r := range notes.rules {
		for pos := range notes.mine {
			valid[fieldpos{r.Name, pos}] = true
		}
	}

	skipTickets := map[int]bool{}
tloop:
	for ticketIdx, t := range notes.nearby {
	vloop:
		for _, v := range t {
			for _, r := range notes.rules {
				if r.Valid(v) {
					continue vloop
				}
			}
			skipTickets[ticketIdx] = true
			continue tloop
		}
	}

	fieldPosPossible := make(map[string][]int)
	for _, r := range notes.rules {
	posloop:
		for pos := range notes.mine {
			for ticketIdx, t := range notes.nearby {
				if skipTickets[ticketIdx] {
					continue
				}
				if !r.Valid(t[pos]) {
					continue posloop
				}
			}

			// All tickets valid at this pos
			fieldPosPossible[r.Name] = append(fieldPosPossible[r.Name], pos)
		}
	}

	fmt.Println(fieldPosPossible)

	fieldPos := make(map[string]int)
	elimPos := map[int]bool{}
	for {
		if len(fieldPosPossible) == 0 {
			break
		}
	fieldNameLoop:
		for fieldName, poses := range fieldPosPossible {
			possible := -1
			for _, pos := range poses {
				if elimPos[pos] {
					continue
				} else if possible == -1 {
					possible = pos
				} else {
					continue fieldNameLoop
				}
			}
			fieldPos[fieldName] = possible
			elimPos[possible] = true
			delete(fieldPosPossible, fieldName)
		}
	}

	fmt.Println(fieldPos)

	prod := 1
	for fieldName, pos := range fieldPos {
		if strings.HasPrefix(fieldName, "departure") {
			fmt.Println(fieldName, pos)
			prod *= notes.mine[pos]
		}
	}
	fmt.Println(prod)
}
