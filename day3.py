'''
37   36  35  34  33  32 31
38   17  16  15  14  13 30
39   18   5   4   3  12 29
40   19   6   1   2  11 28
41   20   7   8   9  10 27
42   21  22  23  24  25 26
43   44  45  46  47  48 49
'''

moves = [
    (-1, 0),
    (0, -1),
    (1, 0),
    (0, 1),
]


def print_grid(grid):
    for row in grid:
        print(' '.join(f'{i:02d}' for i in row))


def make_grid(r):
    size = 1 + 2 * r
    row = col = size // 2

    grid = [[0] * size for _ in range(size)]

    ring = 1
    i = 1
    grid[row][col] = i

    while i < size*size:
        col += 1
        i += 1
        grid[row][col] = i
        for d, (drow, dcol) in enumerate(moves):
            n = ring * 2
            if d == 0:
                n -= 1
            while n > 0:
                i += 1
                row += drow
                col += dcol
                grid[row][col] = i
                n -= 1
        ring += 1
    return grid


def distance(n):
    if n == 1:
        return 0

    r = 1
    while r*r < n:
        r += 1

    size = 1 + 2 * r
    row = col = size // 2

    grid = [[0] * size for _ in range(size)]

    ring = 1
    i = 1
    grid[row][col] = i

    while i < n:
        col += 1
        i += 1
        grid[row][col] = i
        for d, (drow, dcol) in enumerate(moves):
            m = ring * 2
            if d == 0:
                m -= 1
            while m > 0 and i < n:
                i += 1
                row += drow
                col += dcol
                grid[row][col] = i
                m -= 1
        ring += 1

    return abs(row - size // 2) + abs(col - size//2)



print(distance(368078))
