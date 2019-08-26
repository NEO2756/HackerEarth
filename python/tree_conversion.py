class Node:
    data = 0
    left = None
    right = None


def preOrder(root):
    if root is None:
        return
    print(root.data, end=" ")
    preOrder(root.left)
    preOrder(root.right)


def constructBstUtil(root, arr, i):
    if arr[i] < root.data and root.left == None:
        node = Node()
        node.data = arr[i]
        node.left = node.right = None
        root.left = node
        return
    elif arr[i] >= root.data and root.right == None:
        node = Node()
        node.data = arr[i]
        node.left = node.right = None
        root.right = node
        return

    if arr[i] < root.data:
        constructBstUtil(root.left, arr, i)
    else:
        constructBstUtil(root.right, arr, i)


def constructBst(arr, n):
    root = Node()
    root.data = arr[0]
    root.left = root.right = None
    for i in range(1, n):
        constructBstUtil(root, arr, i)
    return root


def main():
    #testcases = int(input())
    testcases = 1
    while(testcases > 0):
        root = None
        #sizeOfArray = int(input())
        #arr = [int(x) for x in input().strip().split()]

        root = constructBst([7, 4, 12, 3, 6, 8, 1, 5, 10], 9)

        preOrder(root)
        print()

        testcases -= 1


if __name__ == "__main__":
    main()
