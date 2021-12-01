def load(file_name):
    with open(file_name) as fp:
        return [int(line.strip()) for line in fp]


nums = load('1.txt')


def part_one():
    n = sum(1 if j > i else 0 for (i, j) in zip(nums, nums[1:]))
    print(n)


def part_two():
    wins = [sum(nums[i:i+3]) for i in range(len(nums)-2)]
    n = sum(1 if j > i else 0 for (i, j) in zip(wins, wins[1:]))
    print(n)


part_one()
part_two()
