from math import floor


def calcFuel(mass):
    return floor(mass/3) - 2


def fuel_fuel(fuel):
    total = 0
    fuel = calcFuel(fuel)
    while fuel > 0:
        total += fuel
        fuel = calcFuel(fuel)
    return total


total = 0
with open('day-1.txt') as fp:
    for line in fp:
        f = calcFuel(float(line))
        ff = fuel_fuel(f)
        total = total + f + ff
print(total)
