# Given a string, count number of subsequences of the form aibjck, i.e., it consists of i ’a’ characters, followed by j ’b’ characters, followed by k ’c’ characters where i >= 1, j >=1 and k >= 1.

# Note: Two subsequences are considered different if the set of array indexes picked for the 2 subsequences are different.
# Expected Time Complexity : O(n)
# Examples:
#Input  : abbc
#Output : 3
# Subsequences are abc, abc and abbc
#Input  : abcabc
#Output : 7
# Subsequences are abc, abc, abbc, aabc
#abcc, abc and abc


def countSequence(sequence):
    count_a = count_b = count_c = 0
    for char in sequence:
        if char == 'a':
            count_a = 1 + 2 * count_a
        elif char == 'b':
            count_b = count_a + 2 * count_b
        else:
            count_c = count_b + 2 * count_c
    return count_c


if __name__ == "__main__":
    T = int(input())
    for i in range(0, T):
        sequence = input()
        print(countSequence(sequence))
