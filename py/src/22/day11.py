from ..common.util import get_puzzle_input_str, print_results
import math
from copy import deepcopy


CONTENTS = get_puzzle_input_str('22', 'day11', '\n\n')


def steal_things(data):
    monkeys = {}
    for idx in range(len(data)):
        monkey = {}
        monkey['items'] = [int(x) for x in data[idx][1].split(':')[-1].rstrip().replace(' ', '').split(',')]
        monkey['operations'] =  {
                'op': data[idx][2].split('old ')[-1].split()[0],
                'val': data[idx][2].split('old ')[-1].split()[-1]
            }
        monkey['test'] = int(data[idx][3].split()[-1])
        monkey['true'] =  int(data[idx][4].split()[-1])
        monkey['false'] = int(data[idx][5].split()[-1])
        monkey['total'] = 0
        monkeys[idx] = monkey
    return monkeys


MONKEYS = steal_things(CONTENTS)

MOD = math.prod([MONKEYS[m]['test'] for m in MONKEYS.keys()])


def part1(data, rounds, worry_reduction):
    for _ in range(rounds):
        for monkey in data.keys():
            current_monkey = data[monkey]
            while len(current_monkey['items']) != 0:
                item = current_monkey['items'].pop(0)
                current_monkey['total'] += 1
                
                if current_monkey['operations']['val'] != 'old':
                    val = int(current_monkey['operations']['val'])
                else:
                    val = item
                worry = eval(f'{item} {current_monkey["operations"]["op"]} {val}')
                if worry_reduction:
                    worry = worry % MOD
                else:
                    worry = worry // 3
                
                if worry % current_monkey['test'] == 0:
                    data[current_monkey['true']]['items'].append(worry)
                else:
                    data[current_monkey['false']]['items'].append(worry)

    res = [data[m]['total'] for m in data.keys()]
    res.sort()
    return math.prod(res[-2:])


def part2(data):
    return 'modded in part1()'


def solve():
    print_results(part1(deepcopy(MONKEYS), 20, False), part1(deepcopy(MONKEYS), 10000, True))

