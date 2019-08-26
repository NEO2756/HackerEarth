
def qsort(arr):
    if len(arr) < 2:
        return arr
    else:
        pivot = arr[0]
        less = [i for i in arr if i < arr[0]]
        high = [i for i in arr if i > arr[0]]
    return qsort(less) + [pivot] + qsort(high)


if __name__ == "__main__":
    print(qsort([10, 5, 2, 3]))
