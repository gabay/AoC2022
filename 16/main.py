import itertools
import os
import re
import sys
from collections import deque
from dataclasses import dataclass


def main():
	valves = parse_input(open('input16').read())
	distances = get_all_distances(valves)

	print('PART 1:')
	max_flow = get_max_flow('AA', 30, valves, distances, get_interesting(valves))
	print(max_flow)

	print('PART 2:')
	max_flow = get_max_flow2('AA', 26, valves, distances, get_interesting(valves))
	print(max_flow)


@dataclass
class Valve:
	flow: int
	adjacent: list[str]

def parse_input(data: str) -> dict:
	valves = {}
	for line in data.splitlines():
		name, flow_str, adjacent_str = re.fullmatch(r'Valve (\w{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)', line).groups()
		flow = int(flow_str)
		adjacent = adjacent_str.split(', ')
		valves[name] = Valve(flow, adjacent)
	return valves

def get_all_distances(valves: dict[str, Valve]) -> dict:
	distances = {}
	for src in valves:
		distances[src] = get_distances(valves, src)
	return distances

def get_distances(valves: dict[str, Valve], src: str) -> dict:
	distances = {src: 0}
	queue = deque([src])
	while queue:
		item = queue.popleft()
		for adj in valves[item].adjacent:
			if adj not in distances:
				distances[adj] = distances[item] + 1
				queue.append(adj)
	return distances

def get_interesting(valves: dict[str, Valve]) -> set[str]:
	return {name for name, valve in valves.items() if valve.flow > 0}

def get_max_flow(curr: str, time_left: int, valves: dict[str, Valve], distances: dict, interesting: set) -> int:
	max_flow = 0
	for next in interesting:
		next_time_left = time_left - distances[curr][next] - 1
		if next_time_left > 0:
			next_flow = valves[next].flow
			next_total_flow = next_flow * next_time_left

			valves[next].flow = 0
			next_max_flow = get_max_flow(next, next_time_left, valves, distances, interesting - {next})
			valves[next].flow = next_flow

			if max_flow < next_total_flow + next_max_flow:
				max_flow = next_total_flow + next_max_flow

	return max_flow

def get_max_flow2(curr: str, time_left: int, valves: dict[str, Valve], distances: dict, interesting: set) -> int:
	max_flow = 0
	for interesting_a in itertools.combinations(interesting, len(interesting) // 2):
		interesting_a = set(interesting_a)
		interesting_b = interesting - interesting_a
		max_flow_a = get_max_flow(curr, time_left, valves, distances, interesting_a)
		max_flow_b = get_max_flow(curr, time_left, valves, distances, interesting_b)
		if max_flow < max_flow_a + max_flow_b:
			max_flow = max_flow_a + max_flow_b

	return max_flow

if __name__ == '__main__':
	main()
