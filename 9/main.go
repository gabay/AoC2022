package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getLines(s string) []string {
	return strings.Split(strings.Trim(s, "\n"), "\n")
}

func main() {
	data, e := os.ReadFile("input9")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	rope := makeRope()
	tailPositions := map[Point]int{}
	for _, direction := range parseInput(data) {
		rope.move(direction)
		tailPositions[rope.tail] = 1
	}
	return len(tailPositions)
}

func task2(data string) int {
	n := 10
	rope := makeRope2(n)
	tailPositions := map[Point]int{}
	for _, direction := range parseInput(data) {
		rope.move(direction)
		tailPositions[rope.knots[n-1]] = 1
	}
	return len(tailPositions)
}

type Direction byte

const (
	Up    Direction = 'U'
	Down  Direction = 'D'
	Left  Direction = 'L'
	Right Direction = 'R'
)

type Point struct {
	x int
	y int
}

func (p *Point) move(d Direction) {
	switch d {
	case Up:
		p.y -= 1
	case Down:
		p.y += 1
	case Left:
		p.x -= 1
	case Right:
		p.x += 1
	}
}

type Rope struct {
	head Point
	tail Point
}

type Rope2 struct {
	knots []Point
}

func makeRope() Rope {
	return Rope{Point{}, Point{}}
}

func makeRope2(knots int) Rope2 {
	return Rope2{make([]Point, knots)}
}

func (r *Rope) move(d Direction) {
	r.head.move(d)
	dx := abs(r.head.x - r.tail.x)
	dy := abs(r.head.y - r.tail.y)
	if dx >= 2 {
		r.tail.x = (r.head.x + r.tail.x) / 2
		r.tail.y = r.head.y
	} else if dy >= 2 {
		r.tail.x = r.head.x
		r.tail.y = (r.head.y + r.tail.y) / 2
	}
}

func (r *Rope2) move(d Direction) {
	r.knots[0].move(d)
	for i := 0; i < len(r.knots)-1; i++ {
		r.propagateMove(i)
	}
}

func (r *Rope2) propagateMove(headIndex int) {
	head := &r.knots[headIndex]
	tail := &r.knots[headIndex+1]
	dx := head.x - tail.x
	dy := head.y - tail.y
	if abs(dx) < 2 && abs(dy) < 2 {
		return
	}

	if dx > 0 {
		tail.move(Right)
	} else if dx < 0 {
		tail.move(Left)
	}

	if dy > 0 {
		tail.move(Down)
	} else if dy < 0 {
		tail.move(Up)
	}
}

func parseInput(data string) []Direction {
	dirs := []Direction{}
	for _, line := range getLines(data) {
		parts := strings.Split(line, " ")
		direction := Direction(line[0])
		count, ok := strconv.Atoi(parts[1])
		check(ok)
		for i := 0; i < count; i++ {
			dirs = append(dirs, direction)
		}
	}
	return dirs
}
