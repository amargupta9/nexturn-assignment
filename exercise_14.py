names = input("Enter the students' names separated by commas: ").split(',')

names = [name.strip() for name in names]

names.sort()

longest_name = max(names , key = len ) if names else None


# max() returns the largest element in the names list.
# The key=len argument tells max() to determine the "largest" based on the length of each name (rather than alphabetical order).
# So, max(names, key=len) will return the name with the most characters in names.

print("sorted names : ", names)

print("longest_names : ",longest_name)