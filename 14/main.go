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
	data, e := os.ReadFile("input14")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	grid := parseInput(data)
	maxY := grid.GetMaxY()

	result := 0
	for {
		if !grid.AddSand(maxY) {
			break
		}
		result += 1
	}
	return result
}

func task2(data string) int {
	grid := parseInput(data)
	maxY := grid.GetMaxY()

	result := 0
	for {
		result += 1
		if grid.AddSand2(maxY) == (Point{500, 0}) {
			break
		}
	}
	return result
}

type Point struct{ x, y int }

type Grid map[Point]bool

func parseInput(data string) Grid {
	grid := Grid{}
	for _, line := range getLines(data) {
		points := parsePoints(line)
		for i := 0; i < len(points)-1; i++ {
			for _, point := range pointsInLine(points[i], points[i+1]) {
				grid[point] = true
			}
		}
	}
	return grid
}

func parsePoints(line string) []Point {
	points := []Point{}
	for _, pointString := range strings.Split(line, " -> ") {
		xy := strings.Split(pointString, ",")
		points = append(points, Point{atoi(xy[0]), atoi(xy[1])})
	}
	return points
}

func pointsInLine(p0, p1 Point) []Point {
	if p0.x > p1.x || p0.y > p1.y {
		return pointsInLine(p1, p0)
	}
	points := []Point{}
	if p0.x < p1.x {
		for x := p0.x; x < p1.x+1; x++ {
			points = append(points, Point{x, p0.y})
		}
	} else if p0.y < p1.y {
		for y := p0.y; y < p1.y+1; y++ {
			points = append(points, Point{p0.x, y})
		}
	} else {
		panic(p0)
	}
	return points
}

func (g Grid) GetMaxY() int {
	maxY := 0
	for point := range g {
		if maxY < point.y {
			maxY = point.y
		}
	}
	return maxY
}

func (g Grid) AddSand(maxY int) bool {
	sand := Point{500, 0}
	for sand.y < maxY {
		down := Point{sand.x, sand.y + 1}
		if _, exist := g[down]; !exist {
			sand = down
			continue
		}
		downLeft := Point{sand.x - 1, sand.y + 1}
		if _, exist := g[downLeft]; !exist {
			sand = downLeft
			continue
		}
		downRight := Point{sand.x + 1, sand.y + 1}
		if _, exist := g[downRight]; !exist {
			sand = downRight
			continue
		}
		// rest
		g[sand] = true
		return true
	}
	// sand falls forever
	return false
}

func (g Grid) AddSand2(maxY int) Point {
	sand := Point{500, 0}
	for sand.y <= maxY {
		down := Point{sand.x, sand.y + 1}
		if _, exist := g[down]; !exist {
			sand = down
			continue
		}
		downLeft := Point{sand.x - 1, sand.y + 1}
		if _, exist := g[downLeft]; !exist {
			sand = downLeft
			continue
		}
		downRight := Point{sand.x + 1, sand.y + 1}
		if _, exist := g[downRight]; !exist {
			sand = downRight
			continue
		}
		break
	}
	// rest
	g[sand] = true
	return sand
}
