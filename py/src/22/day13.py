from ..common.util import get_puzzle_input_str, print_results
from functools import cmp_to_key

CONTENTS = get_puzzle_input_str('22', 'day13', '\n\n')


def compare(left, right):
    if not isinstance(left, list):
        left = [left]
    if not isinstance(right, list):
        right = [right]
    for l, r in zip(left, right):
        if isinstance(l, list) or isinstance(r, list):
            res = compare(l, r)
        else:
            res = r - l
        if res != 0:
            return res
    return len(right) - len(left)


def part1(data):
    res = 0
    for idx, packet in enumerate(data):
        left, right = eval(packet[0]), eval(packet[1])
        if compare(left, right) >0:
            res += idx+1
    return res


def part2(data):
    res = 1
    data_raw = [eval(item) for subitem in data for item in subitem] + [[[2]], [[6]]]
    sorted_data = sorted(data_raw, key=cmp_to_key(compare), reverse=True)
    
    for idx, packet in enumerate(sorted_data, 1):
        if packet in ([[2]], [[6]]):
            res *= idx

    return res


def solve():
    print_results(part1(CONTENTS), part2(CONTENTS))

