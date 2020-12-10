package aoc2020

import (
	"fmt"
	"sort"

	. "aoc2020/helpers"
)

func Problem10a(lines []string) {
	jolts := IntLines(lines)
	maxJ := Max(jolts...)
	jolts = append(jolts, 0, maxJ+3)

	num1, num3 := 0, 0
	sort.Ints(jolts)
	for i, next := range jolts[1:] {
		prev := jolts[i]
		diff := next - prev
		if diff == 1 {
			num1++
		} else if diff == 3 {
			num3++
		}
	}

	fmt.Println(num1, num3, num1*num3)
}

func Problem10b(lines []string) {
	jolts := IntLines(lines)
	maxJ := Max(jolts...)
	jolts = append(jolts, 0, maxJ+3)
	sort.Ints(jolts)

	comboCounts := make([]int, maxJ+3+1)
	comboCounts[0] = 1
	for _, jolt := range jolts[1:] {
		comboCounts[jolt] = Sum(comboCounts[Max(jolt-3, 0):jolt]...)
	}

	fmt.Println(comboCounts[maxJ+3])
}
