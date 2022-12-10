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

func getLines(s string) []string {
	return strings.Split(strings.Trim(s, "\n"), "\n")
}

func main() {
	data, e := os.ReadFile("input6")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	for i := 0; i < len(data)-4; i++ {
		chars := map[byte]int{}
		chars[data[i]] = 1
		chars[data[i+1]] = 1
		chars[data[i+2]] = 1
		chars[data[i+3]] = 1

		if len(chars) == 4 {
			return i + 4
		}
	}
	panic(data)
}

func task2(data string) int {
	for i := 0; i < len(data)-14; i++ {
		chars := map[byte]int{}
		for j := 0; j < 14; j++ {
			chars[data[i+j]] = 1
		}

		if len(chars) == 14 {
			return i + 14
		}
	}
	panic(data)
}
