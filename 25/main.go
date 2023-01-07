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
	data, e := os.ReadFile("input25")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) string {
	numbers := parseInput(data)
	sum := 0
	for _, qn := range numbers {
		sum += QuinaryNumber(qn).toInt()
	}
	return string(toQuinaryNumber(sum))
}

func task2(data string) int {
	return 0
}

func parseInput(data string) []string {
	return getLines(data)
}

type QuinaryNumber string

func (qn QuinaryNumber) toInt() int {
	n := 0
	for _, c := range qn {
		n *= 5
		switch c {
		case '2':
			n += 2
		case '1':
			n += 1
		case '-':
			n -= 1
		case '=':
			n -= 2
		}
	}
	return n
}

func toQuinaryNumber(n int) string {
	// split into quinary digits
	digits := []int{}
	for n != 0 {
		digits = append(digits, n%5)
		n = n / 5
	}
	if digits[len(digits)-1] > 2 {
		digits = append(digits, 0)
	}

	// shift to range [-2..2]
	for i := range digits {
		for digits[i] > 2 {
			digits[i] -= 5
			digits[i+1] += 1
		}
	}

	// reverse so MSB is first
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	// concat to string
	s := []string{}
	for _, d := range digits {
		switch d {
		case -2:
			s = append(s, "=")
		case -1:
			s = append(s, "-")
		case 0, 1, 2:
			s = append(s, strconv.Itoa(d))
		default:
			panic(n)
		}
	}
	return strings.Join(s, "")
}
