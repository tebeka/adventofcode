def parse_line(line):
    rule, passwd = [v.strip() for v in line.split(':', 1)]
    counts, letter = rule.split()
    low, high = (int(v) for v in counts.split('-'))

    return low, high, passwd, letter


# 2-5 l: fllxf
def is_valid(line):
    low, high, passwd, letter = parse_line(line)
    return low <= passwd.count(letter) <= high


with open('day-2.txt') as fp:
    print(sum(is_valid(line) for line in fp))


def is_valid2(line):
    low, high, passwd, letter = parse_line(line)
    return (passwd[low-1] == letter) + (passwd[high-1] == letter) == 1


# lines = '''\
# 1-3 a: abcde
# 1-3 b: cdefg
# 2-9 c: ccccccccc
# '''
# for line in lines.splitlines():
#     print(line, is_valid2(line))

with open('day-2.txt') as fp:
    print(sum(is_valid2(line) for line in fp))
