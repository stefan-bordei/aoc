
# Helper functions
def get_puzzle_input_num(year, day, separator=None):
    with open(f'./data/{year}/{day}.txt', 'r') as scanned_input:
        if separator:
            prep_input = scanned_input.read().split(separator)
        else:
            prep_input = scanned_input.read()
        return [[int(item) for item in line.split('\n') if item] for line in prep_input if line]


def get_puzzle_input_str(year, day, separator=None):
    with open(f'./data/{year}/{day}.txt', 'r') as scanned_input:
        if separator:
            prep_input = scanned_input.read().split(separator)
        else:
            prep_input = scanned_input.read()
        return [[item for item in line.split('\n') if item] for line in prep_input if line]


def get_input(year, day):
    with open(f'./data/{year}/{day}.txt', 'r') as scanned_input:
        return [line.rstrip() for line in scanned_input.readlines()]


def sign(x):
    if x < 0:
        return -1
    elif x > 0: 
        return 1
    return 0


def print_results(part1, part2):
    print(f'Part 1: {part1}\nPart 2: {part2}')
