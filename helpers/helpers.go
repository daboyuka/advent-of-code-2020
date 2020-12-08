package helpers

import (
	"strconv"
	"strings"
)

func MustAtoi(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func AtoiOrZero(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return x
}

func Between(x, low, high int) bool {
	return x >= low && x <= high
}

func Split2(s, sep string) (a, b string) {
	idx := strings.Index(s, sep)
	if idx == -1 {
		return s, ""
	} else {
		return s[:idx], s[idx+len(sep):]
	}
}

func Split3(s, sep1, sep2 string) (a, b, c string) {
	if idx1 := strings.Index(s, sep1); idx1 != -1 {
		a = s[:idx1]
		b, c = Split2(s[idx1+len(sep1):], sep2)
	} else {
		b = ""
		a, c = Split2(s, sep2)
	}
	return a, b, c
}
