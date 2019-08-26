'''
        1
       |  |
      2    3
    |  |  |  |
    4  5  6  7
    
    Internal node: Nodes not visible from right and left, i.e 5 & 6
       '''
tree = []


def printInternalNodes(tree):
    while(tree != None):
        arr.remove(tree.data)
        tree = tree.left
    while(tree.right != None):
        arr.remove(tree.data)
        tree = tree.left
    print(arr)
