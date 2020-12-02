from itertools import permutations

with open('day-1.txt') as fp:
    nums = [int(line.strip()) for line in fp]

# nums = [1721, 979, 366, 299, 675, 1456]

for v1, v2, v3 in permutations(nums, 3):
    if v1 + v2 + v3 == 2020:
        print(v1 * v2 * v3)
        break
