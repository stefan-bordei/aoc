from ..common.util import get_puzzle_input_str, print_results
from string import ascii_lowercase
from collections import deque


CONTENTS = [[char for char in l] for l in [line[0] for line in get_puzzle_input_str('22', 'day12', '\n')]]


def get_start_end(data, n, m):
    for i in range(n):
        for j in range(m):
            if data[i][j] == 'S':
                start = i, j
            if data[i][j] == 'E':
                end = i, j
    return start, end


def get_val(ch):
    if ch in ascii_lowercase:
        return ascii_lowercase.index(ch)
    if ch == 'S':
        return 0
    if ch == 'E':
        return 25


def neighbours(data, i, j, n, m):
    for di, dj in [[1, 0], [-1, 0], [0, 1], [0, -1]]:
        ii = i + di
        jj = j + dj
        if not (0 <= ii < n and 0 <= jj < m):
            continue
        if get_val(data[ii][jj]) <= get_val(data[i][j])+1:
            yield ii, jj


def neighbours_rev(data, i, j, n, m):
    for di, dj in [[1, 0], [-1, 0], [0, 1], [0, -1]]:
        ii = i + di
        jj = j + dj
        if not (0 <= ii < n and 0 <= jj < m):
            continue
        if get_val(data[ii][jj]) >= get_val(data[i][j])-1:
            yield ii, jj


def part1(data):
    n = len(data)
    m = len(data[0])
    
    start, end = get_start_end(data, n, m)
    visited = [[False] * m for _ in range(n)]
    
    mapping = deque()
    mapping.append((0, start[0], start[1]))
    
    while len(mapping) > 0:
        steps, cx, cy = mapping.popleft()
        if visited[cx][cy]:
            continue
        visited[cx][cy] = True

        if (cx, cy) == end:
            return steps
        for ii, jj in neighbours(data, cx, cy, n, m):
            mapping.append((steps+1, ii, jj))
    return 0


def part2(data):
    n = len(data)
    m = len(data[0])
    
    start, end = get_start_end(data, n, m)
    visited = [[False] * m for _ in range(n)]
    
    mapping = deque()
    mapping.append((0, end[0], end[1]))
    
    while len(mapping) > 0:
        steps, cx, cy = mapping.popleft()
        if visited[cx][cy]:
            continue
        visited[cx][cy] = True

        if data[cx][cy] == 'a':
            return steps
        for ii, jj in neighbours_rev(data, cx, cy, n, m):
            mapping.append((steps+1, ii, jj))
    return 0


def solve():
    print_results(part1(CONTENTS), part2(CONTENTS))

