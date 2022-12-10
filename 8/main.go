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
	data, e := os.ReadFile("input8")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	grid := parseInput(data)
	result := 0
	for x := 0; x < grid.width(); x++ {
		for y := 0; y < grid.height(); y++ {
			if grid.isVisible(x, y) {
				result += 1
			}
		}
	}

	return result
}

func task2(data string) int {
	grid := parseInput(data)
	result := 0
	for x := 0; x < grid.width(); x++ {
		for y := 0; y < grid.height(); y++ {
			score := grid.getScenicScore(x, y)
			if result < score {
				result = score
			}
		}
	}
	return result
}

type Grid struct {
	items [][]byte
}

func (g Grid) height() int {
	return len(g.items)
}

func (g Grid) width() int {
	return len(g.items[0])
}

func (g Grid) isVisible(x int, y int) bool {
	if x == 0 || y == 0 || x+1 == g.width() || y+1 == g.height() {
		return true
	}

	height := g.get(x, y)
	return (g.checkAllValuesLowerThan(0, x-1, y, y, height) ||
		g.checkAllValuesLowerThan(x+1, g.width()-1, y, y, height) ||
		g.checkAllValuesLowerThan(x, x, 0, y-1, height) ||
		g.checkAllValuesLowerThan(x, x, y+1, g.height()-1, height))
}

func (g Grid) get(x int, y int) int {
	return int(g.items[y][x])
}

func (g Grid) checkAllValuesLowerThan(x1, x2, y1, y2, value int) bool {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if g.get(x, y) >= value {
				return false
			}
		}
	}
	return true
}

func (g Grid) getScenicScore(x, y int) int {
	if x == 0 || y == 0 || x+1 == g.width() || y+1 == g.height() {
		return 0
	}

	left := g.countVisibleLeft(x, y)
	right := g.countVisibleRight(x, y)
	up := g.countVisibleUp(x, y)
	down := g.countVisibleDown(x, y)
	return up * down * left * right
}

func (g Grid) countVisibleLeft(x, y int) int {
	count := 0
	for xi := x - 1; xi >= 0; xi-- {
		count += 1
		if g.get(x, y) <= g.get(xi, y) {
			break
		}
	}
	return count
}
func (g Grid) countVisibleRight(x, y int) int {
	count := 0
	for xi := x + 1; xi < g.width(); xi++ {
		count += 1
		if g.get(x, y) <= g.get(xi, y) {
			break
		}
	}
	return count
}
func (g Grid) countVisibleUp(x, y int) int {
	count := 0
	for yi := y - 1; yi >= 0; yi-- {
		count += 1
		if g.get(x, y) <= g.get(x, yi) {
			break
		}
	}
	return count
}
func (g Grid) countVisibleDown(x, y int) int {
	count := 0
	for yi := y + 1; yi < g.height(); yi++ {
		count += 1
		if g.get(x, y) <= g.get(x, yi) {
			break
		}
	}
	return count
}

func parseInput(data string) Grid {
	g := Grid{}
	for _, line := range getLines(data) {
		row := []byte(line)
		g.items = append(g.items, row)
	}
	return g
}
