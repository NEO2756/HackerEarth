from sys import stdin, stdout


def nextLargetElement(arr, target):

   # print(arr, target)
    start = 0
    end = len(arr) - 1
    ans = None
    while start <= end:
        mid = int((start + end)/2)
        if arr[mid] < target:
            start = mid + 1
        elif arr[mid] >= target:
            ans = mid
            end = mid - 1
    return ans


def merge2SortedArray(arr1, arr2):

    for j, v in reversed(list(enumerate(arr2))):
        last = arr1[len(arr1) - 1]
        pos = nextLargetElement(arr1, v)
        for i, _ in reversed(list(enumerate(arr1))):
            if i <= pos:
                arr1[pos] = v
                arr2[j] = last
                break
            arr1[i] = arr1[i - 1]


if __name__ == "__main__":
    arr1 = [1, 5, 9, 10, 15, 20]
    arr2 = [2, 3, 8, 13]
    merge2SortedArray(arr1, arr2)
    print(arr1, arr2)
