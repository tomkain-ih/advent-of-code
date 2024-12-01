# # https://www.reddit.com/r/adventofcode/comments/18b4b0r/comment/kc2au1l/
# https://topaz.github.io/paste/#XQAAAQCYAQAAAAAAAAAzHIoib6poHLpewxtGE3pTrRdzrponK6uH8J9SK+0ES1LNdidG34FPT5Z5WfpAJ/epkMmnWNkE3uWVG0Y8HT99F+XNebEMfXdELf9EbJ4kU+4y8Z8mDYq4vFjdXFjqAZ+RcsKBrxM8B4MNbXk+8ON8Eg2UVhAETGWE/hiQTSOEXcL7kKbvTlDhwyye9cGIwJft/OhiPJfgAxk9De6uGFbA+h8JLBXvO7eA+06Kevk3KqYuRoSivXiNXjclP6RVPreuYtpQVlnhamNsljGPuCrFbhVpyfshGpn/tnVhDosGz+KbtVDUqXyEuIxPZd+TxHcBZD6B/P91HOcA
from functools import reduce

seeds, *mappings = open('input.txt').read().split('\n\n')
seeds = map(int, seeds.split()[1:])

def lookup(start, mapping):
    for m in mapping.split('\n')[1:]:
        dst, src, len = map(int, m.split())
        delta = start - src
        if delta in range(len):
            return dst + delta
    else: return start

print(min(reduce(lookup, mappings, int(s)) for s in seeds))