# FIXME: Wrong answer
import re
from collections import namedtuple

Rect = namedtuple('Rect', 'top_x top_y bottom_x bottom_y')


def parse_line(line):
    # #1 @ 861,330: 20x10
    match = re.search(r'(\d+) @ (\d+),(\d+): (\d+)x(\d+)', line)
    assert match, f'bad line: {line}'
    x, y, w, h = map(int, match.groups()[1:])
    return Rect(x, y, x+w, y+h)


def overlap(r1, r2):
    # Max(0, Min(XA2, XB2) - Max(XA1, XB1))
    # *
    # Max(0, Min(YA2, YB2) - Max(YA1, YB1))
    width = max(0, min(r1.bottom_x, r2.bottom_x)-max(r1.top_x, r2.top_x))
    height = max(0, min(r1.bottom_y, r2.bottom_y)-max(r1.top_y, r2.top_y))
    return width * height


data = '''
#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2
'''
# rects = [parse_line(line) for line in data.splitlines() if line.strip()]


with open('day-3.txt') as fp:
    rects = [parse_line(line) for line in fp]

n = 0
for i, r1 in enumerate(rects):
    for r2 in rects[i+1:]:
        n += overlap(r1, r2)
print(n)
