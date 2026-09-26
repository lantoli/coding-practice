def play(a, b):
    score = 0
    i = 0
    for x in a:
        if x > b[i]:
            score += 1
        else:
            i += 1
    return score

d = int(raw_input())

for i in range(1, d + 1):
    n = int(raw_input())
    a = map(float, raw_input().split())
    b = map(float, raw_input().split())
    a = list(reversed(sorted(a)))
    b = list(reversed(sorted(b)))
    print "Case #%d:" % i, n - play(b, a), play(a, b)
