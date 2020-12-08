import re

required = {
    'byr',
    'iyr',
    'eyr',
    'hgt',
    'hcl',
    'ecl',
    'pid',
    # 'cid',
}


data = '''
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
'''

with open('day-4.txt') as fp:
    data = fp.read()


count = 0
for sect in data.split('\n\n'):
    sect = sect.strip()
    if not sect:
        continue
    fields = set(re.findall(r'([a-z]{3}):\S', sect))
    if required - fields:
        continue
    count += 1
print(count)
