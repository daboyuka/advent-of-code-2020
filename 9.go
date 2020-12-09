package aoc2020

import (
	"fmt"

	"aoc2020/helpers"
)

func parseLines(lines []string) (nums []int) {
	for _, line := range lines {
		nums = append(nums, helpers.MustAtoi(line))
	}
	return nums
}

func check(prefix []int, target int) bool {
	for _, a := range prefix {
		for _, b := range prefix {
			if a != b && a+b == target {
				return true
			}
		}
	}
	return false
}

func Problem9a(lines []string) {
	nums := parseLines(lines)

	for i := 25; i < len(nums); i++ {
		prefix := nums[i-25 : i]
		if !check(prefix, nums[i]) {
			fmt.Println(nums[i])
			return
		}
	}
}

func findRange(nums []int, target int) (first, last int) {
	accums := make([]int, len(nums))
	for i, x := range nums {
		accums[i] = x
		if i > 0 {
			accums[i] += accums[i-1]
		}
	}

	for i := range accums {
		for j := range accums[i+1:] {
			if accums[j+i+1]-accums[i] == target {
				return i, j + i + 1
			}
		}
	}

	return -1, -1
}

func Problem9b(lines []string) {
	nums := parseLines(lines)

	target := -1
	for i := 25; i < len(nums); i++ {
		prefix := nums[i-25 : i]
		if !check(prefix, nums[i]) {
			target = nums[i]
			break
		}
	}

	first, last := findRange(nums, target)
	min, max := -1, -1
	for _, x := range nums[first:last] {
		if min == -1 || min > x {
			min = x
		}
		if max == -1 || max < x {
			max = x
		}
	}

	fmt.Println(first, last, min, max)
}
