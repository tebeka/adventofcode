def height(n):
    return 1 + 2*n


def outline(n):
    return 4 * height(n) - 4


def last(n):
    return 1 + sum(outline(i) for i in range(1, n+1))


def first(n):
    return last(n-1) + 1


def find_last(n):
    i = 1
    while True:
        if last(i) >= n:
            return i
        i += 1
