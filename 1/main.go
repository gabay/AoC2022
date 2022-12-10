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

func main() {
	data, e := os.ReadFile("input1")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	calories := getCalories(data)
	sort.Ints(calories)
	return calories[len(calories)-1]
}

func task2(data string) int {
	calories := getCalories(data)
	sort.Ints(calories)
	return sum(calories[len(calories)-3:])
}

func getCalories(data string) []int {
	chunks := []int{0}
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			chunks = append(chunks, 0)
		} else {
			value, e := strconv.Atoi(line)
			check(e)
			chunks[len(chunks)-1] += value
		}
	}
	return chunks
}

func sum(items []int) int {
	result := 0
	for _, item := range items {
		result += item
	}
	return result
}
