package main

import (
	"fmt"
	"math/bits"
)

type customsForm uint

func ParseCustomsForm(line string) (c customsForm) {
	for _, char := range line {
		c = c | (1 << (char - 'a'))
	}
	return c
}

func (c customsForm) Union(other customsForm) customsForm { return c | other }

func (c customsForm) QuestionCount() int { return bits.OnesCount(uint(c)) }

func problem6a(linegroups [][]string) {
	questionCountTotal := 0

	for _, lines := range linegroups {
		var unionForm customsForm // blank form
		for _, line := range lines {
			unionForm = unionForm.Union(ParseCustomsForm(line))
		}
		questionCountTotal += unionForm.QuestionCount()
	}

	fmt.Println(questionCountTotal)
}
