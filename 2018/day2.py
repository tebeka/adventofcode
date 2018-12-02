from collections import Counter


def checksum(fname):
    n2, n3 = 0, 0
    with open(fname) as fp:
        for line in fp:
            counts = set(Counter(line).values())
            n2 += 1 if 2 in counts else 0
            n3 += 1 if 3 in counts else 0
    return n2 * n3


def step_1():
    print(checksum('day-2.txt'))


def diff_locs(s1, s2):
    return list(i for i, (c1, c2) in enumerate(zip(s1, s2)) if c1 != c2)


def step_2():
    with open('day-2.txt') as fp:
        ids = list(line.strip() for line in fp)

    for i, id1 in enumerate(ids):
        for id2 in ids[i+1:]:
            locs = diff_locs(id1, id2)
            if len(locs) == 1:
                idx = locs[0]
                print(id1[:idx] + id1[idx+1:])
                return


if __name__ == '__main__':
    step_1()
    step_2()
