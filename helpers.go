package main

import "strconv"

func mustAtoi(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func atoiOrZero(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return x
}

func between(x, low, high int) bool {
	return x >= low && x <= high
}
