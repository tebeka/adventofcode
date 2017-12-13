from io import StringIO

data = '''\
0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5
'''

graph = {}

# if 1:
#     for line in StringIO(data):
with open('input-12.txt') as fp:
    for line in fp:
        if not line.strip():
            continue

        head, _, *tail = line.split()
        for node in tail:
            node = node.replace(',', '')
            graph.setdefault(head, []).append(node)


def group(graph, start):
    group = set()
    pending = [start]

    while pending:
        node = pending.pop()
        for other in graph[node]:
            if other in group:
                continue
            group.add(other)
            pending.append(other)
    return group


print(len(group(graph, '0')))


def num_groups(graph):
    nodes = set(graph)
    count = 0
    while nodes:
        count += 1
        node = nodes.pop()
        nodes -= group(graph, node)
    return count

print(num_groups(graph))
