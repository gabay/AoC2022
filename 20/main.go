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

func atoiList(s []string) []int {
	ints := []int{}
	for _, ss := range s {
		ints = append(ints, atoi(ss))
	}
	return ints
}

func main() {
	data, e := os.ReadFile("input20")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	numbers := parseInput(data)
	newNumbers := shuffle(numbers, 3)
	zeroIndex := getIndex(newNumbers, 0)
	a := newNumbers[(zeroIndex+1000)%len(newNumbers)]
	b := newNumbers[(zeroIndex+2000)%len(newNumbers)]
	c := newNumbers[(zeroIndex+3000)%len(newNumbers)]
	return a + b + c
}

func task2(data string) int {
	numbers := parseInput(data)
	for i := range numbers {
		numbers[i] *= 811589153
	}

	newNumbers := shuffle(numbers, 10)
	zeroIndex := getIndex(newNumbers, 0)
	a := newNumbers[(zeroIndex+1000)%len(newNumbers)]
	b := newNumbers[(zeroIndex+2000)%len(newNumbers)]
	c := newNumbers[(zeroIndex+3000)%len(newNumbers)]
	return a + b + c
}

func parseInput(data string) []int {
	return atoiList(getLines(data))
}

func shuffle(numbers []int, rounds int) []int {
	next := make([]int, len(numbers))
	prev := make([]int, len(numbers))
	for i := range numbers {
		next[i] = (i + 1) % len(numbers)
		prev[i] = (len(numbers) + i - 1) % len(numbers)
	}

	for j := 0; j < rounds; j++ {
		for i, n := range numbers {
			swap(next, prev, i, n%(len(numbers)-1))
		}
	}

	newNumbers := []int{numbers[0]}
	for i := next[0]; i != 0; i = next[i] {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
}

func swap(next []int, prev []int, i int, amount int) {
	// unlink
	next[prev[i]] = next[i]
	prev[next[i]] = prev[i]

	newPrev := prev[i]
	for ; amount > 0; amount-- {
		newPrev = next[newPrev]
	}
	for ; amount < 0; amount++ {
		newPrev = prev[newPrev]
	}
	// link
	newNext := next[newPrev]
	next[newPrev] = i
	prev[newNext] = i
	next[i] = newNext
	prev[i] = newPrev
}

func getIndex(numbers []int, value int) int {
	for i, v := range numbers {
		if v == value {
			return i
		}
	}
	panic(numbers)
}
