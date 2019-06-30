# Ahad Bawany
# Daily code 6/30:
# Given a list of numbers and a number key, print whether any two numbers from the list add up
# to the key
def oneSum(map, k):
   # cycle through the valid keys of the map
    for key, _ in map.items():
      	# key - value is equal to the number we need to add to a value to get the key
      	# each index is equal to its value, so if an index exists, that value exists
        if (k - map[key]) in map:
            print("true")
            break

# Create a map where the keys and values are identical
map = {
    3 : 3,
    7 : 7,
    10 : 10,
    15 : 15
}
oneSum(map, 17)
