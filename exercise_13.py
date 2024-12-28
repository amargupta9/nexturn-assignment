marks = list(map(float,input("enter marks of students separated by spaces: ").split())) 
#The program reads a list of marks from user input, separated by spaces, and converts each to a float.
# max() is used to find the highest mark, min() for the lowest mark, and sum()/len() for the average.
# If the list is empty, the average calculation avoids division by zero by setting it to 0.
# The highest, lowest, and average marks are then printed.
# This program efficiently calculates the statistics for the list of marks provided by the user.

# # Calculate highest, lowest, and average marks
highest_mark = max(marks)
lowest_mark = min(marks)
average_mark = sum(marks) / len(marks) if marks else 0


# if marks else 0:
# This part is a conditional expression that checks if marks is not empty.
# If marks has elements, it calculates the average as sum(marks) / len(marks).
# If marks is an empty list (meaning len(marks) is 0), it avoids division by zero and sets average_mark to 0.

# Output the results
print("Highest Mark:", highest_mark)
print("Lowest Mark:", lowest_mark)
print("Average Mark:", average_mark)