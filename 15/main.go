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
	data, e := os.ReadFile("input15")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	sensors := parseInput(data)
	clear := map[int]bool{}
	for _, s := range sensors {
		for _, x := range s.GetCoveredPositionsInRow(2000000) {
			clear[x] = true
		}
	}
	return len(clear)
}

func task2(data string) int {
	sensors := parseInput(data)
	for _, sensor := range sensors {
		for _, point := range sensor.GetBoundary() {
			if point.x < 0 || point.x > 4000000 || point.y < 0 || point.y > 4000000 {
				continue
			}
			isCovered := false
			for _, sensor2 := range sensors {
				if sensor2.location.Distance(point) < sensor2.Radius() {
					isCovered = true
					break
				}
			}
			if !isCovered {
				return point.x*4000000 + point.y
			}
		}
	}
	panic(sensors)
}

func parseInput(data string) []Sensor {
	sensors := []Sensor{}
	for _, line := range getLines(data) {
		sensors = append(sensors, parseSensor(line))
	}
	return sensors
}

func parseSensor(line string) Sensor {
	var sx, sy, bx, by int
	fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
	return Sensor{Point{sx, sy}, Point{bx, by}}
}

type Point struct{ x, y int }

func (p Point) Distance(other Point) int {
	dx := p.x - other.x
	if dx < 0 {
		dx = -dx
	}
	dy := p.y - other.y
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

type Sensor struct {
	location Point
	beacon   Point
}

func (s Sensor) GetCoveredPositionsInRow(n int) []int {
	radius := s.Radius()
	distanceFromN := s.location.Distance(Point{s.location.x, n})

	radiusAtN := radius - distanceFromN
	result := []int{}
	for i := -radiusAtN; i <= radiusAtN; i++ {
		x := s.location.x + i
		if s.beacon.y != n || s.beacon.x != x {
			result = append(result, x)
		}
	}
	return result
}

func (s Sensor) GetBoundary() []Point {
	radius := s.Radius()
	top := Point{s.location.x, s.location.y - radius - 1}
	bottom := Point{s.location.x, s.location.y + radius + 1}
	left := Point{s.location.x - radius - 1, s.location.y}
	right := Point{s.location.x + radius + 1, s.location.y}

	boundary := []Point{}
	for i := 0; i < radius; i++ {
		boundary = append(
			boundary,
			Point{top.x + i, top.y + i},
			Point{bottom.x - i, bottom.y - i},
			Point{left.x + i, left.y - i},
			Point{right.x - i, right.y + i},
		)
	}
	return boundary
}

func (s Sensor) Radius() int {
	return s.location.Distance(s.beacon)
}
