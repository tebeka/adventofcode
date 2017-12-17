# size = 5
size = 16

ix = 0
ring = [chr(ord('a') + i) for i in range(size)]

# data = 's1,x3/4,pe/b'
with open('input-16.txt') as fp:
    data = fp.read().strip()

cmds = data.split(',')

seen = {
    ''.join(ring): 0,
}

ntimes = 1_000_000_000

for i in range(ntimes):
    for cmd in cmds:
        op, args = cmd[:1], cmd[1:]
        if op == 's':
            ix = (ix - int(args)) % size
        elif op == 'x':
            a, b = [int(v) for v in args.split('/')]
            a, b = (ix + a) % size, (ix + b) % size
            ring[a], ring[b] = ring[b], ring[a]
        elif op == 'p':
            a, b = args.split('/')
            a, b = ring.index(a), ring.index(b)
            ring[a], ring[b] = ring[b], ring[a]
        else:
            assert False, f'unknown op = {cmd}'
    value = ''.join(ring[ix:] + ring[:ix])
    if i == 0:
        print(value)  # step 1

    if value in seen:
        # TODO: Wrong
        cycle = i - seen[value]
        m = (ntimes - i) % cycle
        val = seen[value] + m
        for k, v in seen.items():
            if v == val:
                print(k)
                break
        break
    seen[value] = i + 1
