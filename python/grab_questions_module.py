from sys import stdin, stdout
from collections import defaultdict
from collections import deque
# Reverse a sentence of words
arr = []


def reverse_sentence_of_words(string):
    arr = string.split()
    while len(arr) > 0:
        print(arr.pop(), end=' ')


def reverse_sentence_with_reversed_words(string):
    arr = string.split()
    while len(arr) > 0:
        print("".join(reversed(arr.pop())), end=' ')


def firstNonRepeating(arr):
    dic = defaultdict(lambda: int())
    for v in arr:
        dic[v] += 1
    for k, v in dic.items():
        if v == 1:
            print(k)
            break


def makeBinaryTree():
    dic = defaultdict(list)
    dic[1] = [2, 3]
    dic[2] = [4, 5]
    dic[3] = [6, 7]
    print(dic)


def levelOrderTraversal(tree):
    q = deque([tree[1]])  # queue root of tree
    mark = defaultdict()
    mark[1] = 0
    while len(q) > 0:
        p = q.popleft()
        left = 2 * p
        right = 2 * p + 1
        if left < len(tree) and tree[left] != 0:
            mark[left] = mark[p] - 1
            q.append(left)
        if left < len(tree) and tree[right] != 0:
            mark[right] = mark[p] + 1
            q.append(right)
    print(mark)


def topView(tree):
    q = deque([tree[1]])  # queue root of tree
    topview = defaultdict()
    mark = defaultdict()
    topview[1] = 0  # we can see root in top view always
    mark[1] = 0  # mark root as 0
    while len(q) > 0:
        p = q.popleft()
        left = 2 * p
        right = 2 * p + 1
        if left < len(tree) and tree[left] != 0:
            mark[left] = mark[p] - 1
            q.append(left)
            if not mark[left] in topview.values():
                topview[left] = mark[left]
        if left < len(tree) and tree[right] != 0:
            mark[right] = mark[p] + 1
            q.append(right)
            if not mark[right] in topview.values():
                topview[right] = mark[right]
    print(topview)
    view = [[k, v] for k, v in topview.items()]
    view.sort(key=lambda x: x[1])
    for i in range(0, len(view)):
        print(view[i][0], end=' ')


def bottomView(tree):
    q = deque([tree[1]])  # queue root of tree
    topview = defaultdict()
    mark = defaultdict()
    topview[1] = 0  # we can see root in top view always
    mark[1] = 0  # mark root as 0
    while len(q) > 0:
        p = q.popleft()
        left = 2 * p
        right = 2 * p + 1
        if left < len(tree) and tree[left] != 0:
            mark[left] = mark[p] - 1
            q.append(left)
            # if not mark[left] in topview.values():
            topview[mark[left]] = left
        if left < len(tree) and tree[right] != 0:
            mark[right] = mark[p] + 1
            q.append(right)
            # if not mark[right] in topview.values():
            topview[mark[right]] = right
    print(topview)
    view = [[k, v] for k, v in topview.items()]
    view.sort(key=lambda x: x[0])
    for i in range(0, len(view)):
        print(view[i][1], end=' ')


def columnTraverse(tree):
    q = deque([tree[1]])  # queue root of tree
    topview = defaultdict(list)
    mark = defaultdict()
    topview[0].append(1)  # we can see root in top view always
    mark[1] = 0  # mark root as 0
    while len(q) > 0:
        p = q.popleft()
        left = 2 * p
        right = 2 * p + 1
        if left < len(tree) and tree[left] != 0:
            mark[left] = mark[p] - 1
            q.append(left)
            # if not mark[left] in topview.values():
            topview[mark[left]].append(left)
        if left < len(tree) and tree[right] != 0:
            mark[right] = mark[p] + 1
            q.append(right)
            # if not mark[right] in topview.values():
            topview[mark[right]].append(right)
    print(topview)
    view = [[k, v] for k, v in topview.items()]
    print(view)
    view.sort(key=lambda x: x[0])
    for i in range(0, len(view)):
        print(view[i][1], end=' ')


if __name__ == "__main__":
    s = "sandeep sharma"
    reverse_sentence_of_words(s)
    print()
    reverse_sentence_with_reversed_words(s)
    print()
    arr = [-1, 2, -1, 3, 2]
    firstNonRepeating(arr)
    # makeBinaryTree()
    tree = [0, 1, 2, 3, 4, 5, 6, 7]
    levelOrderTraversal(tree)
    print("\ntop view")
    topView(tree)
    print("\nbottom view")
    bottomView(tree)
    print("\nColumn Traverse view")
    columnTraverse(tree)
