package main

import (
	"bufio"
	"io"
	"os"
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
		problem1a(readlines(os.Stdin))
	case "1b":
		problem1b(readlines(os.Stdin))
	case "2a":
		problem2a(readlines(os.Stdin))
	case "2b":
		problem2b(readlines(os.Stdin))
	case "3a":
		problem3a(readlines(os.Stdin))
	case "3b":
		problem3b(readlines(os.Stdin))
	case "4a":
		problem4a(readlinegroups(os.Stdin))
	case "4b":
		problem4b(readlinegroups(os.Stdin))
	case "5a":
		problem5a(readlines(os.Stdin))
	case "5b":
		problem5b(readlines(os.Stdin))
	case "6a":
		problem6a(readlinegroups(os.Stdin))
	case "6b":
		problem6b(readlinegroups(os.Stdin))
	case "7a":
		problem7a(readlines(os.Stdin))
	case "7b":
		problem7b(readlines(os.Stdin))
	case "8a":
		problem8a(readlines(os.Stdin))
	case "8b":
		problem8b(readlines(os.Stdin))
	}
}
