package aoc2020

import (
	"fmt"
	"strings"

	. "aoc2020/helpers"
)

func doProblem15(targetPos int, initial []int) {
	lastPos := map[int]int{} // val -> last pos (0-based)
	curNum := 0              // last spoken number
	curAge := 0              // age of last spoken number
	for i, x := range initial {
		if p, ok := lastPos[x]; ok {
			curAge = i - p
		} else {
			curAge = 0
		}
		lastPos[x] = i
		curNum = x
	}

	for i := len(initial); i < targetPos; i++ {
		curNum = curAge
		if p, ok := lastPos[curNum]; ok {
			curAge = i - p
		} else {
			curAge = 0
		}
		lastPos[curNum] = i
	}

	fmt.Println(curNum)
}

func Problem15a(lines []string) {
	doProblem15(2020, IntLines(strings.Split(lines[0], ",")))
}

func Problem15b(lines []string) {
	doProblem15(30000000, IntLines(strings.Split(lines[0], ",")))
}
