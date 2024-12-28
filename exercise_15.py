def calculate_grade(marks):
    if marks >= 90:
        print(f"your obtained grade is : A")
    elif marks >=75:
        print(f"your obtained grade is : B")
    elif marks >= 50:
        print(f"your obtained grade is : C")
    else:
        print(f"your obtained grade is : F")

marks = int(input("enter your maks: "))
calculate_grade(marks)
