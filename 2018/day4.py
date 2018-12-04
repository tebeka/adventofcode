# [1518-11-01 00:00] Guard #10 begins shift
# [1518-11-01 00:05] falls asleep
# [1518-11-01 00:25] wakes up

from collections import Counter, namedtuple, defaultdict
from datetime import datetime


Record = namedtuple('Record', 'time id state')


awake, asleep = 'awake', 'asleep'

example = '''\
[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up
'''


def parse_line(line):
    # [1518-11-01 00:00] Guard #10 begins shift
    # [1518-11-01 00:05] falls asleep
    # [1518-11-01 00:25] wakes up
    fields = line.split()
    # [1518-11-01 00:25]
    time = datetime.strptime(line[1:line.find(']')], '%Y-%m-%d %H:%M')
    # Entering shift is awake state
    state = asleep if fields[2] == 'falls' else awake
    # #99
    id = int(fields[3][1:]) if fields[2] == 'Guard' else None

    return Record(time, id, state)


def load_records(fp):
    lines = sorted(fp)  # Sort by time stamp
    return sorted(parse_line(line) for line in lines)


def part_1(records):
    """We assume records are sorted by date/time"""
    sleeps = defaultdict(Counter)  # id -> minutes asleep
    curr_id = None
    for i, record in enumerate(records):
        prev = records[i-1] if i > 0 else None
        if record.id:
            curr_id = record.id
        if record.state == awake and prev and prev.state == asleep:
            minutes = range(prev.time.minute, record.time.minute)
            sleeps[curr_id].update(minutes)

    def key(id):
        return sum(sleeps[key].values())

    id = sorted(sleeps, key=key)[0]
    (minute, _), = sleeps[id].most_common(1)
    return id * minute


# from io import StringIO
# records = load_records(StringIO(example))
with open('day-4.txt') as fp:
    records = load_records(fp)
print(part_1(records))
