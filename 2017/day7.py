import re
from collections import Counter

# puzzle = '''
# pbga (66)
# xhth (57)
# ebii (61)
# havc (66)
# ktlj (57)
# fwft (72) -> ktlj, cntj, xhth
# qoyq (66)
# padx (45) -> pbga, havc, qoyq
# tknk (41) -> ugml, padx, fwft
# jptl (61)
# ugml (68) -> gyxo, ebii, jptl
# gyxo (61)
# cntj (57)
# '''

with open('input-7.txt') as fp:
    puzzle = fp.read()

counts = Counter(re.findall('[a-z]+', puzzle))
print([name for name, count in counts.items() if count == 1][0])

if 0:
    counts = {}

    for line in puzzle.splitlines():
        match = re.search('([a-z]+) \((\d+)\)( -> (.+))?', line)
        if not match:
            continue
        name = match.group(1).strip()
        weight = int(match.group(2))
        nodes = (match.group(4) or '').split(', ')
        if name not in counts:
            counts[name] = 0
        # print(f'name: {name}')
        # print(f'weight: {weight}')
        # print(f'nodes: {nodes}')
        for node in nodes:
            if not node.strip():
                continue
            counts[node] = counts.get(node, 0) + 1

    print(sorted(counts, key=counts.get)[0])

