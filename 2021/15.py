from time import monotonic
from functools import lru_cache


def load_mat(file_name):
    mat = []
    with open(file_name) as fp:
        for line in fp:
            row = tuple(int(c) for c in line.strip())
            mat.append(row)
    return tuple(mat)


# file_name = '15_small.txt'
file_name = '15.txt'
mat = load_mat(file_name)
# print(mat)


@lru_cache
def shortest_path(mat, x, y):
    if x == len(mat)-1 and y == len(mat[0])-1:
        return mat[x][y]

    px, py = float('inf'), float('inf')

    if x < len(mat)-1:
        px = shortest_path(mat, x+1, y)

    if y < len(mat[0])-1:
        py = shortest_path(mat, x, y+1)

    init = mat[x][y] if (x, y) != (0, 0) else 0
    return init + min(px, py)


start = monotonic()
p = shortest_path(mat, 0, 0)
duration = monotonic() - start
print(p, duration)
