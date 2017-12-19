from collections import defaultdict
from io import StringIO

regs = defaultdict(int)

code = '''\
set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2
'''


def valof(n):
    try:
        return int(n)
    except ValueError:
        return regs[n]


prog = []
# with StringIO(code) as fp:
with open('input-18.txt') as fp:
    for line in fp:
        if not line.strip():
            continue
        cmd, x, *y = line.split()
        y = y[0] if y else None
        prog.append((cmd, x, y))


def prn():
    for i, (op, x, y) in enumerate(prog):
        y = y or ''
        prefix = '> ' if i == pc else '  '
        print(f'{prefix}{op} {x:<3} {y:<3} [{regs["a"]:02d}]')
    print('')


pc = 0
freq = None
rec = None

while pc >= 0 and pc < len(prog):
    offset = 1
    cmd, x, y = prog[pc]
    # prn()

    if cmd == 'snd':
        freq = valof(x)
    elif cmd == 'set':
        regs[x] = valof(y)
    elif cmd == 'add':
        regs[x] += valof(y)
    elif cmd == 'mul':
        regs[x] *= valof(y)
    elif cmd == 'mod':
        regs[x] %= valof(y)
    elif cmd == 'rcv':
        if valof(x):
            rec = freq
            break
    elif cmd == 'jgz':
        if valof(x):
            offset = valof(y)
    else:
        assert False, f'{pc}: unknown command: {cmd}'

    pc += offset

print(rec)
