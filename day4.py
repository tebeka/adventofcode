def is_valid(passphrase):
    words = passphrase.split()
    return len(words) == len(set(words))


with open('input.txt') as fp:
    print(sum(1 for line in fp if is_valid(line)))
