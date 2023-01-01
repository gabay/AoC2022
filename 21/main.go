package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getLines(s string) []string {
	return strings.Split(strings.Trim(s, "\n"), "\n")
}

func atoi(s string) int {
	result, ok := strconv.Atoi(s)
	check(ok)
	return result
}

func main() {
	data, e := os.ReadFile("input21")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	monkeys := parseInput(data)
	return monkeys.eval("root")
}

func task2(data string) int {
	monkeys := parseInput(data)
	result := monkeys.isolateHUMN()
	return result
}

func parseInput(data string) Monkeys {
	monkeys := Monkeys{}
	for _, line := range getLines(data) {
		name, expr, ok := strings.Cut(line, ": ")
		if !ok {
			panic(line)
		}
		monkeys[name] = &Monkey{expr}
	}
	return monkeys
}

type Monkey struct {
	expr string
}

type Monkeys map[string]*Monkey

func (ms Monkeys) eval(name string) int {
	if ms.isSimple(name) {
		return atoi(ms[name].expr)
	}

	va, operand, vb := ms.subexpressions(name)
	return eval(va, vb, operand)
}

func (ms Monkeys) isSimple(name string) bool {
	return !strings.Contains(ms[name].expr, " ")
}

func eval(a, b int, operand string) int {
	switch operand {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	panic(operand)
}

func evalOppositeRHS(rhs, value int, operand string) int {
	switch operand {
	case "+":
		return value - rhs
	case "-":
		return value + rhs
	case "*":
		return value / rhs
	case "/":
		return value * rhs
	}
	panic(operand)
}

func evalOppositeLHS(lhs, value int, operand string) int {
	switch operand {
	case "+":
		return value - lhs
	case "-":
		return lhs - value
	case "*":
		return value / lhs
	case "/":
		return lhs / value
	}
	panic(operand)
}

func (ms Monkeys) submonkeys(name string) (string, string, string) {
	parts := strings.Split(ms[name].expr, " ")
	ma, operand, mb := parts[0], parts[1], parts[2]
	return ma, operand, mb
}

func (ms Monkeys) subexpressions(name string) (int, string, int) {
	ma, operand, mb := ms.submonkeys(name)
	va, vb := ms.eval(ma), ms.eval(mb)
	return va, operand, vb
}

func (ms Monkeys) isolateHUMN() int {
	a, _, b := ms.submonkeys("root")
	if !ms.hasHUMN(a) {
		a, b = b, a
		if !ms.hasHUMN(a) {
			panic(ms)
		}
	}

	value := ms.eval(b)

	for !ms.isSimple(a) {
		aa, op, ab := ms.submonkeys(a)
		if ms.hasHUMN(aa) {
			value = evalOppositeRHS(ms.eval(ab), value, op)
			a = aa
		} else {
			value = evalOppositeLHS(ms.eval(aa), value, op)
			a = ab
		}
	}
	return value
}

func (ms Monkeys) hasHUMN(name string) bool {
	if name == "humn" {
		return true
	}
	if ms.isSimple(name) {
		return false
	}

	a, _, b := ms.submonkeys(name)
	return ms.hasHUMN(a) || ms.hasHUMN(b)
}
