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
	data, e := os.ReadFile("input2")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	score := 0
	for _, line := range strings.Split(data, "\n") {
		if line != "" {
			a := getMove(string(line[0]))
			b := getMove(string(line[2]))
			score += b.score() + b.scoreAgainst(a)
		}

	}
	return score
}

func task2(data string) int {
	score := 0
	for _, line := range strings.Split(data, "\n") {
		if line != "" {
			a := getMove(string(line[0]))
			outcome := getOutcome(string(line[2]))
			b := outcome.getMyMove(a)
			score += b.score() + b.scoreAgainst(a)
		}
	}
	return score
}

type Move int

const (
	Rock Move = iota
	Paper
	Scissors
)

func (m Move) score() int {
	switch m {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	panic(m)
}

func (m Move) scoreAgainst(mEnemy Move) int {
	diff := (m.score() - mEnemy.score() + 3) % 3
	switch diff {
	case 0:
		return 3
	case 1:
		return 6
	case 2:
		return 0
	}
	panic(m)
}

func getMove(s string) Move {
	switch s {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	panic(s)
}

type Outcome int

const (
	Lose Outcome = iota
	Draw
	Win
)

func getOutcome(s string) Outcome {
	switch s {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	}
	panic(s)
}

func (o Outcome) getMyMove(mEnemy Move) Move {
	switch o {
	case Draw:
		return mEnemy
	case Win:
		return mEnemy.getWinningMove()
	case Lose:
		return mEnemy.getLosingMove()
	}
	panic(o)
}

func (m Move) getWinningMove() Move {
	switch m {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	}
	panic(m)
}

func (m Move) getLosingMove() Move {
	switch m {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	}
	panic(m)
}
