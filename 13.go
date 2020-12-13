package aoc2020

type thingy struct {
	S string
}

func ParseInput(lines []string) (out []thingy) {
	for _, line := range lines {
		out = append(out, thingy{line})
	}
	return out
}

func Problem13a(lines []string) {
	in := ParseInput(lines)
}

func Problem13b(lines []string) {
}
