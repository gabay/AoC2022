package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	data, e := os.ReadFile("input23")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	grid := parseInput(data)
	for i := 0; i < 10; i++ {
		grid.next()
	}
	return grid.getEmptySquares()
}

func task2(data string) int {
	grid := parseInput(data)
	round := 1
	for grid.next() {
		round++
	}
	return round
}

func parseInput(data string) Grid {
	grid := newGrid([]int{3, 1, 2, 0}) // N, S, W, E
	for y, line := range strings.Split(data, "\n") {
		for x, value := range line {
			if value == '#' {
				grid.elves[Point{x, y}] = unit
			}
		}
	}
	return grid
}

type Point struct{ x, y int }

func (p Point) add(x, y int) Point {
	return Point{p.x + x, p.y + y}
}

func (p Point) around() []Point {
	points := []Point{}
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x != 0 || y != 0 {
				points = append(points, p.add(x, y))
			}
		}
	}
	return points
}

type Unit struct{}

var unit = Unit{}

type Dirs []int

func (d Dirs) rotate() Dirs {
	return append(d[1:], d[0])
}

type Grid struct {
	elves map[Point]Unit
	dirs  Dirs
}

func newGrid(dirs []int) Grid {
	return Grid{make(map[Point]Unit), dirs}
}

func (g Grid) getEmptySquares() int {
	top, left := math.MaxInt, math.MaxInt
	bottom, right := math.MinInt, math.MinInt
	for p := range g.elves {
		left = min(left, p.x)
		right = max(right, p.x)
		top = min(top, p.y)
		bottom = max(bottom, p.y)
	}
	width := right - left + 1
	height := bottom - top + 1
	return (width * height) - len(g.elves)
}

func (g *Grid) next() bool { // Return true if at least one elf could move, false otherwise
	// 1 - proposes
	proposes := map[Point]Point{}
	collisions := map[Point]Unit{}
	for elf := range g.elves {
		if propose, ok := g.propose(elf); ok {
			if _, exist := proposes[propose]; exist {
				collisions[propose] = unit
			} else {
				proposes[propose] = elf
			}
		}
	}
	if len(proposes) == 0 {
		return false
	}

	// 2 - moves
	for dst, src := range proposes {
		if _, exist := collisions[dst]; !exist {
			delete(g.elves, src)
			g.elves[dst] = unit
		}
	}

	// 3 - rotate
	g.dirs = g.dirs.rotate()

	return true
}

func (g Grid) propose(elf Point) (Point, bool) {
	isAlone := true
	for _, p := range elf.around() {
		if g.hasElfIn(p) {
			isAlone = false
			break
		}
	}
	if isAlone {
		return Point{0, 0}, false
	}

	NW, N, NE := elf.add(-1, -1), elf.add(0, -1), elf.add(1, -1)
	W, E := elf.add(-1, 0), elf.add(1, 0)
	SW, S, SE := elf.add(-1, 1), elf.add(0, 1), elf.add(1, 1)
	for _, dir := range g.dirs {
		switch dir {
		case 0: // E
			if !g.hasElfIn(NE) && !g.hasElfIn(E) && !g.hasElfIn(SE) {
				return E, true
			}
		case 1: // S
			if !g.hasElfIn(SW) && !g.hasElfIn(S) && !g.hasElfIn(SE) {
				return S, true
			}
		case 2: // W
			if !g.hasElfIn(NW) && !g.hasElfIn(W) && !g.hasElfIn(SW) {
				return W, true
			}
		case 3: // N
			if !g.hasElfIn(NW) && !g.hasElfIn(N) && !g.hasElfIn(NE) {
				return N, true
			}
		default:
			panic(g)
		}
	}
	return Point{0, 0}, false
}

func (g Grid) hasElfIn(p Point) bool {
	_, exist := g.elves[p]
	return exist
}
