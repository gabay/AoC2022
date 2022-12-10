package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, e := os.ReadFile("input3")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	result := 0
	for _, line := range strings.Split(data, "\n") {
		middle := len(line) / 2
		items := map[Item]int{}
		for _, item := range line[:middle] {
			items[Item(item)] = 1
		}
		for _, item := range line[middle:] {
			if items[Item(item)] == 1 {
				result += Item(item).priority()
				break
			}
		}
	}
	return result
}

func task2(data string) int {
	result := 0
	rucksacks := []map[Item]int{}
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}
		rucksacks = append(rucksacks, map[Item]int{})
		for _, item := range line {
			rucksacks[len(rucksacks)-1][Item(item)] = 1
		}
	}
	for index := 0; index < len(rucksacks); index += 3 {
		a := rucksacks[index]
		b := rucksacks[index+1]
		c := rucksacks[index+2]
		for item := range a {
			_, ok1 := b[item]
			_, ok2 := c[item]
			if ok1 && ok2 {
				result += item.priority()
				break
			}
		}
	}
	return result
}

type Item byte

func (i Item) priority() int {
	if i >= 'a' && i <= 'z' {
		return int(i - 'a' + 1)
	}
	if i >= 'A' && i <= 'Z' {
		return int(i - 'A' + 27)
	}
	panic(i)
}
