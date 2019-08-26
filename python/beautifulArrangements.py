def isKthBitSet(n, k):
    if n & (1 << (k - 1)):
        print(1)
    else:
        print(0)
