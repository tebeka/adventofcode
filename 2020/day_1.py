prev = []
with open('day-1.txt') as fp:
    try:
        for line in fp:
            val = int(line.strip())
            for p in prev:
                if p + val == 2020:
                    print(p * val)
                    raise StopIteration
            prev.append(val)
    except StopIteration:
        pass
