import sys


def min(a, b, c):
    t = a if a <= b else b
    t1 = t if t <= c else c
    return t1


def countMaxSizeSubMatrix_of_1s(matrix, N, M):
    result = [[0 for x in range(M)] for y in range(N)]
    result = matrix
    max = 0
    # result[i][j] = Represents the Max size of submatrix with i and j as its bootom right corner
    for i in range(1, N):
        for j in range(1, M):
            if matrix[i][j] == 1:
                result[i][j] = min(result[i - 1][j], result[i]
                                   [j - 1], result[i - 1][j - 1]) + 1
                if result[i][j] > max:
                    max = result[i][j]
    # print(result)
    return max


if __name__ == "__main__":
    T = int(input())
    for _ in range(T):
        N, M = input().split()
        N = int(N)
        M = int(M)
        # M = input()
        # N, M = tuple([int(x) for x in sys.stdin.readline().split()])
        matrix = [[int(x) for x in sys.stdin.readline(
            2 * M).strip().split()] for y in range(N)]
        print(countMaxSizeSubMatrix_of_1s(matrix, int(N), int(M)))
