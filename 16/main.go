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
	data, e := os.ReadFile("input16")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	state := parseInput(data, 30)
	maxPressureReleased := state.GetMaxPressureReleased()
	return maxPressureReleased
}

func task2(data string) int {
	// state := parseInput(data, 26)
	// maxPressureReleased := state.GetMaxPressureReleased()
	// return maxPressureReleased
	return 0
}

func parseInput(data string, time int) State {
	valves := make(map[string]Valve)
	for _, line := range getLines(data) {
		valve := parseValve(line)
		valves[valve.name] = valve
	}
	interesting := map[string]bool{}
	for name, valve := range valves {
		interesting[name] = valve.IsInteresting()
	}
	return State{time, 0, "AA", "AA", 0, 0, valves, GetDistances(valves), interesting}
}

func parseValve(line string) Valve {
	var name string
	var flow int
	fmt.Sscanf(line, "Valve %s has flow rate=%d", &name, &flow)
	adjacentString := strings.SplitN(line, " ", 10)[9]
	adjacent := strings.Split(adjacentString, ", ")
	return Valve{name, flow, adjacent}
}

type Valve struct {
	name     string
	flow     int
	adjacent []string
}

func (v Valve) IsInteresting() bool {
	return v.flow > 0
}

type State struct {
	timeLeft         int
	pressureReleased int
	current1         string
	current2         string
	readytAt1        int
	readyAt2         int
	valves           map[string]Valve
	distances        map[string]map[string]int
	interesting      map[string]bool
}

func (s State) GetMoves() []State {
	moves := []State{}

	for name, isInteresting := range s.interesting {
		distance := s.distances[s.current1][name]
		valve := s.valves[name]
		if !isInteresting || s.timeLeft <= distance {
			continue
		}

		newTimeLeft := s.timeLeft - distance - 1
		newPressureReleased := s.pressureReleased + (valve.flow * newTimeLeft)
		newValves := Copy(s.valves)
		newValves[valve.name] = valve
		newIntersting := Copy2(s.interesting)
		newIntersting[valve.name] = false
		newState := State{newTimeLeft, newPressureReleased, valve.name, s.current2, 0, 0,
			newValves, s.distances, newIntersting}
		moves = append(moves, newState)
	}
	return moves
}

func (s State) GetMaxPressureReleased() int {
	maxFlow := s.pressureReleased

	if s.timeLeft > 1 && len(s.interesting) > 0 {
		for _, nextState := range s.GetMoves() {
			stateFlow := nextState.GetMaxPressureReleased()
			if maxFlow < stateFlow {
				maxFlow = stateFlow
			}
		}
	}

	return maxFlow
}

func (s State) GetMoves2() []State {
	moves := []State{}

	for name, distance := range s.distances[s.current1] {
		valve := s.valves[name]
		if valve.flow == 0 || s.timeLeft <= distance {
			continue
		}

		newTimeLeft := s.timeLeft - distance - 1
		newPressureReleased := s.pressureReleased + (valve.flow * newTimeLeft)
		newValves := Copy(s.valves)
		newValves[valve.name] = valve
		newIntersting := Copy2(s.interesting)
		newIntersting[valve.name] = false
		newState := State{newTimeLeft, newPressureReleased, valve.name, s.current2, 0, 0,
			newValves, s.distances, newIntersting}
		moves = append(moves, newState)
	}
	return moves
}

func (s State) GetMaxPressureReleased2() int {
	maxFlow := s.pressureReleased

	if s.timeLeft > 1 && len(s.interesting) > 0 {
		for _, nextState := range s.GetMoves2() {
			stateFlow := nextState.GetMaxPressureReleased2()
			if maxFlow < stateFlow {
				maxFlow = stateFlow
			}
		}
	}

	return maxFlow
}

func Copy(m map[string]Valve) map[string]Valve {
	newMap := make(map[string]Valve)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func Copy2(m map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func GetDistances(valves map[string]Valve) map[string]map[string]int {
	distances := map[string]map[string]int{}
	for name := range valves {
		distances[name] = GetDistancesFromValve(valves, name)
	}
	return distances
}

func GetDistancesFromValve(valves map[string]Valve, sourceName string) map[string]int {
	queue := make(chan string, len(valves))
	queue <- sourceName
	distances := map[string]int{sourceName: 0}
	for len(queue) > 0 {
		name := <-queue
		for _, adjacent := range valves[name].adjacent {
			if _, exist := distances[adjacent]; !exist {
				queue <- adjacent
				distances[adjacent] = distances[name] + 1
			}
		}
	}
	return distances
}
