package main

import (
	"fmt"
	"strings"
)

type passport map[string]string

func (p passport) Has(field string) bool { _, ok := p[field]; return ok }

var reqPassportFields = [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // cid

func (p passport) IsComplete() bool {
	for _, req := range reqPassportFields {
		if !p.Has(req) {
			return false
		}
	}
	return true
}

func (p passport) ParseAndAddLine(line string) {
	for _, word := range strings.Split(line, " ") {
		colonIdx := strings.Index(word, ":")
		key, value := word[:colonIdx], word[colonIdx+1:]
		p[key] = value
	}
}

func ParsePassports(lines []string) (ps []passport) {
	curP := make(passport)
	for _, line := range lines {
		if line == "" {
			if len(curP) > 0 {
				ps = append(ps, curP)
				curP = make(passport)
			}
		} else {
			curP.ParseAndAddLine(line)
		}
	}
	if len(curP) > 0 {
		ps = append(ps, curP)
	}
	return ps
}

func problem4a(lines []string) {
	valid, invalid := 0, 0

	for _, p := range ParsePassports(lines) {
		if p.IsComplete() {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Println(valid, invalid)
}
