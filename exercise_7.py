# your_age = int(input("enter your age: "))

# if your_age < 5 : 
#     print("No price because your age is below 5: ")
# elif your_age <= 12 :
#     print("movie  ticket price  for you 5$ ")
# elif your_age <= 60 :
#     print(" movie  ticket price  for you 12$ ")
# else: 
#     print(" movie  ticket price  for you 8$ ")


age = int(input("Enter your age: "))

if age < 5:
  price = 0
elif age <= 12:
  price = 5
elif age <= 60:
  price = 12
else:
  price = 8

print(f"Your ticket price is: ${price}")