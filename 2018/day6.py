# FIXME: Wrong answer
from collections import defaultdict, namedtuple
from functools import partial

Coord = namedtuple('Point', 'x y')


def dist(c1, c2):
    return abs(c1.x - c2.x) + abs(c1.y - c2.y)


def calc_areas(coords):
    min_x = min(c.x for c in coords)
    max_x = max(c.x for c in coords)
    min_y = min(c.y for c in coords)
    max_y = max(c.y for c in coords)

    areas = defaultdict(list)
    for x in range(min_x, max_x + 1):
        for y in range(min_y, max_y + 1):
            c = Coord(x, y)
            key = partial(dist, c)
            c1 = min(coords, key=key)
            c2 = min((co for co in coords if co != c1), key=key)
            if key(c1) == key(c2):
                continue
            areas[c1].append(c)
    return areas


def parse_line(line):
    x, y = (int(v.strip()) for v in line.split(','))
    return Coord(x, y)


def is_inf_area(c, coords):
    if min(co.x for co in coords) >= c.x:
        return True
    if max(co.x for co in coords) <= c.x:
        return True
    if min(co.y for co in coords) >= c.y:
        return True
    if max(co.y for co in coords) <= c.y:
        return True
    return False


def part_1(coords):
    areas = calc_areas(coords)
    finite = {c: cs for c, cs in areas.items() if not is_inf_area(c, cs)}
    print(max(len(cs) for cs in finite.values()))


data = '''\
1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
'''

# coords = [parse_line(line) for line in data.splitlines() if line.strip()]
with open('day-6.txt') as fp:
    coords = [parse_line(line) for line in fp]
part_1(coords)
