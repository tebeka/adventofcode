# file_name = '3_small.txt'
# size = 5
size = 12
file_name = '3.txt'
counts = [{} for _ in range(size)]
with open(file_name) as fp:
    for line in fp:
        for i, b in enumerate(line.strip()):
            counts[i][b] = counts[i].get(b, 0) + 1

bits = [0] * size
for i in reversed(range(size)):
    b = max(counts[i], key=counts[i].get)
    bits[i] = b
gamma = int(''.join(bits), 2)
mask = int(''.join(['1' * size]), 2)
eps = (~gamma) & mask
print(eps * gamma)
