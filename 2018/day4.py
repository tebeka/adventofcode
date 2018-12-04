import re
from collections import Counter, defaultdict


def parse_line(line):
    # [1518-11-01 00:00] Guard #10 begins shift
    # [1518-11-01 00:05] falls asleep
    # [1518-11-01 00:25] wakes up
    match = re.search(r'(\d+)\]', line)
    assert match, 'cannot find time in {!r}'.format(line)
    minute = int(match.group(1))

    match = re.search(r'Guard #(\d+)', line)
    guard = int(match.group(1)) if match else None
    state = 'asleep' if 'falls asleep' in line else 'up'

    return minute, guard, state


def load_records(fp):
    lines = sorted(fp)  # Sort by time stamp
    return [parse_line(line) for line in lines]


def part_1(records):
    """We assume records are sorted by date/time"""
    sleeps = defaultdict(Counter)  # id -> (minute -> count)
    sleep_minute = None
    curr_guard = None
    for (minute, guard, state) in records:
        if guard:
            curr_guard = guard
            sleep_minute = None
            continue

        if state == 'asleep':
            sleep_minute = minute
            continue

        if state == 'up' and sleep_minute:
            sleeps[curr_guard].update(range(sleep_minute, minute))
            sleep_minute = None

    def key(guard):
        return sum(sleeps[guard].values())

    guard = max(sleeps, key=key)
    (minute, _), = sleeps[guard].most_common(1)
    return guard * minute


# file_name = 'day-4-example.txt'
file_name = 'day-4.txt'
with open(file_name) as fp:
    records = load_records(fp)
print(part_1(records))
