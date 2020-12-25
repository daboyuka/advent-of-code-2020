package aoc2020

import (
	//"fmt"

	"fmt"
	"strings"

	. "aoc2020/helpers"
)

type rawRule struct {
	simple  string
	options [][]int // [option][subrules]
}

type rule struct {
	id      int
	simple  string
	options [][]*rule
}

func ParseRulesAndMessages(linegroups [][]string) (rules map[int]rawRule, messages []string) {
	ruleLines, messages := linegroups[0], linegroups[1]
	rules = make(map[int]rawRule)
	for _, ruleLine := range ruleLines {
		idStr, def := Split2(ruleLine, ": ")

		var r rawRule
		if def[0] == '"' {
			r.simple = def[1 : len(def)-1]
		} else {
			for _, opt := range strings.Split(def, " | ") {
				r.options = append(r.options, IntLines(strings.Split(opt, " ")))
			}
		}

		rules[Atoi(idStr)] = r
	}
	return rules, messages
}

func FuseRules(inRules map[int]rawRule) (outRules []*rule) {
	maxRuleId := 0
	for ruleId := range inRules {
		if maxRuleId < ruleId {
			maxRuleId = ruleId
		}
	}

	outRules = make([]*rule, maxRuleId+1)
	for i := range outRules {
		outRules[i] = &rule{id: i}
	}

	for i, r := range outRules {
		rr := inRules[i]
		r.simple = rr.simple
		for _, opt := range rr.options {
			optRules := make([]*rule, 0, len(opt))
			for _, optIdx := range opt {
				optRules = append(optRules, outRules[optIdx])
			}
			r.options = append(r.options, optRules)
		}
	}

	return outRules
}

func (r *rule) matchPrefix(message string, outSkips map[int]bool, outSkipBasis int, indent string) bool {
	if r.simple != "" {
		if strings.HasPrefix(message, r.simple) {
			outSkips[outSkipBasis+len(r.simple)] = true
			return true
		} else {
			return false
		}
	}

	any := false
optionLoop:
	for _, opt := range r.options {
		optSkips := map[int]bool{0: true}
		for _, optRule := range opt {
			nextSkips := make(map[int]bool)
			for skip := range optSkips {
				optRule.matchPrefix(message[skip:], nextSkips, skip, indent+" ")
			}
			if len(nextSkips) == 0 {
				continue optionLoop
			} else {
				optSkips = nextSkips
			}
		}

		any = true
		for skip := range optSkips {
			outSkips[outSkipBasis+skip] = true
		}
	}

	return any
}

func MatchWithRules(rules []*rule, ruleIdx int, message string) bool {
	outSkips := make(map[int]bool)
	anyMatch := rules[0].matchPrefix(message, outSkips, 0, "")
	return anyMatch && outSkips[len(message)]
}

func Problem19a(linegroups [][]string) {
	rawRules, messages := ParseRulesAndMessages(linegroups)
	rules := FuseRules(rawRules)

	nMatches := 0
	for _, message := range messages {
		matches := MatchWithRules(rules, 0, message)
		fmt.Println(message, matches)
		if matches {
			nMatches++
		}
	}

	fmt.Println(nMatches)
}

func Problem19b(linegroups [][]string) {
	rawRules, messages := ParseRulesAndMessages(linegroups)

	//8: 42 | 42 8
	//11: 42 31 | 42 11 31
	rawRules[8] = rawRule{options: [][]int{{42}, {42, 8}}}
	rawRules[11] = rawRule{options: [][]int{{42, 31}, {42, 11, 31}}}

	rules := FuseRules(rawRules)

	nMatches := 0
	for _, message := range messages {
		matches := MatchWithRules(rules, 0, message)
		fmt.Println(message, matches)
		if matches {
			nMatches++
		}
	}

	fmt.Println(nMatches)
}
