package main

import (
	"fmt"
	"os"
	"sort"
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
	data, e := os.ReadFile("input11")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	monkeys := parseInput(data, 3)
	for i := 0; i < 20; i++ {
		doRound(monkeys)
	}

	inspectionCounts := getInspectionCounts(monkeys)
	sort.Sort(sort.Reverse(sort.IntSlice(inspectionCounts)))
	return inspectionCounts[0] * inspectionCounts[1]
}

func task2(data string) int {
	monkeys := parseInput(data, 1)
	for i := 0; i < 10000; i++ {
		doRound(monkeys)
	}

	inspectionCounts := getInspectionCounts(monkeys)
	sort.Sort(sort.Reverse(sort.IntSlice(inspectionCounts)))
	return inspectionCounts[0] * inspectionCounts[1]
}

func parseInput(data string, finalDivisor int) []Monkey {
	monkeys := []Monkey{}
	for _, monkeyData := range strings.Split(data, "\n\n") {
		monkeys = append(monkeys, parseMonkey(monkeyData, finalDivisor))
	}
	return monkeys
}

func parseMonkey(monkeyData string, finalDivisor int) Monkey {
	lines := getLines(monkeyData)
	itemsString := strings.Split(lines[1], ": ")[1]
	items := []int{}
	for _, itemString := range strings.Split(itemsString, ", ") {
		items = append(items, atoi(itemString))
	}
	operation := Operation(strings.Split(lines[2], " ")[6][0])
	arg2 := strings.Split(lines[2], " ")[7]
	divider := atoi(strings.Split(lines[3], " ")[5])
	monkeyTrue := atoi(strings.Split(lines[4], " ")[9])
	monkeyFalse := atoi(strings.Split(lines[5], " ")[9])

	return Monkey{items, operation, arg2, divider, monkeyTrue, monkeyFalse, 0, finalDivisor}
}

func doRound(monkeys []Monkey) {
	modulus := 1
	for _, monkey := range monkeys {
		modulus *= monkey.modulus
	}
	for i := 0; i < len(monkeys); i++ {
		monkeys[i].inspectItems(monkeys, modulus)
	}
}

func getInspectionCounts(monkeys []Monkey) []int {
	inspectionCounts := []int{}
	for _, monkey := range monkeys {
		inspectionCounts = append(inspectionCounts, monkey.inspectedItemsCount)
	}
	return inspectionCounts
}

type Monkey struct {
	items               []int
	operation           Operation
	arg2                string
	modulus             int
	monkeyTrue          int
	monkeyFalse         int
	inspectedItemsCount int
	finalDivisor        int
}

func (m *Monkey) inspectItems(monkeys []Monkey, modulus int) {
	for _, item := range m.items {
		newValue, nextMonkey := m.inspectItem(item % modulus)
		monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, newValue%modulus)
	}
	m.items = nil
}

func (m *Monkey) inspectItem(item int) (int, int) {
	m.inspectedItemsCount += 1

	arg2 := item
	if m.arg2 != "old" {
		arg2 = atoi(m.arg2)
	}

	new := 0
	switch m.operation {
	case Add:
		new = (item + arg2) / m.finalDivisor
	case Mul:
		new = (item * arg2) / m.finalDivisor
	default:
		panic(m)
	}

	if new%m.modulus == 0 {
		return new, m.monkeyTrue
	} else {
		return new, m.monkeyFalse
	}
}

type Operation byte

const (
	Add Operation = '+'
	Mul Operation = '*'
)
