package aoc2020

import (
	"fmt"
	"regexp"
	"strings"

	. "aoc2020/helpers"
)

type passport map[string]string

func (p passport) Has(field string) bool { _, ok := p[field]; return ok }

var reqPassportFields = [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // cid

var (
	hclRegexp = regexp.MustCompile("^#[0-9a-f]{6}$")
	eclRegexp = regexp.MustCompile("^(?:amb|blu|brn|gry|grn|hzl|oth)$")
	pidRegexp = regexp.MustCompile("^[0-9]{9}$")
)

func (p passport) IsComplete() bool {
	for _, req := range reqPassportFields {
		if !p.Has(req) {
			return false
		}
	}
	return true
}

func (p passport) checkByr() bool {
	return Between(AtoiOrZero(p["byr"]), 1920, 2002)
}
func (p passport) checkIyr() bool {
	return Between(AtoiOrZero(p["iyr"]), 2010, 2020)
}
func (p passport) checkEyr() bool {
	return Between(AtoiOrZero(p["eyr"]), 2020, 2030)
}
func (p passport) checkHgt() bool {
	hgt := p["hgt"]
	if len(hgt) < 2 {
		return false
	}

	hgtNum, hgtSuff := AtoiOrZero(hgt[:len(hgt)-2]), hgt[len(hgt)-2:]

	switch hgtSuff {
	case "in":
		return Between(hgtNum, 59, 76)
	case "cm":
		return Between(hgtNum, 150, 193)
	default:
		return false
	}
}
func (p passport) checkHcl() bool {
	return hclRegexp.MatchString(p["hcl"])
}
func (p passport) checkEcl() bool {
	return eclRegexp.MatchString(p["ecl"])
}
func (p passport) checkPid() bool {
	return pidRegexp.MatchString(p["pid"])
}

func (p passport) IsCompleteAndValid() bool {
	return p.IsComplete() && p.checkByr() && p.checkEyr() && p.checkIyr() && p.checkHgt() && p.checkHcl() && p.checkEcl() && p.checkPid()
}

func (p passport) ParseAndAddLine(line string) {
	for _, word := range strings.Split(line, " ") {
		colonIdx := strings.Index(word, ":")
		key, value := word[:colonIdx], word[colonIdx+1:]
		p[key] = value
	}
}

func ParsePassports(linegroups [][]string) (ps []passport) {
	for _, lines := range linegroups {
		curP := make(passport)
		for _, line := range lines {
			curP.ParseAndAddLine(line)
		}
		ps = append(ps, curP)
	}
	return ps
}

func Problem4a(linegroups [][]string) {
	valid, invalid := 0, 0

	for _, p := range ParsePassports(linegroups) {
		if p.IsComplete() {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Println(valid, invalid)
}

func Problem4b(linegroups [][]string) {
	valid, invalid := 0, 0

	for _, p := range ParsePassports(linegroups) {
		if p.IsCompleteAndValid() {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Println(valid, invalid)
}
