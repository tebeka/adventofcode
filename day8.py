import operator
from collections import defaultdict

ops = {
    'inc': operator.add,
    'dec': operator.sub,
}

regs = defaultdict(int)
max_reg = 0

with open('input-8.txt') as fp:
    for line in fp:
        reg, op, val, _, creg, cop, cval = line.split()
        assert op in ops, f'unknown op - {op}'

        if not eval(f'{creg} {cop} {cval}', None, regs):
            continue

        regs[reg] = ops[op](regs[reg], int(val))
        if regs[reg] > max_reg:
            max_reg = regs[reg]

print(max(regs.values()))
print(max_reg)
