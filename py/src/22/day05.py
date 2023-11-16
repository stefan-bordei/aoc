from ..common.util import get_puzzle_input_str, print_results
import copy


unprocessed_data = get_puzzle_input_str('22', 'day05', '\n\n')


def get_crate_stacks(crate_data):
    flattened = [
        item for sublist in [
            [
                (item[i], i//4+1) for i, c in enumerate(item) if c.isalpha()
            ] for item in crate_data
        ] for item in sublist
    ]
    
    res = {}
    # Not optimal but my brain hurts....
    for i in list(range(9)):
        res[i+1] = []
        for item in flattened:
            if item[1] == i+1:
                res[i+1].append(item[0])
    return res


PUZZLE_INPUT = dict(
    zip(
        ['crates', 'operations'],
        [
            get_crate_stacks(unprocessed_data[0][:-1]),
            [
                dict(
                    zip(
                        line.split()[::2],
                        [int(i) for i in line.split()[1::2]]
                    )
                ) for line in unprocessed_data[1:][0]
            ]
        ]
    )
)


PART1 = copy.deepcopy(PUZZLE_INPUT)
PART2 = copy.deepcopy(PUZZLE_INPUT)


def part1(data):
    stack = data.get('crates')
    for move in data.get('operations'):
        for _ in list(range(move.get('move'))):
            stack[move.get('to')].insert(0, stack[move.get('from')][0])
            stack[move.get('from')] = stack[move.get('from')][1:]
    return ''.join([crates[0] for crates in stack.values()])


def part2(data):
    stack = data.get('crates')
    for move in data.get('operations'):
        if move.get('move') == 1:
            stack[move.get('to')].insert(0, stack[move.get('from')][0])
            stack[move.get('from')] = stack[move.get('from')][1:]
        else:
            temp = stack[move.get('from')].copy()
            stack[move.get('to')] = temp[:move.get('move')] + stack[move.get('to')]
            stack[move.get('from')] = stack.get(move.get('from'))[move.get('move'):] 
    return ''.join([crates[0] for crates in stack.values()])


def solve():
    print_results(part1(PART1), part2(PART2))

