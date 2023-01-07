package main

import (
	"container/heap"
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
	data, e := os.ReadFile("input24")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	grid := parseInput(data)
	state := State{grid.start, grid.end, 0, &grid}
	moves := countMoves(state)
	return moves
}

func task2(data string) int {
	grid := parseInput(data)
	moves1 := countMoves(State{grid.start, grid.end, 0, &grid})
	moves2 := countMoves(State{grid.end, grid.start, moves1, &grid})
	moves3 := countMoves(State{grid.start, grid.end, moves2, &grid})
	return moves3
}

func parseInput(data string) Grid {
	lines := getLines(data)
	height := len(lines) - 2
	width := len(lines[0]) - 2
	grid := newGrid(width, height)
	for y, line := range lines[1 : len(lines)-1] {
		for x, value := range line[1 : len(line)-1] {
			switch value {
			case '^':
				grid.up[Point{x, y}] = unit
			case 'v':
				grid.down[Point{x, y}] = unit
			case '<':
				grid.left[Point{x, y}] = unit
			case '>':
				grid.right[Point{x, y}] = unit
			}
		}
	}
	return grid
}

func countMoves(initialState State) int {
	queue := States{initialState}
	heap.Init(&queue)
	seen := map[State]Unit{}

	for len(queue) > 0 {
		s := heap.Pop(&queue).(State)

		for _, ns := range s.nextStates() {
			if _, exist := seen[ns]; exist {
				continue
			}
			seen[ns] = unit

			if ns.pos == ns.goal {
				return ns.turn
			}

			heap.Push(&queue, ns)
		}
	}
	panic(initialState)
}

type Point struct{ x, y int }

func (p Point) add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

func (p Point) distance(other Point) int {
	return abs(p.x-other.x) + abs(p.y-other.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Unit struct{}

var unit = Unit{}

type Grid struct {
	size         Point
	start        Point
	end          Point
	up           map[Point]Unit
	down         map[Point]Unit
	left         map[Point]Unit
	right        map[Point]Unit
	obstructions map[int]map[Point]Unit
}

func newGrid(width, height int) Grid {
	start := Point{0, -1}
	end := Point{width - 1, height}
	return Grid{Point{width, height}, start, end,
		map[Point]Unit{}, map[Point]Unit{}, map[Point]Unit{}, map[Point]Unit{},
		map[int]map[Point]Unit{}}
}

func (g *Grid) isObstructed(turn int, p Point) bool {
	o, exist := g.obstructions[turn]
	if !exist {
		o = map[Point]Unit{}
		for p := range g.up {
			op := Point{p.x, (p.y + (g.size.y-1)*turn) % g.size.y}
			o[op] = unit
		}
		for p := range g.down {
			op := Point{p.x, (p.y + turn) % g.size.y}
			o[op] = unit
		}
		for p := range g.left {
			op := Point{(p.x + (g.size.x-1)*turn) % g.size.x, p.y}
			o[op] = unit
		}
		for p := range g.right {
			op := Point{(p.x + turn) % g.size.x, p.y}
			o[op] = unit
		}
		g.obstructions[turn] = o
	}
	_, obstructed := o[p]
	return obstructed
}

type State struct {
	pos  Point
	goal Point
	turn int
	grid *Grid
}

func (s State) nextStates() []State {
	ns := []State{}
	moves := []Point{{0, 0}, {0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, p := range moves {
		newState := s.move(p)
		if newState.isValid() {
			ns = append(ns, newState)
		}
	}
	return ns
}

func (s State) move(p Point) State {
	return State{s.pos.add(p), s.goal, s.turn + 1, s.grid}
}

func (s State) isValid() bool {
	if s.pos == s.grid.start || s.pos == s.grid.end || s.pos == s.goal {
		return true
	}

	if s.pos.x < 0 || s.pos.x >= s.grid.size.x || s.pos.y < 0 || s.pos.y >= s.grid.size.y {
		return false
	}

	if s.grid.isObstructed(s.turn, s.pos) {
		return false
	}
	return true
}

func (s State) priority() int {
	return s.turn + s.pos.distance(s.goal)
}

type States []State

func (s States) Len() int {
	return len(s)
}

func (s States) Less(i, j int) bool {
	return s[i].priority() < s[j].priority()
}

func (s States) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *States) Push(x any) {
	*s = append(*s, x.(State))
}

func (s *States) Pop() any {
	old := *s
	length := len(old)
	item := old[length-1]
	*s = old[:length-1]
	return item
}
