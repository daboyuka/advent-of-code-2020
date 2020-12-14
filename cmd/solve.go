package main

import (
	"bufio"
	"io"
	"os"

	"aoc2020"
)

func readlines(r io.Reader) (lines []string) {
	rbuf := bufio.NewReader(r)
	for {
		switch line, _, err := rbuf.ReadLine(); err {
		case nil:
			lines = append(lines, string(line))
		case io.EOF:
			return
		default:
			panic(err)
		}
	}
}

func readlinegroups(r io.Reader) (linegroups [][]string) {
	lines := readlines(r)

	var curGroup []string
	for _, line := range lines {
		if line == "" {
			if len(curGroup) > 0 {
				linegroups = append(linegroups, curGroup)
				curGroup = nil
			}
		} else {
			curGroup = append(curGroup, line)
		}
	}
	if len(curGroup) > 0 {
		linegroups = append(linegroups, curGroup)
	}
	return linegroups
}

func main() {
	switch os.Args[1] {
	case "1a":
		aoc2020.Problem1a(readlines(os.Stdin))
	case "1b":
		aoc2020.Problem1b(readlines(os.Stdin))
	case "2a":
		aoc2020.Problem2a(readlines(os.Stdin))
	case "2b":
		aoc2020.Problem2b(readlines(os.Stdin))
	case "3a":
		aoc2020.Problem3a(readlines(os.Stdin))
	case "3b":
		aoc2020.Problem3b(readlines(os.Stdin))
	case "4a":
		aoc2020.Problem4a(readlinegroups(os.Stdin))
	case "4b":
		aoc2020.Problem4b(readlinegroups(os.Stdin))
	case "5a":
		aoc2020.Problem5a(readlines(os.Stdin))
	case "5b":
		aoc2020.Problem5b(readlines(os.Stdin))
	case "6a":
		aoc2020.Problem6a(readlinegroups(os.Stdin))
	case "6b":
		aoc2020.Problem6b(readlinegroups(os.Stdin))
	case "7a":
		aoc2020.Problem7a(readlines(os.Stdin))
	case "7b":
		aoc2020.Problem7b(readlines(os.Stdin))
	case "8a":
		aoc2020.Problem8a(readlines(os.Stdin))
	case "8b":
		aoc2020.Problem8b(readlines(os.Stdin))
	case "9a":
		aoc2020.Problem9a(readlines(os.Stdin))
	case "9b":
		aoc2020.Problem9b(readlines(os.Stdin))
	case "10a":
		aoc2020.Problem10a(readlines(os.Stdin))
	case "10b":
		aoc2020.Problem10b(readlines(os.Stdin))
	case "11a":
		aoc2020.Problem11a(readlines(os.Stdin))
	case "11b":
		aoc2020.Problem11b(readlines(os.Stdin))
	case "12a":
		aoc2020.Problem12a(readlines(os.Stdin))
	case "12b":
		aoc2020.Problem12b(readlines(os.Stdin))
	case "13a":
		aoc2020.Problem13a(readlines(os.Stdin))
	case "13b":
		aoc2020.Problem13b(readlines(os.Stdin))
	case "14a":
		aoc2020.Problem14a(readlines(os.Stdin))
	case "14b":
		aoc2020.Problem14b(readlines(os.Stdin))
	}
}
