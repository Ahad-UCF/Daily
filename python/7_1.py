# Ahad Bawany
# Daily code 7/1/2019:
# Given a list of numbers, return a list of numbers where each number is the result of multiplying
# all other numbers in the list

def multiplArr(a):
    b = []
    # Go through all the elements of b and make them 1
    for i in a:
        b.append(1)

    # len b is the length of b, using range goes from 0 -> len b
    # if we didn't use range, we'd have for i in b which would cycle through the VALUES of b not the indices
    for i in range(len(b)):
        for j in range(len(a)):
            # if the indecies of b and a are the same, we know that we're at the index we want to skip
            if (i != j):
                b[i] *= a[j]
    return b

a = [3, 2, 1]
b = multiplArr(a)
print(b)
