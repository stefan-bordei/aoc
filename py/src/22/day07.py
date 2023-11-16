from ..common.util import get_input, print_results
 

CONTENTS = get_input('22', 'day07')[1:]


class Dir:
    def __init__(self, name, prev=None, size=0):
        self.name = name
        self.prev = prev
        self.size = size
        self.contents = {}
        self.contents['dirs'] = {}
        self.contents['files'] = {}
    
    def __repr__(self):
        return str(self.contents)

    def update_size(self, val):
        self.size += val
        if self.prev != None:
            self.prev.update_size(val)
    
    def get_all_with_limit(self, max):
        res = 0
        if self.size <= max:
            res += self.size

        if len(self.contents['dirs'].values()) > 0:
            for v in self.contents['dirs'].values():
                res += v.get_all_with_limit(max)
        return res

    def get_min_to_remove(self, min, available):
        marked = []
        if available + self.size >= min:
            marked.append(self.size)
        if len(self.contents['dirs'].values()) > 0:
            for v in self.contents['dirs'].values():
                marked += v.get_min_to_remove(min, available)
        return list(marked)
        

def generate_file_tree(data):
    root = Dir('/', None)
    current = root
    for line in data:
        first, second = line.split()[0], line.split()[1]
        if first == '$':
            if second == 'ls':
                continue
            if second == 'cd':
                if len(line.split()) == 2:
                    current = root
                    continue
                if len(line.split()) == 3:
                    if line.split()[-1] == '..':
                        current = current.prev
                        continue
                    else:
                        current = current.contents['dirs'][line.split()[-1]]
                        continue
        if first.isnumeric():
            current.contents['files'] = {'size': int(first), 'cur': current.name}
            current.update_size(int(first))
            continue
        if first == 'dir':
            new = Dir(second, current)
            current.contents['dirs'][second] = new
    return root


def part1(obj, limit):
    return obj.get_all_with_limit(limit)


def part2(obj, min, available):
    res = obj.get_min_to_remove(min, available)
    res.sort()
    return res[0]


def solve():
    file_contents = generate_file_tree(CONTENTS)
    print_results(part1(file_contents, 100000), part2(file_contents, 30000000, 70000000-file_contents.size))

