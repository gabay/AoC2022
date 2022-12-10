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

func main() {
	data, e := os.ReadFile("input5")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) string {
	parts := strings.Split(data, "\n\n")
	stacks := getStacks(parts[0])
	for _, line := range getLines(parts[1]) {
		count, from, to := parseLine(line)
		stacks.move(count, from, to)
	}
	return stacks.top()
}

func task2(data string) string {
	parts := strings.Split(data, "\n\n")
	stacks := getStacks(parts[0])
	for _, line := range getLines(parts[1]) {
		count, from, to := parseLine(line)
		stacks.move2(count, from, to)
	}
	return stacks.top()
}

type Stack struct {
	items []byte
}

func (s *Stack) push(b byte) {
	s.items = append(s.items, b)
}

func (s *Stack) pop() byte {
	b := s.top()
	s.items = s.items[:len(s.items)-1]
	return b
}

func (s *Stack) top() byte {
	return s.items[len(s.items)-1]
}

type Stacks struct {
	stacks []Stack
}

func (s *Stacks) move(count int, from int, to int) {
	for i := 0; i < count; i++ {
		s.stacks[to-1].push(s.stacks[from-1].pop())
	}
}

func (s *Stacks) move2(count int, from int, to int) {
	tempStack := Stack{}
	for i := 0; i < count; i++ {
		tempStack.push(s.stacks[from-1].pop())
	}
	for i := 0; i < count; i++ {
		s.stacks[to-1].push(tempStack.pop())
	}
}

func (s *Stacks) top() string {
	builder := strings.Builder{}
	for _, stack := range s.stacks {
		builder.WriteByte(stack.top())
	}
	return builder.String()
}

func getStacks(data string) Stacks {
	lines := getLines(data)
	n := (len(lines[0]) + 3) / 4
	stacks := Stacks{make([]Stack, n)}
	for row := len(lines) - 2; row >= 0; row-- {
		for stackIndex := 0; stackIndex < n; stackIndex++ {
			b := lines[row][stackIndex*4+1]
			if b != ' ' {
				stacks.stacks[stackIndex].push(b)
			}
		}
	}
	return stacks
}

func parseLine(line string) (int, int, int) {
	words := strings.Split(line, " ")
	count, ok := strconv.Atoi(words[1])
	check(ok)
	from, ok := strconv.Atoi(words[3])
	check(ok)
	to, ok := strconv.Atoi(words[5])
	check(ok)
	return count, from, to
}
