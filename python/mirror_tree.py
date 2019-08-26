# Initial Template for Python 3
class Node:

    def __init__(self, val):
        self.right = None
        self.data = val
        self.left = None

    def inorderTraversalUtil(root):
        # Code here
        if root is None:
            return
        inorderTraversalUtil(root.left)
        print(root.data, end=' ')
        inorderTraversalUtil(root.right)

    def inorderTraversal(root):
        # Code here
        inorderTraversalUtil(root)
        print()


if __name__ == '__main__':
    t = int(input())

    for i in range(t):
        n = int(input())
        arr = input().strip().split()
        if n == 0:
            print(0)
            continue

        root = None
        dictTree = dict()

        for j in range(n):
            if arr[3*j] not in dictTree:
                dictTree[arr[3*j]] = Node(int(arr[3*j]))
                parent = dictTree[arr[3*j]]

                if j is 0:
                    root = parent
            else:
                parent = dictTree[arr[3*j]]

            child = Node(int(arr[3*j+1]))

            if(arr[3*j+2] == 'L'):
                parent.left = child
            else:
                parent.right = child
            dictTree[arr[3*j+1]] = child

        mirror(root)

        inorderTraversal(root)

''' This is a function problem.You only need to complete the function given below '''
# User function Template for python3
'''
class Node:
    def __init__(self, val):
        self.right = None
        self.data = val
        self.left = None
'''
# your task is to complete this function


def mirror(root):
    if root == Node:
        return
    if root.left != None or root.right != None:
        root.left, root.right = root.right, root.left

    if root.left != None:
        mirror(root.left)
    if root.right != None:
        mirror(root.right)
