package aoc2020

import (
	"fmt"
	"strings"

	. "aoc2020/helpers"
)

type thingy struct {
	S string
}

func ParseInput(lines []string) (depart int, buses []int) {

	depart = Atoi(lines[0])
	for _, t := range strings.Split(lines[1], ",") {
		if t == "x" {
			buses = append(buses, 0)
		} else {
			buses = append(buses, Atoi(t))
		}
	}
	return depart, buses
}

func Problem13a(lines []string) {
	depart, buses := ParseInput(lines)

	min, minBus := -1, -1
	for _, bus := range buses {
		if bus == 0 {
			continue
		}
		if mod := depart % bus; mod == 0 {
			minBus = bus
			min = 0
			break
		} else if dtime := bus - mod; min == -1 || min > dtime {
			minBus = bus
			min = dtime
		}
	}

	fmt.Println(min, minBus, min*minBus)
}

func Problem13b(lines []string) {
	_, buses := ParseInput(lines)

	chunk := 1
	t := 0
loop:
	for i, bus := range buses {
		if i == 0 {
			chunk *= bus
			continue
		} else if bus == 0 {
			continue
		}

		tMod := t % bus
		chunkMod := chunk % bus

		for chunks := 0; ; chunks++ {
			fmt.Printf("%d: %d*%d + %d + %d %% %d -> %d (%d)\n", i, chunks, chunk, t, bus, chunks*chunk+t+i, (chunks*chunk+tMod+i)%bus)
			if (chunks*chunkMod+tMod+i)%bus == 0 {

				t += chunks * chunk
				chunk *= bus
				continue loop
			}
		}
	}

	fmt.Println(t)
}
