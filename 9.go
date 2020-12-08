package aoc2020

type thingy struct{}

func parseThingy(line string) thingy { return thingy{} }

func parseLines(lines []string) (items []thingy) {
	for _, line := range lines {
		items = append(items, parseThingy(line))
	}
	return items
}

func Problem9a(lines []string) {
	items := parseLines(lines)
	_ = items
}

func Problem9b(lines []string) {
	items := parseLines(lines)
	_ = items
}
