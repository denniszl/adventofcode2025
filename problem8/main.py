import math
from decimal import *
from scipy.spatial import distance

# def eucliD(p, q):
#     print("p:", p, "q:", q)
#     x = p[0] - q[0]
#     y = p[1] - q[1]
#     z = p[2] - q[2]
#     print("x:", x, "y:", y, "z", z)
#     r = Decimal(x**x) + Decimal(y**y) + Decimal(z**z)
#     return math.sqrt(r)

def coord(c):
    split = c.split(',')
    return tuple(int(dim) for dim in split)

def find(sets, ds, x):
    if sets[x] != x:
        sets[x] = find(sets, ds, sets[x])
    return sets[x]

def getLargestMult(sets, n=3):
    disjointSets = {}
    for k in sets.keys():
        root = find(sets, k)
        if root not in disjointSets:
            disjointSets[root] = {}
        disjointSets[root][k] = True
    # print("disjoint sets:", disjointSets)
    lengths = [len(disjointSets[k]) for k in disjointSets.keys()]

    lengths.sort(reverse=True)
    # print("lengths:", lengths)

    total = lengths[0]
    for i in range(1, n):
        total *= lengths[i]
    return total

def connectCircuit(sets, ds, t):
    root0 = find(sets, ds, t[0])
    root1 = find(sets, ds, t[1])

    if root0 != root1:
        if root1 in ds:
            ds.pop(root1)
        sets[root1] = root0

def main():
    with open('input.txt') as f:
            coords = [coord(line.rstrip()) for line in f]
    sets = {}
    disjointSets = {}
    distances = []
    seenDistances = {}
    for obj in coords:
        disjointSets[obj] = {obj: True}
        sets[obj] = obj
        for c in coords:
            if c == obj:
                # print("skipping")
                continue
            key = tuple((obj, c))
            if tuple((c, obj)) in seenDistances or key in seenDistances:
                # print("skipping due to reverse been added already")
                continue
            seenDistances[key] = True
            distances.append(tuple((distance.euclidean(obj, c), key)))
    sortedDistances = sorted(distances, key=lambda distances: distances[0])
    # print([d[0] for d in sortedDistances ])
    # for i in range(0, 1000):
    #     connectCircuit(sets, sortedDistances[i][1])
    i = 0
    while len(disjointSets) > 1:
        lastTuple = sortedDistances[i][1]
        connectCircuit(sets, disjointSets, sortedDistances[i][1])
        i += 1
    print(lastTuple)
    print(lastTuple[0][0] * lastTuple[1][0])

if __name__=="__main__":
    main()
