package main

import "fmt"

func problem1a(lines []string) {
	vals := make([]int, len(lines))
	for i, line := range lines {
		vals[i] = mustAtoi(line)
	}

	for _, a := range vals {
		for _, b := range vals {
			if a+b == 2020 {
				fmt.Println(a, b, a*b)
				return
			}
		}
	}
}

func problem1b(lines []string) {
	vals := make([]int, len(lines))
	for i, line := range lines {
		vals[i] = mustAtoi(line)
	}

	for _, a := range vals {
		for _, b := range vals {
			for _, c := range vals {
				if a+b+c == 2020 {
					fmt.Println(a, b, c, a*b*c)
					return
				}
			}
		}
	}
}
