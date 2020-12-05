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
		problem4a(readlines(os.Stdin))
	}
}
