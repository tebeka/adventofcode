from itertools import permutations


def is_valid(passphrase):
    words = passphrase.split()
    return len(words) == len(set(words))


with open('input-4.txt') as fp:
    print(sum(1 for line in fp if is_valid(line)))


def is_valid2(passphrase):
    words = passphrase.split()
    for i, word in enumerate(words):
        other = words[:i] + words[i+1:]
        for perm in permutations(word):
            if ''.join(perm) in other:
                return False
    return True


with open('input-4.txt') as fp:
    print(sum(1 for line in fp if is_valid2(line)))
