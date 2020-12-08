package aoc2020

import (
	"fmt"
	"strings"

	. "aoc2020/helpers"
)

type bagType string

type bagTypeWithCount struct {
	Type  bagType
	Count int
}

type bagRule struct {
	Type     bagType
	Contains map[bagType]int
}

type bagRulesMap map[bagType]map[bagType]int

func ParseBagRule(line string) (rule bagRule) {
	line = line[:len(line)-1] // remove period
	containsIdx := strings.Index(line, " contain ")
	fromType := line[:containsIdx-1] // subtract the 's' in bags
	containsList := strings.Split(line[containsIdx+len(" contain "):], ", ")

	rule.Type = bagType(fromType)
	rule.Contains = make(map[bagType]int)
	for _, contains := range containsList {
		firstSpaceIdx := strings.Index(contains, " ")
		numStr := contains[:firstSpaceIdx]
		if numStr == "no" {
			break
		}

		num := MustAtoi(numStr)

		containsType := contains[firstSpaceIdx+1:]
		if num > 1 {
			containsType = containsType[:len(containsType)-1] // remove s
		}

		rule.Contains[bagType(containsType)] = num
	}

	return rule
}

func Problem7a(lines []string) {
	bmap := bagRulesMap{}
	for _, line := range lines {
		rule := ParseBagRule(line)
		bmap[rule.Type] = rule.Contains
	}

	const targetType = "shiny gold bag"

	bagTypeContainsTarget := map[bagType]bool{}
	bagTypeContainsTarget[targetType] = true // trivially ; remove from count later

	var dft func(bagType) bool // return true if contains target type
	dft = func(bt bagType) bool {
		if doesContain, ok := bagTypeContainsTarget[bt]; ok {
			return doesContain
		}

		doesContain := false
		for containsType := range bmap[bt] {
			doesContain = doesContain || dft(containsType)
		}

		bagTypeContainsTarget[bt] = doesContain
		return doesContain
	}

	for bt := range bmap {
		dft(bt)
	}

	containsCount := 0
	for bt, doesContain := range bagTypeContainsTarget {
		fmt.Println(bt, doesContain)
		if doesContain && bt != targetType {
			containsCount++
		}
	}
	fmt.Println(containsCount)
}

func Problem7b(lines []string) {
	bmap := bagRulesMap{}
	for _, line := range lines {
		rule := ParseBagRule(line)
		bmap[rule.Type] = rule.Contains
	}

	const targetType = "shiny gold bag"

	var dft func(bagType) int // return num bags, including self
	dft = func(bt bagType) int {
		numBags := 1 // self
		for containsType, containsCount := range bmap[bt] {
			numBags += containsCount * dft(containsType)
		}
		fmt.Println(bt, "contains", numBags)
		return numBags
	}

	fmt.Println(dft(targetType) - 1)
}
