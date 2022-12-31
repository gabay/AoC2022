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

func getLines(s string) []string {
	return strings.Split(strings.Trim(s, "\n"), "\n")
}

func atoi(s string) int {
	result, ok := strconv.Atoi(s)
	check(ok)
	return result
}

func atoiList(s []string) []int {
	ints := []int{}
	for _, ss := range s {
		ints = append(ints, atoi(ss))
	}
	return ints
}

func findAll(regexString string, data string) []string {
	re, ok := regexp.Compile(regexString)
	check(ok)
	return re.FindAllString(data, -1)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	data, e := os.ReadFile("input19")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	blueprints := parseInput(data)
	qualityLevel := 0
	for i, blueprint := range blueprints {
		qualityLevel += (i + 1) * blueprint.GetMaxGeode(24)
	}
	return qualityLevel
}

func task2(data string) int {
	blueprints := parseInput(data)[:3]
	out := make(chan int)

	go blueprints[0].GetMaxGeode2(32, out)
	go blueprints[1].GetMaxGeode2(32, out)
	go blueprints[2].GetMaxGeode2(32, out)
	return <-out * <-out * <-out
}

func parseInput(data string) []Blueprint {
	blueprints := []Blueprint{}
	for _, line := range getLines(data) {
		numbers := atoiList(findAll("\\d+", line))
		oreCost := Cost{numbers[1], 0, 0}
		clayCost := Cost{numbers[2], 0, 0}
		obsidianCost := Cost{numbers[3], numbers[4], 0}
		geodeCost := Cost{numbers[5], 0, numbers[6]}
		blueprint := Blueprint{oreCost, clayCost, obsidianCost, geodeCost}
		blueprints = append(blueprints, blueprint)
	}
	return blueprints
}

type Cost struct{ ore, clay, obsidian int }

type Blueprint struct {
	OreCost      Cost
	clayCost     Cost
	obsidianCost Cost
	geodeCost    Cost
}

type Bots struct{ ore, clay, obsidian, geode int }
type Resources struct{ ore, clay, obsidian, geode int }

func initialBots() Bots {
	return Bots{1, 0, 0, 0}
}
func InitialResource() Resources {
	return Resources{0, 0, 0, 0}
}

func (r *Resources) canPay(c Cost) bool {
	return (r.ore >= c.ore && r.clay >= c.clay && r.obsidian >= c.obsidian)
}

func (r *Resources) step(b Bots, c Cost) {
	r.ore += b.ore - c.ore
	r.clay += b.clay - c.clay
	r.obsidian += b.obsidian - c.obsidian
	r.geode += b.geode
}

type State struct {
	minutes   int
	bots      Bots
	resources Resources
}

func (s *State) next(c Cost) {
	s.minutes -= 1
	s.resources.step(s.bots, c)
}

func (b Blueprint) GetMaxGeode(minutes int) int {
	initialState := State{minutes, initialBots(), InitialResource()}
	states := []State{initialState}

	maxGeode := 0
	for len(states) > 0 {
		state := states[len(states)-1]
		states = states[:len(states)-1]

		if state.minutes == 0 {
			maxGeode = max(maxGeode, state.resources.geode)
			continue
		}

		states = append(states, b.nextStates(state)...)
	}
	return maxGeode
}

func (b Blueprint) nextStates(s State) []State {
	nextStates := []State{}

	ss := s
	if s.bots.ore < 4 {
		for ss.minutes > 0 && !ss.resources.canPay(b.OreCost) {
			ss.next(Cost{})
		}
		if ss.minutes > 0 {
			ss.next(b.OreCost)
			ss.bots.ore += 1
			nextStates = append(nextStates, ss)
		}
	}

	ss = s
	for ss.minutes > 0 && !ss.resources.canPay(b.clayCost) {
		ss.next(Cost{})
	}
	if ss.minutes > 0 {
		ss.next(b.clayCost)
		ss.bots.clay += 1
		nextStates = append(nextStates, ss)
	}

	ss = s
	for ss.minutes > 0 && !ss.resources.canPay(b.obsidianCost) {
		ss.next(Cost{})
	}
	if ss.minutes > 0 {
		ss.next(b.obsidianCost)
		ss.bots.obsidian += 1
		nextStates = append(nextStates, ss)
	}

	ss = s
	for ss.minutes > 0 && !ss.resources.canPay(b.geodeCost) {
		ss.next(Cost{})
	}
	if ss.minutes > 0 {
		ss.next(b.geodeCost)
		ss.bots.geode += 1
		nextStates = append(nextStates, ss)
	}

	// do nothing and count geode
	geode := s.resources.geode + (s.minutes * s.bots.geode)
	ss = State{0, s.bots, Resources{0, 0, 0, geode}}
	nextStates = append(nextStates, ss)

	return nextStates
}

func (b Blueprint) GetMaxGeode2(minutes int, out chan int) {
	result := b.GetMaxGeode(minutes)
	out <- result
}
