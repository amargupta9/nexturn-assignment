temperature = float(input("Enter the temperature: "))
unit = input("Enter the unit of temperature (C for celsius, F for Fahrenheit): ").strip().upper()#input("Enter the Unit (C for celsius, F for Fahrenheit): "): This part displays the prompt asking the user to enter the temperature unit.
# .strip(): This removes any leading or trailing whitespace from the user's input. This ensures the input is consistent, even if the user accidentally adds spaces.
# .upper(): This converts the user's input to uppercase. This makes the code less sensitive to case variations, as the user could enter "c" or "f" instead of "C" or "F".

if(unit == "C"):
   temp_in_F = (temperature * 9/5 ) + 32
   print(f"{temperature:.2f}°C is equal to {temp_F:.2f}°F." )
elif(unit == "F"):
   temp_in_C = (temperature - 32) * 5/9
   print(f"{temperature:.2f}°F is equal to {temp_in_C}°C.")
else:
   print("enter valid temperature or unit")