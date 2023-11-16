from ..common.util import get_puzzle_input_str, print_results


CONTENTS =  [[int(a) for a in b] for b in [list(i[0]) for i in get_puzzle_input_str('22', 'day08', '\n')]]


# Soulution Part1
def part1(data):
    vis = len(data[0]) * 2 + len(data) * 2 - 4
    for i in range(1, len(data[0])-1):
        for j in range(1, len(data[i])-1):
            if max(data[i][:j]) < data[i][j] or\
                max(data[i][j+1:len(data[0])]) < data[i][j] or\
                max([a[j] for a in data[:i]]) < int(data[i][j]) or\
                max([a[j] for a in data[i+1:len(data)]]) < int(data[i][j]):
                vis += 1
    return vis

    
# Solution Part2
def part2(data):
    # Could not come up with a worst solution even if I tried to....
    def get_vis_rank(lst, check):
        cleaned = []
        for elem in lst:
            if elem <= check:
                cleaned.append(elem)
            if elem == check:
                break
        return len(cleaned)
                
    scores = []
    for i in range(len(data[0])):
        for j in range(len(data[i])):
            top = [a[j] for a in data[:i]]
            top.reverse()
            bottom = [a[j] for a in data[i+1:len(data)]]
            left = data[i][:j]
            left.reverse()
            right = data[i][j+1:len(data[0])]

            scores.append(
                get_vis_rank(top, data[i][j]) *\
                get_vis_rank(bottom, data[i][j]) *\
                get_vis_rank(left, data[i][j]) *\
                get_vis_rank(right, data[i][j])
            )
    return max(scores)
            

def solve():
    print_results(part1(CONTENTS), part2(CONTENTS))

