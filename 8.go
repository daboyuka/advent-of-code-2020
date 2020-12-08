package aoc2020

import (
	"fmt"
	"strings"

	. "aoc2020/helpers"
)

type instruction struct {
	Kind string
	Amt  int
}

type state struct {
	Accum int
	CP    int
}

func ParseProgram(lines []string) (instrs []instruction) {
	for _, line := range lines {
		spaceIdx := strings.Index(line, " ")
		op := line[:spaceIdx]
		amt := MustAtoi(line[spaceIdx+2:])
		if line[spaceIdx+1] == '-' {
			amt *= -1
		}

		instr := instruction{Kind: op, Amt: amt}
		instrs = append(instrs, instr)
	}
	return instrs

}

func Problem8a(lines []string) {
	code := ParseProgram(lines)

	instrsRun := map[int]bool{}
	curState := state{Accum: 0, CP: 0}
	instrsRun[0] = true
	for {
		nextState := curState
		instr := code[curState.CP]
		switch instr.Kind {
		case "nop":
			nextState.CP++
		case "acc":
			nextState.Accum += instr.Amt
			nextState.CP++
		case "jmp":
			nextState.CP += instr.Amt
		}

		if instrsRun[nextState.CP] {
			fmt.Println(curState.Accum)
			return
		}
		instrsRun[nextState.CP] = true
		curState = nextState
	}
}

func runProgram(code []instruction) (int, bool) {
	instrsRun := map[int]bool{}
	curState := state{Accum: 0, CP: 0}
	instrsRun[0] = true
	for {
		nextState := curState
		instr := code[curState.CP]
		switch instr.Kind {
		case "nop":
			nextState.CP++
		case "acc":
			nextState.Accum += instr.Amt
			nextState.CP++
		case "jmp":
			nextState.CP += instr.Amt
		}

		if nextState.CP == len(code) {
			return nextState.Accum, true
		} else if instrsRun[nextState.CP] {
			return 0, false
		}
		instrsRun[nextState.CP] = true
		curState = nextState
	}
}

func Problem8b(lines []string) {
	code := ParseProgram(lines)

	for i, instr := range code {
		var accum int
		var ok bool
		if instr.Kind == "nop" {
			code[i].Kind = "jmp"
			accum, ok = runProgram(code)
			code[i].Kind = "nop"
		} else if instr.Kind == "jmp" {
			code[i].Kind = "nop"
			accum, ok = runProgram(code)
			code[i].Kind = "jmp"

		}
		if ok {
			fmt.Println(accum)
			return
		}
	}
}
