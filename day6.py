def mem2key(mem):
    return ','.join(str(v) for v in mem)


def find_max(mem):
    imax, vmax = 0, mem[0]
    for i, v in enumerate(mem[1:], 1):
        if v > vmax:
            imax, vmax = i, v
    return imax, vmax


# mem = [0, 2, 7, 0]
mem = [14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4]
seen = set([mem2key(mem)])
count = 1

while True:
    i, v = find_max(mem)
    mem[i] = 0

    while v:
        i = (i+1) % len(mem)
        mem[i] += 1
        v -= 1

    key = mem2key(mem)
    if key in seen:
        print(count)
        break

    seen.add(key)
    count += 1
