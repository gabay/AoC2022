package main

import (
	"encoding/json"
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
	data, e := os.ReadFile("input13")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	signals := parseInput(data)
	result := 0
	for i := 0; i < len(signals)/2; i++ {
		left := signals[i*2]
		right := signals[i*2+1]
		if Compare(left, right) == -1 {
			result += i + 1
		}
	}
	return result
}

func task2(data string) int {
	signals := parseInput(data)
	signal1, signal2 := makeSignal("[[2]]"), makeSignal("[[6]]")
	signals = append(append(signals, signal1), signal2)
	sort.Sort(signals)

	indices := []int{}
	for i, v := range signals {
		if Compare(v, signal1) == 0 || Compare(v, signal2) == 0 {
			indices = append(indices, i+1)
		}
	}

	return indices[0] * indices[1]
}

func parseInput(data string) Signals {
	signals := Signals{}
	for _, line := range getLines(data) {
		if line != "" {
			signals = append(signals, makeSignal(line))
		}
	}
	return signals
}

func makeSignal(line string) Signal {
	var signal Signal
	json.Unmarshal([]byte(line), &signal)
	return signal
}

type Signal []any

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Compare(a, b Signal) int {
	commonBound := MinInt(len(a), len(b))
	for i := 0; i < commonBound; i++ {
		ae := a[i]
		be := b[i]
		aFloat, aok := ae.(float64)
		bFloat, bok := be.(float64)
		if aok && bok {
			if result := CompareInt(int(aFloat), int(bFloat)); result != 0 {
				return result
			}
		} else {
			if aok {
				ae = []any{ae}
			}
			if bok {
				be = []any{be}
			}
			if result := Compare(ae.([]any), be.([]any)); result != 0 {
				return result
			}
		}
	}
	return CompareInt(len(a), len(b))
}

func CompareInt(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

type Signals []Signal

func (s Signals) Len() int {
	return len(s)
}

func (s Signals) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Signals) Less(i, j int) bool {
	return Compare(s[i], s[j]) == -1
}
