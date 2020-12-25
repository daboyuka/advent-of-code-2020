package aoc2020

import (
	"fmt"

	. "aoc2020/helpers"
)

const modulus = 20201227
const startingSubject = 7

func Transform(subject int, loopSize int) int {
	x := 1
	for i := 0; i < loopSize; i++ {
		x *= subject
		x %= modulus
	}
	return x
}

func FindLoopSize(subject int, targetPubKey int) int {
	x := 1
	for i := 0; i < modulus; i++ {
		if x == targetPubKey {
			return i
		}
		x *= subject
		x %= modulus
	}
	return -1
}

func Problem25a(lines []string) {
	pubKeys := IntLines(lines)

	firstLoopSize := FindLoopSize(startingSubject, pubKeys[0])

	fmt.Println(firstLoopSize)
	fmt.Println(Transform(pubKeys[1], firstLoopSize))
}

func Problem25b(lines []string) {

}
