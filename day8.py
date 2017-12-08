import operator
from collections import defaultdict

code = '''
b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10
'''

ops = {
    'inc': operator.add,
    'dec': operator.sub,
}

regs = defaultdict(int)
max_reg = 0

# for line in code.splitlines():
with open('input-8.txt') as fp:
    for line in fp:
        if not line.strip():
            continue

        fields = line.split()
        reg, op, val, _, creg, cop, cval = fields
        assert op in ops, f'unknown op - {op}'

        if not eval(f'{creg} {cop} {cval}', None, regs):
            continue

        regs[reg] = ops[op](regs[reg], int(val))
        if regs[reg] > max_reg:
            max_reg = regs[reg]

print(max(regs.values()))
print(max_reg)
