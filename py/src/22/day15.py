from ..common.util import get_puzzle_input_str, print_results

CONTENTS = [line[0] for line in get_puzzle_input_str('22', 'day15', '\n')]


def parse(data):
    result = []
    for item in data:
        s, b = item.split(':')[0], item.split(':')[-1]
        s, b = list(map(int, s.split('x=')[1].split(', y='))), list(map(int, b.split('x=')[1].split(', y=')))
        result.append({'sensor': s, 'beacon': b})
    return result


# Solution Part1
def part1(data):
    return parse(data)


# Solution Part2
def part2(data):
    return 0


def solve():
    print_results(part1(CONTENTS), part2(CONTENTS))

