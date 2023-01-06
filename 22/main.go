package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(s string) int {
	result, ok := strconv.Atoi(s)
	check(ok)
	return result
}

func findAll(reString string, s string) []string {
	re, ok := regexp.Compile(reString)
	check(ok)
	return re.FindAllString(s, -1)
}

func main() {
	data, e := os.ReadFile("input22")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	forest := parseInput(data)
	return forest.move(1)
}

func task2(data string) int {
	forest := parseInput(data)
	return forest.move(2)
}

func parseInput(data string) Forest {
	grid, moves, ok := strings.Cut(data, "\n\n")
	if !ok {
		panic(data)
	}

	forest := Forest{strings.Split(grid, "\n"), []Move{}}
	for _, move := range findAll("\\d+|L|R", moves) {
		switch move {
		case "L":
			forest.moves = append(forest.moves, Move{0, 3})
		case "R":
			forest.moves = append(forest.moves, Move{0, 1})
		default:
			forest.moves = append(forest.moves, Move{atoi(move), 0})
		}
	}
	return forest
}

type Forest struct {
	grid  []string
	moves []Move
}

type Move struct {
	steps int
	turn  int
}

type Point struct{ x, y int }

func (f Forest) move(style int) int {
	pos, dir := f.initialPosDir()
	for _, move := range f.moves {
		switch style {
		case 1:
			f.doMove(&pos, &dir, move)
		case 2:
			f.doMove2(&pos, &dir, move)
		default:
			panic(style)
		}

	}
	return (pos.y+1)*1000 + (pos.x+1)*4 + dir
}

func (f Forest) initialPosDir() (Point, int) {
	for x, cell := range f.grid[0] {
		if cell == '.' {
			return Point{x, 0}, 0
		}
	}
	panic(f)
}

func (f Forest) doMove(p *Point, dir *int, move Move) {
	for i := 0; i < move.steps; i++ {
		switch *dir {
		case 0:
			f.moveRight(p)
		case 1:
			f.moveDown(p)
		case 2:
			f.moveLeft(p)
		case 3:
			f.moveUp(p)
		default:
			panic(f)
		}
	}
	*dir = (*dir + move.turn) % 4
}

func (f Forest) moveRight(p *Point) {
	y := p.y
	x := (p.x + 1) % f.widthAt(y)
	for f.isOffMap(x, y) {
		x = (x + 1) % f.widthAt(y)
	}
	if !f.isBlocked(x, y) {
		p.x = x
	}
}

func (f Forest) moveLeft(p *Point) {
	y := p.y
	x := (p.x - 1 + f.widthAt(y)) % f.widthAt(y)
	for f.isOffMap(x, y) {
		x = (x - 1 + f.widthAt(y)) % f.widthAt(y)
	}
	if !f.isBlocked(x, y) {
		p.x = x
	}
}

func (f Forest) moveDown(p *Point) {
	y := (p.y + 1) % f.height()
	x := p.x
	for f.isOffMap(x, y) {
		y = (y + 1) % f.height()
	}
	if !f.isBlocked(x, y) {
		p.y = y
	}
}

func (f Forest) moveUp(p *Point) {
	y := (p.y - 1 + f.height()) % f.height()
	x := p.x
	for f.isOffMap(x, y) {
		y = (y - 1 + f.height()) % f.height()
	}
	if !f.isBlocked(x, y) {
		p.y = y
	}
}

func (f Forest) isBlocked(x, y int) bool {
	return f.grid[y][x] == '#'
}

func (f Forest) isOffMap(x, y int) bool {
	if y < 0 || y >= len(f.grid) || x < 0 || x >= len(f.grid[y]) {
		return true
	}
	return f.grid[y][x] == ' '
}

func (f Forest) widthAt(y int) int {
	return len(f.grid[y])
}
func (f Forest) height() int {
	return len(f.grid)
}

func (f Forest) doMove2(p *Point, dir *int, move Move) {
	// Assuming this layout:
	// _21
	// _3_
	// 54_
	// 6__
	for i := 0; i < move.steps; i++ {
		switch *dir {
		case 0:
			f.moveRight2(p, dir)
		case 1:
			f.moveDown2(p, dir)
		case 2:
			f.moveLeft2(p, dir)
		case 3:
			f.moveUp2(p, dir)
		default:
			panic(f)
		}
	}
	*dir = (*dir + move.turn) % 4
}

func (f Forest) moveRight2(p *Point, dir *int) {
	x := p.x + 1
	y := p.y
	newDir := *dir

	if f.isOffMap(x, y) {
		ymod := y % 50
		switch y / 50 {
		case 0: // 1 right -> 4 right
			x = 99
			y = 149 - ymod
			newDir = 2
		case 1: // 3 right -> 1 bottom
			x = 100 + ymod
			y = 49
			newDir = 3
		case 2: // 4 right -> 1 right
			x = 149
			y = 49 - ymod
			newDir = 2
		case 3: // 6 right -> 4 bottom
			x = 50 + ymod
			y = 149
			newDir = 3
		default:
			panic(y)
		}
		// fmt.Println("R offmap", *p, *dir, "to", x, y, newDir)
	}
	if !f.isBlocked(x, y) {
		*p = Point{x, y}
		*dir = newDir
	}
}

func (f Forest) moveLeft2(p *Point, dir *int) {
	x := p.x - 1
	y := p.y
	newDir := *dir

	if f.isOffMap(x, y) {
		ymod := y % 50
		switch y / 50 {
		case 0: // 2 left -> 5 left
			x = 0
			y = 149 - ymod
			newDir = 0
		case 1: // 3 left -> 5 top
			x = ymod
			y = 100
			newDir = 1
		case 2: // 5 left -> 2 left
			x = 50
			y = 49 - ymod
			newDir = 0
		case 3: // 6 left -> 2 top
			x = 50 + ymod
			y = 0
			newDir = 1
		default:
			panic(y)
		}
	}
	if !f.isBlocked(x, y) {
		*p = Point{x, y}
		*dir = newDir
	}
}

func (f Forest) moveDown2(p *Point, dir *int) {
	x := p.x
	y := p.y + 1
	newDir := *dir

	if f.isOffMap(x, y) {
		xmod := x % 50
		switch x / 50 {
		case 0: // 6 bottom -> 1 top
			x = 100 + xmod
			y = 0
			newDir = 1
		case 1: // 4 bottom -> 6 right
			x = 49
			y = 150 + xmod
			newDir = 2
		case 2: // 1 bottom -> 3 right
			x = 99
			y = 50 + xmod
			newDir = 2
		default:
			panic(x)
		}
	}
	if !f.isBlocked(x, y) {
		*p = Point{x, y}
		*dir = newDir
	}
}

func (f Forest) moveUp2(p *Point, dir *int) {
	x := p.x
	y := p.y - 1
	newDir := *dir

	if f.isOffMap(x, y) {
		xmod := x % 50
		switch x / 50 {
		case 0: // 5 top -> 3 left
			x = 50
			y = 50 + xmod
			newDir = 0
		case 1: // 2 top -> 6 left
			x = 0
			y = 150 + xmod
			newDir = 0
		case 2: // 1 top -> 6 bottom
			x = xmod
			y = 199
			newDir = 3
		default:
			panic(x)
		}
	}
	if !f.isBlocked(x, y) {
		*p = Point{x, y}
		*dir = newDir
	}
}
