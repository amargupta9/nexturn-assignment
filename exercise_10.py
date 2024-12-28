first_range = int(input("enter a starting  number: "))
second_range = int(input("enter a end number: "))

for num in range (first_range , second_range):
    if num % 2 == 0:
       print(f"{num}")
      