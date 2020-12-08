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
