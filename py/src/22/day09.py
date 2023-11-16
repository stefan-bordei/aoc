from ..common.util import get_puzzle_input_str, print_results


CONTENTS = [(x[0], int(x[2:])) for x in get_puzzle_input_str('22', 'day09', '\n\n')[0]]


# Mappings
move = {
    "R": (1, 0),
    "L": (-1, 0),
    "U": (0, 1),
    "D": (0, -1)
}


def sign(x):
    if x < 0:
        return -1
    elif x > 0: 
        return 1
    return 0


def step(x, y):
    return max(abs(x[0]-y[0]), abs(x[1]-y[1]))


# Solution Part1
def part1(data, rope_len):
    res = set()
    rope = []
    base = (0, 0)
    
    for i in range(rope_len):
        rope.append(base[:])

    for i, j in data:
        for _ in range(j):
            rope[0] = [sum(tup) for tup in zip(rope[0][:], move[i])]
            for c in range(len(rope)-1):
                if abs(step(rope[c], rope[c+1])) > 1:
                    rope[c+1] = [
                        sum(tup) for tup in zip(
                            rope[c+1], 
                            (
                                sign(rope[c][:][0] - rope[c+1][:][0]), 
                                sign(rope[c][:][1] - rope[c+1][:][1])
                            )
                        )
                    ]
                else:
                    break
            res.add(tuple(rope[-1]))
    if rope_len == 2:
        return len(set(res))
    return len(set(res))


# Solution Part2
def part2(data, rope_len):
    return part1(data, rope_len)
            

def solve():
    print_results(part1(CONTENTS, 2), part2(CONTENTS, 10))

