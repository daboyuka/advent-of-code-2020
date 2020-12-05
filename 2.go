package main

import (
	"fmt"
	"regexp"
	"strings"
)

var lineRegexp = regexp.MustCompile(`(\d+)-(\d+) (.): (.+)`)

func problem2a(lines []string) {
	valid, invalid := 0, 0
	for _, line := range lines {
		fields := lineRegexp.FindStringSubmatch(line)
		if fields == nil {
			panic(fmt.Errorf("wut %s", line))
		}

		from, to, char, password := mustAtoi(string(fields[1])), mustAtoi(fields[2]), fields[3], fields[4]

		reps := strings.Count(password, char)
		if reps >= from && reps <= to {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Println(valid, invalid)
}

func problem2b(lines []string) {
	valid, invalid := 0, 0
	for _, line := range lines {
		fields := lineRegexp.FindStringSubmatch(line)
		if fields == nil {
			panic(fmt.Errorf("wut %s", line))
		}

		at1, at2, char, password := mustAtoi(string(fields[1])), mustAtoi(fields[2]), fields[3][0], fields[4]

		matches := 0
		for _, pos := range [...]int{at1 - 1, at2 - 1} {
			if pos >= 0 && pos < len(password) && password[pos] == char {
				matches++
			}
		}

		if matches == 1 {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Println(valid, invalid)
}