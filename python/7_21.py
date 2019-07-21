# Ahad Bawany
# Daily code 6/30:
# Given 2 lists of numbers, find the intersection point within them

# Lists to test our function 
a = [1, 2, 3, 5, 6, 8, 9]
b = [0, 7, 5, 6, 7, 8]

# Given a list, a, and a list, b, return the intersection
def findIntersect(a,b):
    
    # Ensure both a and b exist
    if (a == None or b == None):
        return False
    
    # Only cycle through the length of a if it's the smaller list to prevent index errors
    if (len(b) >= len(a)):
        for i in range(len(a)):
            if a[i] == b[i]:
                return a[i]
    
    # Vice versa, make sure to cycle through the smaller list!
    else:
        for i in range(len(b)):
            if b[i] == a[i]:
                return a[i]

    # If no repeating elements are found, just return false.
    return False

print(findIntersect(a,b))