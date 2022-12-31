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
	data, e := os.ReadFile("input18")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	return parseInput(data).surfaceArea()
}

func task2(data string) int {
	return parseInput(data).exteriorSurfaceArea()
}

func parseInput(data string) Blocks {
	blocks := make(Blocks)
	for _, line := range getLines(data) {
		xyz := strings.Split(line, ",")
		block := Block{atoi(xyz[0]), atoi(xyz[1]), atoi(xyz[2])}
		blocks[block] = true
	}
	return blocks
}

type Block struct{ x, y, z int }

func (b Block) adjacent() []Block {
	blocks := []Block{
		Block{b.x + 1, b.y, b.z},
		Block{b.x - 1, b.y, b.z},
		Block{b.x, b.y + 1, b.z},
		Block{b.x, b.y - 1, b.z},
		Block{b.x, b.y, b.z + 1},
		Block{b.x, b.y, b.z - 1},
	}
	return blocks
}

func (b Block) inBounds(x1, x2, y1, y2, z1, z2 int) bool {
	return (b.x >= x1 && b.x <= x2 &&
		b.y >= y1 && b.y <= y2 &&
		b.z >= z1 && b.z <= z2)
}

type Blocks map[Block]bool

func (b Blocks) surfaceArea() int {
	area := 0
	for block := range b {
		for _, adjacent := range block.adjacent() {
			if _, exist := b[adjacent]; !exist {
				area += 1
			}
		}
	}
	return area
}

func (b Blocks) exteriorSurfaceArea() int {
	exteriorBlocks := b.getInterestingExteriorBlocks()

	area := 0
	for block := range b {
		for _, adjacent := range block.adjacent() {
			if _, exist := exteriorBlocks[adjacent]; exist {
				area += 1
			}
		}
	}
	return area
}

func (b Blocks) xBounds() (int, int) {
	min, max := 0, 0
	for block := range b {
		if min > block.x {
			min = block.x
		}
		if max < block.x {
			max = block.x
		}
	}
	return min - 1, max + 1
}

func (b Blocks) yBounds() (int, int) {
	min, max := 0, 0
	for block := range b {
		if min > block.y {
			min = block.y
		}
		if max < block.y {
			max = block.y
		}
	}
	return min - 1, max + 1
}

func (b Blocks) zBounds() (int, int) {
	min, max := 0, 0
	for block := range b {
		if min > block.z {
			min = block.z
		}
		if max < block.z {
			max = block.z
		}
	}
	return min - 1, max + 1
}

func (b Blocks) getInterestingExteriorBlocks() Blocks {
	x1, x2 := b.xBounds()
	y1, y2 := b.yBounds()
	z1, z2 := b.zBounds()
	initialExteriorBlock := Block{x1, y1, z1}

	seen := Blocks{}
	interestingExteriorBlocks := Blocks{}
	queue := []Block{initialExteriorBlock}
	for len(queue) > 0 {
		block := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if _, exist := seen[block]; exist {
			continue
		}
		seen[block] = true

		for _, adj := range block.adjacent() {
			if !adj.inBounds(x1, x2, y1, y2, z1, z2) {
				continue
			}
			if _, exist := b[adj]; !exist {
				queue = append(queue, adj)
			} else {
				interestingExteriorBlocks[block] = true
			}
		}
	}
	return interestingExteriorBlocks
}
