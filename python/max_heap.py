from sys import stdin, stdout

tree = [0, 10, 40, 15, 50, 100, 30, 10]
heap = []


def insert(new):
    heap.append(new)
    idx = len(heap) - 1
    while heap[idx//2] != -1:
        if heap[idx//2] < heap[idx]:
            heap[idx//2], heap[idx] = heap[idx], heap[idx//2]
            idx = idx//2
            continue
        else:
            break


def heapify(idx):
    left = 2 * idx
    right = 2 * idx + 1

    heapSize = len(heap) - 1
    if left <= heapSize and heap[left] > heap[idx]:
        heap[left], heap[idx] = heap[idx],  heap[left]
        return heapify(left)

    if right <= heapSize and heap[right] > heap[idx]:
        heap[right], heap[idx] = heap[idx],  heap[right]
        return heapify(right)


def delete(targetnodeIdx):
    if targetnodeIdx == len(heap) - 1:
        heap.pop()
        return
    heap[targetnodeIdx] = heap.pop()
    heapify(targetnodeIdx)


def ExtractMax():
    e = heap[1]
    heap[1] = heap[len(heap) - 1]
    heapify(1)
    return e


if __name__ == "__main__":
    heap.append(-1)  # parent of root is -1
    for v in range(1, len(tree)):
        insert(tree[v])
    print('Heap Formed')
    print(heap)
    print('Delete node at index 3 in heap')
    delete(3)
    print('After deletion')
    print(heap)
    print('Extract Max')
    print(ExtractMax())
    print('This is the result')
    print(heap)
    print('Extract Max Again')
    print(ExtractMax())
    print('This is the result again')
    print(heap)
