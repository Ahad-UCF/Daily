# Ahad Bawany
# Daily code 7_3_19:
# Given a list of numbers, return the smallest positive integer that isn't part of the list

def smallestNotPresent(a):
    # Set b to 1 since that's the smallest positive value possible
    b = 1
    for i in a:
        # If the value within a is equal to our current b, that can't be the smallest possible value not present
        if i == b and i > 0:
            b = b + 1
    return b

a = [3, 4, -1, 1]
print(smallestNotPresent(a))
