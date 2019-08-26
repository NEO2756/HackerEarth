import sys


def calculate(a, b, opr):
    if opr == '+':
        return a+b
    elif opr == '-':
        return a - b
    elif opr == '*':
        return a*b
    else:
        return a//b


def evalExp(arr):
    for i in range(0, len(arr)-1):
        i += 1
        if arr[i].isdigit():
            break
    i -= 1
    while i >= 0:
        arr[i] = calculate(int(arr[2 * i + 1]), int(arr[2 * i + 2]), arr[i])
        i -= 1
    return arr[0]


if __name__ == "__main__":
    T = int(input())
    for i in range(0, T):
        n = int(input())
        arr = [x for x in sys.stdin.readline().strip().split()]
        print(evalExp(arr))
