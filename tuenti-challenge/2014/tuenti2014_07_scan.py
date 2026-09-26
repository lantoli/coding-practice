#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 7 - Yes we scan

We do equivalence classes with all the persons connected amongst them.
Two persons are connected (are in the same equivalence class) iif their root are the same.
"""

def main():
    source, target = int(input()), int(input())
    parents = {source: source, target: target}
    with open('phone_call.log', 'r') as f:
        for idx, line in enumerate(f):
            a, b = line.split()
            add_contacts(parents, int(a), int(b))
            if root(parents, source) == root(parents, target):
                print("Connected at {}".format(idx))
                return
    print("Not connected")


def add_contacts(parents, a, b):
    if a not in parents:
        parents[a] = a
    if b not in parents:
        parents[b] = b
    parents[root(parents,a)] = root(parents,b)


def root(parents, node):
    current = node
    changes = set()
    while current != parents[current]:
        changes.add(current)
        current = parents[current]
    for change in changes:
        parents[change] = current  # path compression
    return current


if __name__ == '__main__':
    main()
