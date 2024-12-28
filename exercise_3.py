while True:
    try:
       age = int(input("Enter the employee's age: ")) # Try to convert the input to an integer
        
       if age > 0 :
        break
       else:
        print("Please enter a valid age (greater than 0).")
    except ValueError:
      print("Invalid input. Please enter an integer.")
print(f"The employee's age is {age}.")


