package aoc2020

import (
	"fmt"
	"regexp"

	. "aoc2020/helpers"
)

var mathExprRegexp = regexp.MustCompile(`[0-9]+|\(|\)|\*|\+`)

type exprTokens []string

type expr interface{ Eval() int }

type opExpr struct {
	L, R expr
	Op   string
}
type intExpr int

func (op opExpr) Eval() int {
	switch op.Op {
	case "*":
		return op.L.Eval() * op.R.Eval()
	case "+":
		return op.L.Eval() + op.R.Eval()
	default:
		panic(op.Op)
	}
}
func (v intExpr) Eval() int { return int(v) }

func (op opExpr) String() string { return fmt.Sprintf("%s(%s)(%s)", op.Op, op.L, op.R) }
func (v intExpr) String() string { return fmt.Sprintf("%d", v) }

func (toks exprTokens) Parse() expr {
	e, skip := toks.parseExpr()
	if skip != len(toks) {
		panic(skip)
	}
	return e
}

func (toks exprTokens) parseOperand() (e expr, skip int) {
	switch t := toks[0]; t {
	case "(":
		nested, nestedSkip := toks[1:].parseExpr()
		if toks[1+nestedSkip] != ")" {
			return nil, 0
		}
		return nested, nestedSkip + 2
	default: // number
		return intExpr(Atoi(t)), 1
	case "*", "+", ")":
		return nil, 0
	}
}

func (toks exprTokens) parseExpr() (e expr, skip int) {
	e, skip = toks.parseOperand()
	for skip < len(toks) {
		switch t := toks[skip]; t {
		case "+", "*":
			nested, nestedSkip := toks[skip+1:].parseOperand()
			e, skip = opExpr{L: e, R: nested, Op: t}, skip+1+nestedSkip
		default:
			return e, skip
		}
	}
	return e, skip
}

func (toks exprTokens) Parse2() expr {
	e, skip := toks.parseExpr2()
	if skip != len(toks) {
		panic(skip)
	}
	return e
}

func (toks exprTokens) parseOperand2() (e expr, skip int) {
	defer func() { fmt.Println("parseOperand2", toks[:skip]) }()
	switch t := toks[0]; t {
	case "(":
		nested, nestedSkip := toks[1:].parseExpr2()
		if toks[1+nestedSkip] != ")" {
			return nil, 0
		}
		return nested, nestedSkip + 2
	default: // number
		return intExpr(Atoi(t)), 1
	case "*", "+", ")":
		return nil, 0
	}
}

func (toks exprTokens) parseSum2() (e expr, skip int) {
	defer func() { fmt.Println("parseSum2", toks[:skip]) }()
	e, skip = toks.parseOperand2()
	for skip < len(toks) {
		switch t := toks[skip]; t {
		case "+":
			nested, nestedSkip := toks[skip+1:].parseOperand2()
			e, skip = opExpr{L: e, R: nested, Op: t}, skip+1+nestedSkip
		default:
			return e, skip
		}
	}
	return e, skip
}

func (toks exprTokens) parseExpr2() (e expr, skip int) {
	defer func() { fmt.Println("parseExpr2", toks[:skip]) }()
	e, skip = toks.parseSum2()
	for skip < len(toks) {
		switch t := toks[skip]; t {
		case "*":
			nested, nestedSkip := toks[skip+1:].parseSum2()
			e, skip = opExpr{L: e, R: nested, Op: t}, skip+1+nestedSkip
		default:
			return e, skip
		}
	}
	return e, skip
}

func ParseExpressions(lines []string) (exprs []expr) {
	for _, line := range lines {
		toks := exprTokens(mathExprRegexp.FindAllString(line, -1))
		exprs = append(exprs, toks.Parse())
	}
	return exprs
}

func ParseExpressions2(lines []string) (exprs []expr) {
	for _, line := range lines {
		toks := exprTokens(mathExprRegexp.FindAllString(line, -1))
		e := toks.Parse2()
		fmt.Println(line, e)
		exprs = append(exprs, e)
	}
	return exprs
}

func Problem18a(lines []string) {
	exprs := ParseExpressions(lines)

	sum := 0
	for _, e := range exprs {
		sum += e.Eval()
	}
	fmt.Println(sum)
}

func Problem18b(lines []string) {
	exprs := ParseExpressions2(lines)

	sum := 0
	for _, e := range exprs {
		sum += e.Eval()
	}
	fmt.Println(sum)
}
