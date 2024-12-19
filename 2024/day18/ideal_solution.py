# https://www.reddit.com/r/adventofcode/comments/1hguacy/comment/m2missp
from collections import deque

bytes = [eval(line) for line in open("day18/input.txt").read().splitlines()]


def bfs(grid, start, end):
    queue = deque([(start, 0)])
    seen = set()
    while queue:
        loc, dist = queue.popleft()
        if loc == end: return dist
        queue.extend((loc + d, dist + 1) for d in (-1, 1, -1j, 1j) if
                     loc + d in grid and loc not in seen)
        seen.add(loc)


def pathfinder(bytes, H, W, C):
    start, end = complex(0, 0), complex(W - 1, H - 1)
    grid = {complex(x, y) for y in range(H) for x in range(W)} - set(
        complex(x, y) for x, y in bytes[:C])
    return bfs(grid, start, end)


print('part 1: ', pathfinder(bytes, 71, 71, 1024))

for i in range(1024, len(bytes)):
    path = pathfinder(bytes, 71, 71, i + 1)
    if not path:
        print('part 2:', bytes[i])
        break
