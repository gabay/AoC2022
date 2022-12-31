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

func atoi(s string) int {
	result, ok := strconv.Atoi(s)
	check(ok)
	return result
}

func main() {
	data, e := os.ReadFile("input12")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	g := parseInput(data)
	result := dfs(g, g.GetStart(), g.GetEnd())
	return result
}

func task2(data string) int {
	result := task1(data)

	g := parseInput(data)
	end := g.GetEnd()
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			p := Point{x, y}
			if g.Get(p) == 'a' {

				if result2 := dfs(g, p, end); result2 != -1 && result2 < result {
					result = result2
				}
			}
		}
	}
	return result
}

func parseInput(data string) Grid {
	lines := getLines(data)
	return Grid{lines}
}

type Point struct{ x, y int }

func (p Point) Left() Point {
	return Point{p.x - 1, p.y}
}
func (p Point) Right() Point {
	return Point{p.x + 1, p.y}
}
func (p Point) Up() Point {
	return Point{p.x, p.y - 1}
}
func (p Point) Down() Point {
	return Point{p.x, p.y + 1}
}

type Grid struct {
	cells []string
}

func (g Grid) Width() int {
	return len(g.cells[0])
}

func (g Grid) Height() int {
	return len(g.cells)
}

func (g Grid) GetStart() Point {
	for y, line := range g.cells {
		if x := strings.Index(line, "S"); x != -1 {
			return Point{x, y}
		}
	}
	panic(g)
}

func (g Grid) GetEnd() Point {
	for y, line := range g.cells {
		if x := strings.Index(line, "E"); x != -1 {
			return Point{x, y}
		}
	}
	panic(g)
}

func (g Grid) Get(p Point) byte {
	return g.cells[p.y][p.x]
}

func (g Grid) NextPoints(p Point) []Point {
	nextPoints := []Point{}
	if left := p.Left(); left.x >= 0 && g.CanMove(p, left) {
		nextPoints = append(nextPoints, left)
	}
	if right := p.Right(); right.x < g.Width() && g.CanMove(p, right) {
		nextPoints = append(nextPoints, right)
	}
	if up := p.Up(); up.y >= 0 && g.CanMove(p, up) {
		nextPoints = append(nextPoints, up)
	}
	if down := p.Down(); down.y < g.Height() && g.CanMove(p, down) {
		nextPoints = append(nextPoints, down)
	}
	return nextPoints
}

func (g Grid) CanMove(a Point, b Point) bool {
	return CanMove(g.Get(a), g.Get(b))
}

func CanMove(a byte, b byte) bool {
	if a == 'S' {
		a = 'a'
	}
	if b == 'E' {
		b = 'z'
	}
	return a+1 >= b
}

func dfs(g Grid, start Point, end Point) int {
	queue := make(chan Point, g.Width()*g.Height())
	queue <- start
	// queue := []Point{start}
	depth := map[Point]int{start: 0}
	for {
		if len(queue) == 0 {
			return -1
		}

		point := <-queue
		// point := queue[0]
		// queue = queue[1:]
		nextDepth := depth[point] + 1
		for _, nextPoint := range g.NextPoints(point) {
			if nextPoint == end {
				return nextDepth
			}

			if _, exists := depth[nextPoint]; !exists {
				depth[nextPoint] = nextDepth
				queue <- nextPoint
				// queue = append(queue, nextPoint)
			}
		}
	}
}
