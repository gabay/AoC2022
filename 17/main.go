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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	data, e := os.ReadFile("input17")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	j := Jets{strings.Trim(data, "\n"), 0}
	c := Chamber{0, map[Point]bool{}}

	for i := 0; i < 2022; i++ {
		pos := Point{2, c.top + 4}
		r := Rock{pos, i % 5}
		c.dropRock(r, &j)
	}

	return c.top
}

func task2(data string) int {
	N := 1000000000000
	j := Jets{strings.Trim(data, "\n"), 0}
	c := Chamber{0, map[Point]bool{}}

	heights := map[RelativeState]int{}
	turns := map[RelativeState]int{}
	postPeriodTurn := 0
	postPeriodExtraHeight := 0

	for i := 0; i < N; i++ {
		h := c.relativeHeights()
		relativeState := RelativeState{h[0], h[1], h[2], h[3], h[4], h[5], h[6], i % 5, j.nextIndex}
		if _, exist := heights[relativeState]; exist {
			periodTurns := i - turns[relativeState]
			periodHeight := c.top - heights[relativeState]
			periodCounts := (N - i) / periodTurns
			postPeriodTurn = (periodTurns * periodCounts) + i
			postPeriodExtraHeight = periodHeight * periodCounts
			break
		}

		heights[relativeState] = c.top
		turns[relativeState] = i
		pos := Point{2, c.top + 4}
		r := Rock{pos, i % 5}
		c.dropRock(r, &j)
	}

	for i := postPeriodTurn; i < N; i++ {
		pos := Point{2, c.top + 4}
		r := Rock{pos, i % 5}
		c.dropRock(r, &j)
	}

	return c.top + postPeriodExtraHeight
}

type Jets struct {
	data      string
	nextIndex int
}

func (j *Jets) next() int {
	dx := 0
	if j.data[j.nextIndex] == '<' {
		dx = -1
	} else {
		dx = 1
	}
	j.nextIndex = (j.nextIndex + 1) % len(j.data)
	return dx
}

type Point struct{ x, y int }

func (p *Point) add(dx, dy int) Point {
	return Point{p.x + dx, p.y + dy}
}

type Rock struct {
	pos   Point
	shape int
}

func (r *Rock) points() Shape {
	switch r.shape {
	case 0: // - shape
		return Shape{r.pos, r.pos.add(1, 0), r.pos.add(2, 0), r.pos.add(3, 0)}
	case 1: // + shape
		return Shape{r.pos.add(0, 1), r.pos.add(1, 0), r.pos.add(1, 1), r.pos.add(2, 1), r.pos.add(1, 2)}
	case 2: // J shape
		return Shape{r.pos, r.pos.add(1, 0), r.pos.add(2, 0), r.pos.add(2, 1), r.pos.add(2, 2)}
	case 3: // I shape
		return Shape{r.pos, r.pos.add(0, 1), r.pos.add(0, 2), r.pos.add(0, 3)}
	case 4: // box shape
		return Shape{r.pos, r.pos.add(1, 0), r.pos.add(0, 1), r.pos.add(1, 1)}
	default:
		panic(r)
	}
}

type Shape []Point

type Chamber struct {
	top     int
	blocked map[Point]bool
}

func (c *Chamber) pointBlocked(p Point) bool {
	if p.x < 0 || p.x >= 7 || p.y <= 0 {
		return true
	}
	_, exists := c.blocked[p]
	return exists
}

func (c *Chamber) rockBlocked(r Rock) bool {
	for _, p := range r.points() {
		if c.pointBlocked(p) {
			return true
		}
	}
	return false
}

func (c *Chamber) restRock(r Rock) {
	for _, pos := range r.points() {
		c.blocked[pos] = true
		if c.top < pos.y {
			c.top = pos.y
		}
	}
}

func (c *Chamber) dropRock(r Rock, j *Jets) {
	for {
		// try to move by jet
		rockAfterJet := Rock{r.pos.add(j.next(), 0), r.shape}
		if !c.rockBlocked(rockAfterJet) {
			r = rockAfterJet
		}
		// try to move down
		rockDown := Rock{r.pos.add(0, -1), r.shape}
		if c.rockBlocked(rockDown) {
			break
		}
		r = rockDown
	}
	c.restRock(r)
}

func (c Chamber) relativeHeights() []int {
	relativeHeights := []int{0, 0, 0, 0, 0, 0, 0}
	for p := range c.blocked {
		relativeHeights[p.x] = max(relativeHeights[p.x], p.y)
	}
	for i := 0; i < 7; i++ {
		relativeHeights[i] -= c.top
	}
	return relativeHeights
}

type RelativeState struct {
	h1, h2, h3, h4, h5, h6, h7 int
	rockType                   int
	jetIndex                   int
}
