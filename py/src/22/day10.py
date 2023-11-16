from ..common.util import get_puzzle_input_str, print_results


CONTENTS = [x for nested in get_puzzle_input_str('22', 'day10', '\n') for x in nested]


def print_msg(msg):
    w = 40
    h = 6

    for i in range(h):
        print(
            msg[i*w:i*w+w-2]
        )


# Solution Part1
def part1(data):
    res, rep, buffer, register, check, msg = 0, 0, 0, 1, False, ''
    while buffer < len(data):
        rep +=1
        if rep in [20, 60, 100, 140, 180, 220]:
            res += rep * register

        if (rep-1) % 40 in range(register-1, register+2):
            msg += '#'
        else:
            msg += ' '
        
        if data[buffer] == 'noop':
            buffer += 1
        else:
            if check:
                register += int(data[buffer].split()[-1])
                check = False
                buffer +=1
            else:
                check = True
    return res, msg

    
# Solution Part2
def part2(data):
    print_msg(part1(data)[1])
    return ' '
            

def solve():
    print_results(part1(CONTENTS)[0], part2(CONTENTS))

