
attempts = 0

password = "amar@123#" # Replace with your desired password

while attempts < 3:
    user_password = input("enter your password: ")
    if password == user_password : 
           print("access granted")
           break
    else :
          attempts + 1
          print (f"Incorrect password. {3 - attempts} attempts remaining.")

if attempts == 3 :
      print("Too many attempts. Access denied.")