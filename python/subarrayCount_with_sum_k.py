
from collections import defaultdict


def findSubarraySum(arr, n, Sum):

    prevSum = defaultdict(lambda: 0)

    res = 0

    # Sum of elements so far.
    currsum = 0

    for i in range(0, n):

        # Add current element to sum so far.
        currsum += arr[i]

        if currsum == Sum:
            res += 1

        if (currsum - Sum) in prevSum:
            res += prevSum[currsum - Sum]

        prevSum[currsum] += 1

    return res


if __name__ == "__main__":

    arr = [10, 19, 19, 10]
    Sum = 29
    n = len(arr)
    print(findSubarraySum(arr, n, Sum))
