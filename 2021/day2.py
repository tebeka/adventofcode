

# file_name = '2_small.txt'
file_name = '2.txt'

# part 1
if 0:
    loc, depth = 0, 0
    with open(file_name) as fp:
        for line in fp:
            cmd, n = line.split()
            n = int(n)
            if cmd == 'forward':
                loc += n
            elif cmd == 'up':
                depth -= n
            elif cmd == 'down':
                depth += n

aim, depth, loc = 0, 0, 0
with open(file_name) as fp:
    for line in fp:
        cmd, n = line.split()
        n = int(n)
        if cmd == 'down':
            aim += n
        elif cmd == 'up':
            aim -= n
        elif cmd == 'forward':
            loc += n
            depth += aim * n


print(loc * depth)
