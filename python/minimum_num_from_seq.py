from sys import stdin, stdout
arr = []


def minimum_num_from_I_D_sequence(seq):
    for i in range(1, len(seq) + 2):
        arr.append(i)
        if i <= len(seq) and seq[i-1] == 'D':
            continue
        elif (i <= len(seq) and seq[i-1] == 'I') or (i == (len(seq) + 1)):
            while len(arr) > 0:
                print(arr.pop(), end='')


if __name__ == "__main__":
    minimum_num_from_I_D_sequence("DDIDDIID")
