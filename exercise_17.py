student = [("amar",95),("rohit",56),("suman",67),("saharash",87)]

sorted_by_marks = sorted(student ,key=lambda x:student[1])

# Print the sorted list
print("Students sorted by marks:")
for name, mark in sorted_by_marks:
    print(f"{name}: {mark}")


    """ The key difference between sorted() and filter() when used with a lambda function lies in how each function utilizes the lambda expression and in the format of the function's arguments.

Here’s a breakdown of each:

1. sorted() with a Lambda Function
Purpose: sorted() sorts a list based on the specified sorting key.
Syntax with Lambda: sorted(iterable, key=lambda element: condition)
The key parameter requires a function (or lambda) that specifies what part of each element should be used for sorting.
The lambda function should return a value for comparison but does not need to return a Boolean.
Example:

python
Copy code
sorted_students = sorted(students, key=lambda student: student[1])
Here, student[1] represents the mark of each student, and sorted() uses this mark to sort the list.

2. filter() with a Lambda Function
Purpose: filter() selects elements based on a condition.
Syntax with Lambda: filter(lambda element: condition, iterable)
The lambda function in filter() needs to return a Boolean value (either True or False).
Only elements that return True from the lambda function will be included in the filtered output.
Example:

python
Copy code
passing_students = list(filter(lambda student: student[1] >= 50, students))
Here, the lambda student[1] >= 50 checks if the student’s mark is 50 or above. If the condition is True, that student is included in passing_students.

Summary of Key Differences
Return Type of Lambda:
sorted()’s lambda should return a sorting key (e.g., student[1]).
filter()’s lambda should return a Boolean indicating whether to keep each element.
Parameter:
sorted() uses key= to apply the lambda to each element.
filter() directly takes the lambda as its first argument.
If you swap these behaviors (e.g., use a Boolean return for sorted() or a value return for filter()), it will result in an error or unexpected behavior.
"""





