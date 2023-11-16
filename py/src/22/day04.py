from ..common.util import get_input, print_results


PUZZLE_INPUT = [[[int(boundry) for boundry in zone.split('-')] for zone in item.split(',')] for item in get_input('22', 'day04')]


def part1(data):
    total = 0
    for rng in data:
        if (rng[0][0] >= rng[-1][0] and rng[0][-1] <= rng[-1][-1]) or\
            (rng[-1][0] >= rng[0][0] and rng[-1][-1] <= rng[0][-1]):
            total += 1
    return total


def part2(data):
    total = 0
    for rng in data:
        first, second = list(range(rng[0][0], rng[0][-1] + 1)), list(range(rng[-1][0], rng[-1][-1] + 1))
        total += 1 if [i for i in first if i in second] else 0
    return total


def solve():
    print_results(part1(PUZZLE_INPUT), part2(PUZZLE_INPUT))

