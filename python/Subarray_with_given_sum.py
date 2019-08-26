from sys import stdin, stdout


def Subarray_with_given_sum():
    T = int(input())
    for _ in range(T):
        N, S = [int(x) for x in stdin.readline().split()]
        arr = [int(x) for x in stdin.readline().split()]
        start = 0
        sum = arr[start]
        i = 1
        while i <= N:
            while sum > S and start <= i - 1:
                sum = sum - arr[start]
                start += 1
            if sum == S:
                print(start + 1, i)
                break
            if i < N:
                sum += arr[i]
            i += 1
        else:
            print("-1")


Subarray_with_given_sum()
