# https://www.reddit.com/r/adventofcode/comments/18b4b0r/comment/kc2au1l/
# https://topaz.github.io/paste/#XQAAAQDyAgAAAAAAAAAzHIoib6poHLpewxtGE3pTrRdzrponK6uH8J9SK+0ES1LNdidG34FPT5Z5WfpAJ/epkMmnWNkE3uWVG0Y8HT99F+XNebEMfXdELf9EbJ4kU+4y8Z8mDTo7PYq6+UpBxHW3EhVkOrGX3a+DvB8oTz2hTXJdALhnxjOj4y3tG0GpOl9HlHOht5s9poDZYbVSMAOVoydIcuw6HPI5RHgaj0Dmnf4zU0plUYr3EoKI5ThFrVqKROa/OfWIF2UvFpNQjhWxoE9A1pEE/WsGzUpXdzMxbeUNCAyUDivAxRpAVwehvDQRA+tTYu0KlD1s7ZzjZbesgD+q055DMBD2vWEFXxTS8RaJeZmJXhqifO4l0xwb3NU/jMCEI2odBCcCZwzS6p9uh3Am1pILaAn1HGYLq6s22CWhDBx5BszlI+cy14Zxa5o28HYmiDcE9Q1fODzW5vhJoUt0t4sdraVoaMXTJk7bmgAAlz5+e/4JIbI=
from functools import reduce


seeds, *mappings = open('input.txt').read().split('\n\n')
seeds = list(map(int, seeds.split()[1:]))


def lookup(inputs, mapping):
    for start, length in inputs:
        while length > 0:
            for m in mapping.split('\n')[1:]:
                dst, src, len = map(int, m.split())
                delta = start - src
                if delta in range(len):
                    len = min(len - delta, length)
                    yield (dst + delta, len)
                    start += len
                    length -= len
                    break
            else:
                yield (start, length); break


print(*[min(reduce(lookup, mappings, s))[0] for s in [
    zip(seeds, [1] * len(seeds)),
    zip(seeds[0::2], seeds[1::2])]])
