def show(i, steps):
    print(' '.join(f'({v})' if i == j else str(v) for j, v in enumerate(steps)))


def execute(steps):
    steps = steps[:]
    count = 0
    i = 0

    while True:
        count += 1
        # show(i, steps)
        val = steps[i]
        steps[i] += 1
        i += val
        if i < 0 or i >= len(steps):
            return count


def execute2(steps):
    steps = steps[:]
    count = 0
    i = 0

    while True:
        count += 1
        # show(i, steps)
        val = steps[i]
        if val >= 3:
            steps[i] -= 1
        else:
            steps[i] += 1
        i += val
        if i < 0 or i >= len(steps):
            return count


with open('input-5.txt') as fp:
    steps = [int(line) for line in fp]
n = execute(steps)
print(n)
n = execute2(steps)
print(n)
