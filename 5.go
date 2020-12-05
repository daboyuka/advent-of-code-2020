package main

import "fmt"

type boardingPass struct {
	Row, Col int
}

func ParseBoardingPass(line string) (pass boardingPass) {
	for _, char := range line[:7] {
		pass.Row <<= 1
		if char == 'B' {
			pass.Row++
		}
	}
	for _, char := range line[7:] {
		pass.Col <<= 1
		if char == 'R' {
			pass.Col++
		}
	}
	return pass
}

func (p boardingPass) SeatID() int { return p.Row*8 + p.Col }

func ParseBoardingPasses(lines []string) (passes []boardingPass) {
	for _, line := range lines {
		passes = append(passes, ParseBoardingPass(line))
	}
	return passes
}

func problem5a(lines []string) {
	maxSeatID := -1
	for _, pass := range ParseBoardingPasses(lines) {
		if seatID := pass.SeatID(); maxSeatID < seatID {
			maxSeatID = seatID
		}
	}
	fmt.Println(maxSeatID)
}

func problem5b(lines []string) {
	const numIds = 1 << 10
	seatIdsFound := make([]bool, numIds)

	for _, pass := range ParseBoardingPasses(lines) {
		seatIdsFound[pass.SeatID()] = true
	}

	for i, found := range seatIdsFound {
		if !found && i > 0 && i < numIds-1 && seatIdsFound[i-1] && seatIdsFound[i+1] {
			fmt.Println(i)
		}
	}
}
