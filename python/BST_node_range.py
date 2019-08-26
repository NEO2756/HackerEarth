class Node:
    def __init__(self):
        self.data = 0
        self.left = None
        self.right = None


def createNewNode(value):
    temp = Node()
    temp.data = value
    temp.left = None
    temp.right = None
    return temp


def newNode(root, data):
    if(root is None):
        root = createNewNode(data)
    elif(data < root.data):
        root.left = newNode(root.left, data)
    else:
        root.right = newNode(root.right, data)

    return root


def main():
    testcases = int(input())
    while(testcases > 0):
        root = None
        global count
        count = 0
        sizeOfArray = int(input())
        arr = [int(x) for x in input().strip().split()]
        for i in arr:
            root = newNode(root, i)

        lr = [int(x) for x in input().strip().split()]
        print(getCountOfNode(root, lr[0], lr[1]))
        testcases -= 1


count = 0


def getCountOfNode(root, l, h):
    if root == None:
        return
    if l <= root.data and h >= root.data:
        global count
        count = count + 1

    if root.data > h:
        getCountOfNode(root.left, l, h)
    elif root.data < l:
        getCountOfNode(root.right, l, h)
    else:
        getCountOfNode(root.left, l, h)
        getCountOfNode(root.right, l, h)
    return count


if __name__ == "__main__":

    main()
