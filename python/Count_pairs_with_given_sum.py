import sys
# import os
from collections import defaultdict


def countPairs(arr, targetSum):
    dic = defaultdict(list)
    count = 0
    for v in arr:
        if (targetSum - v) in dic:
            count += len(dic[targetSum - v])
        dic[v].append(1)
    return count


if __name__ == "__main__":

    #T = int(input())

    for i in range(0, 1):
       # _, k = tuple([int(x) for x in sys.stdin.readline().split()])
        #arr = [int(x) for x in sys.stdin.readline().split()]
        arr = [1, 5, 7, 1]
        print(countPairs(arr, 6))
