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
	data, e := os.ReadFile("input10")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	cpu := makeCPU()
	for _, line := range getLines(data) {
		cpu.exec(line)
	}
	return (cpu.signalStrength(20) +
		cpu.signalStrength(60) +
		cpu.signalStrength(100) +
		cpu.signalStrength(140) +
		cpu.signalStrength(180) +
		cpu.signalStrength(220))
}

func task2(data string) int {
	cpu := makeCPU()
	for _, line := range getLines(data) {
		cpu.exec(line)
	}
	cpu.drawCRT()
	return -1
}

type CPU struct {
	x []int
}

func makeCPU() CPU {
	return CPU{[]int{1}}
}

func (c *CPU) nop() {
	x := c.x[len(c.x)-1]
	c.x = append(c.x, x)
}

func (c *CPU) add(v int) {
	x := c.x[len(c.x)-1]
	c.x = append(c.x, x, x+v)
}

func (c *CPU) exec(line string) {
	parts := strings.Split(line, " ")
	switch parts[0] {
	case "noop":
		c.nop()
	case "addx":
		c.add(atoi(parts[1]))
	}
}

func (c *CPU) signalStrength(cycle int) int {
	return cycle * c.x[cycle-1]
}

func (c *CPU) drawCRT() {
	cycles := len(c.x)
	for i := 0; i < cycles; i += 40 {
		for j := 0; j < 40; j++ {
			x := c.x[i+j]
			pixel := "."
			if j-1 <= x && x <= j+1 {
				pixel = "#"
			}
			fmt.Print(pixel)
		}
		fmt.Println()
	}
}
