from ..common.util import get_puzzle_input_str, sign, print_results
from copy import deepcopy


CONTENTS = [[[int(x) for x in coord.strip().split(',')] for coord in line[0].split('->')] for line in get_puzzle_input_str('22', 'day14', '\n')]
BOTTOM = []


def map_rock(map, coord):
    start, end = coord[0], coord[-1]
    dx = sign(end[0] - start[0])
    dy = sign(end[1] - start[1])

    point = start
    map.add(tuple(point))
    while point != end:
        point[0] += dx
        point[1] += dy
        BOTTOM.append(point[1])
        map.add(tuple(point))


def map_cave(data):
    cave = set()
    for step in data:
        for idx in range(len(step)-1):
            map_rock(cave, step[idx:idx+2])
    return cave

def drop_sand(map, depth, floor=False):
    sand = (500, 0)
    if sand in map and floor:
        return None
    while True:
        next = fall(map, sand)
        if not next:
            break
        if next[1] > depth and not floor:
            return None
        if next[1] > depth + 1 and floor:
            break
        sand = next
    return sand

def fall(map, pos):
    for delta in [0, -1, 1]:
        fallen = (pos[0] + delta, pos[1] + 1)
        if not fallen in map:
            return (fallen)
    return None


# Solution Part1
def part1(data, depth, floor=False):
    count = 0
    while True:
        next = drop_sand(data, depth, floor)
        if not next:
            break   
        data.add(next)
        count +=1
    return count

    
# Solution Part2
def part2(data, depth, floor=True):
    count = 0
    while True:
        next = drop_sand(data, depth, floor)
        if not next:
            break
        data.add(next)
        count +=1
    return count
            

def solve():
    cave = map_cave(CONTENTS)
    depth = max(BOTTOM)
    print_results(part1(deepcopy(cave), depth), part2(deepcopy(cave), depth))

