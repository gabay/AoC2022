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

func lines(s string) []string {
	return strings.Split(strings.Trim(s, "\n"), "\n")
}

func main() {
	data, e := os.ReadFile("input4")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	result := 0
	for _, line := range lines(data) {
		a, b := parseLine(line)
		if a.contains(b) || b.contains(a) {
			result += 1
		}
	}
	return result
}

func task2(data string) int {
	result := 0
	for _, line := range lines(data) {
		a, b := parseLine(line)
		if a.overlaps(b) {
			result += 1
		}
	}
	return result
}

type Assignment struct {
	start int
	end   int
}

func toAssignment(s string) Assignment {
	parts := strings.Split(s, "-")
	start, e := strconv.Atoi(parts[0])
	check(e)
	end, e := strconv.Atoi(parts[1])
	check(e)
	return Assignment{start, end}
}

func (a Assignment) contains(b Assignment) bool {
	return a.start <= b.start && a.end >= b.end
}

func (a Assignment) overlaps(b Assignment) bool {
	return a.start <= b.end && a.end >= b.start
}

func parseLine(line string) (Assignment, Assignment) {
	parts := strings.Split(line, ",")
	a := toAssignment(parts[0])
	b := toAssignment(parts[1])
	return a, b
}
