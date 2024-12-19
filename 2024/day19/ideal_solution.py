# https://www.reddit.com/r/adventofcode/comments/1hhlb8g/comment/m2s56mc
from functools import cache

P, _,*T = open('day19/input.txt').read().split('\n')
P = P.split(', ')

@cache
def count(t):
    return sum(count(t.removeprefix(p)) for p in P
        if t.startswith(p)) or t == ''

# for t in T:
#     print(t, count(t))

for t in bool, int:
    print(sum(map(t, map(count, T))))