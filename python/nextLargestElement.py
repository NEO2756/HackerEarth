from sys import stdin, stdout


class items:
    def __init__(self, v, i):
        self.val = v
        self.idx = i


def nextLargetElement(arr, target):

    start = 0
    end = len(arr) - 1
    ans = None
    while start <= end:
        mid = int((start + end)/2)
        if arr[mid].val > target:
            start = mid + 1
        elif arr[mid].val <= target:
            ans = arr[mid]
            end = mid - 1
    return ans


if __name__ == "__main__":
    T = int(input())
    while T > 0:
        stack = []
        arr = []
        ans = i = 0
        T -= 1
        N = int(input())
        for x in stdin.readline().split():
            arr.append(items(int(x), i))
            i += 1
        i = 0
        #arr = [int(x) for x in stdin.readline().split()]
        for v in arr:
            if len(stack) == 0:
                stack.append(v)
                i += 1
                continue
            if v.val < stack[len(stack) - 1].val:  # newitem < top
                stack.append(v)
            else:
                e = nextLargetElement(stack, v.val)
                temp = i - e.idx
                if temp > ans:
                    #print(stack, v, arr[e])
                    ans = temp
            i += 1
        print(ans)
