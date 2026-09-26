import sys

FILENAME = "B-large-practice"

sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')


def ReadInts():
  return list(map(int, sys.stdin.readline().strip().split(" ")))

def Cross(a, b):
  for i in a:
    for j in b:
      yield (i, j)

def Neighbours(ui, uj, m, n):
  if ui - 1 >= 0: yield (ui - 1, uj)
  if uj - 1 >= 0: yield (ui, uj - 1)
  if uj + 1 < n: yield (ui, uj + 1)
  if ui + 1 < m: yield (ui + 1, uj)

N = ReadInts()[0]
for prob in range(1, N + 1):
  # Read the map
  (m, n) = ReadInts()
  maze = [ReadInts() for _ in range(m)]
  answer = [["" for _ in range(n)] for _ in range(m)]

  # The map from sinks to labels.
  label = {}
  next_label = 'a'

  # Brute force each cell.
  for (ui, uj) in Cross(range(m), range(n)):
    (i, j) = (-1, -1)
    (nexti, nextj) = (ui, uj)
    while (i, j) != (nexti, nextj):
      (i, j) = (nexti, nextj)
      for (vi, vj) in Neighbours(i, j, m, n):
        if maze[vi][vj] < maze[nexti][nextj]:
          (nexti, nextj) = (vi, vj)

    # Cell (ui, uj) drains to (i, j).
    if (i, j) not in label:
      label[(i, j)] = next_label
      next_label = chr(ord(next_label) + 1)
    answer[ui][uj] = label[(i, j)]

  # Output the labels.
  print ("Case #%d:" % prob)
  for i in range(m):
    print (" ".join(answer[i]))
