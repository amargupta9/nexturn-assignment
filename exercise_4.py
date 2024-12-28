name = input("Enter your name: ")

while True:
    try:
            salary = float(input("Enter your salary: "))
            break
    except  ValueError:
            print("Invalid input! Ensure salary is a number (float).")
while True:
    try:
         no_of_independent = int(input("enter the no of dependent: ")) 
         break
    except ValueError:
              print("Invalid input! Ensure number of dependent is an integer.")

print(f"Your name is {name},your salary is {salary} and you have {no_of_independent} dependets.")