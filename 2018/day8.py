import re
from collections import namedtuple

Tree = namedtuple('Tree', 'children meta')


def parse_input(data):
    return (int(v) for v in re.findall(r'\d+', data))


def parse_tree(data):
    n_child = next(data)
    n_meta = next(data)
    children = [parse_tree(data) for _ in range(n_child)]
    meta = [next(data) for _ in range(n_meta)]
    return Tree(children, meta)


def sum_meta(tree):
    return sum(tree.meta) + sum(sum_meta(child) for child in tree.children)


def tree_value(tree):
    if not tree.children:
        return sum(tree.meta)

    mc = (tree.children[i-1] for i in tree.meta if i <= len(tree.children))
    return sum(tree_value(c) for c in mc)


# nums = parse_input("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2")
with open('day-8.txt') as fp:
    nums = parse_input(fp.read())
tree = parse_tree(nums)
print(sum_meta(tree))
print(tree_value(tree))
