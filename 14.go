package aoc2020

import (
	"fmt"

	. "aoc2020/helpers"
)

type mask struct {
	mask uint64
	set  uint64
}

type assignment struct {
	addr  uint64
	value uint64
}

func (m mask) Apply(v uint64) uint64 {
	return (v & m.mask) | (m.set &^ m.mask)
}

func (m mask) AllVariants(in uint64) (v []uint64) {
	v = []uint64{in | m.set}
	for i := 0; i < 64; i++ {
		curbit := uint64(1 << i)
		if m.mask&curbit != 0 {
			var v2 []uint64
			for _, x := range v {
				v2 = append(v2, x&^curbit, x|curbit)
			}
			v = v2
			fmt.Printf("%b %d\n", m.mask, len(v))
		}
	}
	return v
}

func ParseMask(s string) (m mask) {
	for _, c := range s {
		m.mask, m.set = m.mask<<1, m.set<<1
		switch c {
		case 'X':
			m.mask = m.mask | 1
		case '0', '1':
			m.set = m.set | uint64(c-'0')
		}
	}
	return m
}

func ParseAssignment(line string) (a assignment) {
	addrS, valS := Split2(line, " = ")
	a.value = uint64(Atoi(valS))
	a.addr = uint64(Atoi(addrS[len("mem[") : len(addrS)-1]))
	return a
}

type maskOrAssignment struct {
	isMask bool
	mask
	assignment
}

func ParseMemProgram(lines []string) (as []maskOrAssignment) {
	for _, line := range lines {
		leftS, rightS := Split2(line, " = ")
		if leftS == "mask" {
			as = append(as, maskOrAssignment{isMask: true, mask: ParseMask(rightS)})
		} else {
			as = append(as, maskOrAssignment{isMask: false, assignment: ParseAssignment(line)})
		}
	}
	return as
}

func Problem14a(lines []string) {
	as := ParseMemProgram(lines)
	mem := make(map[uint64]uint64)

	var m mask
	for _, a := range as {
		if a.isMask {
			m = a.mask
		} else {
			mem[a.assignment.addr] = m.Apply(a.assignment.value)
		}
	}

	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	fmt.Println(sum)
}

func Problem14b(lines []string) {
	as := ParseMemProgram(lines)
	mem := make(map[uint64]uint64)

	var m mask
	for _, a := range as {
		if a.isMask {
			m = a.mask
		} else {
			for _, addr := range m.AllVariants(a.assignment.addr) {
				fmt.Println(a.assignment.value, "to", addr)
				mem[addr] = a.assignment.value
			}
		}
	}

	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	fmt.Println(sum)
}
