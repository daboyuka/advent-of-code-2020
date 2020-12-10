package aoc2020

import (
	"fmt"
	"sort"

	. "aoc2020/helpers"
)

func ParseThingies(lines []string) (out []int) {
	for _, line := range lines {
		out = append(out, MustAtoi(line))
	}
	return out
}

func Problem10a(lines []string) {
	jolts := ParseThingies(lines)
	maxJ := 0
	for _, x := range jolts {
		if maxJ < x {
			maxJ = x
		}
	}
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
	jolts := ParseThingies(lines)
	maxJ := 0
	for _, x := range jolts {
		if maxJ < x {
			maxJ = x
		}
	}
	jolts = append(jolts, 0, maxJ+3)

	sort.Ints(jolts)

	comboCounts := make([]uint64, maxJ+3+1)
	comboCounts[0] = 1
	for _, jolt := range jolts {
		if jolt >= 1 {
			comboCounts[jolt] += comboCounts[jolt-1]
		}
		if jolt >= 2 {
			comboCounts[jolt] += comboCounts[jolt-2]
		}
		if jolt >= 3 {
			comboCounts[jolt] += comboCounts[jolt-3]
		}
	}

	fmt.Println(comboCounts[maxJ+3])
}
