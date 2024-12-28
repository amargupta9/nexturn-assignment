salary = float (input("Enter your monthly salary: "))
room_rent = float(input("Enter your room rent: "))
grocery =  float(input("Enter your grocery expediture: "))
utility_costs =  float(input("Enter your utility_costs expediture: "))


surplus = salary - (room_rent + grocery + utility_costs)

rent_percentage = float((room_rent*100)/salary)
grocery_percentage = float((grocery*100)/salary)
utility_percentage = float((utility_costs*100)/salary)

# Display results
print("\nMonthly Expenses Breakdown:")
 
print(f" rent is ${room_rent:.2f} & expenditure  percentage of rent {rent_percentage:.2f}")
print(f"grocery is ${grocery:.2f} & expenditure  percentage of grocery {grocery_percentage:.2f}")
print(f"utility cost is ${utility_costs} &expenditure  percentage of utility {utility_percentage:.2f}")

# Display surplus or deficit
if surplus >= 0:
  print(f"\nYou have a surplus of ${surplus:.2f} this month.")
else:
  print(f"\nYou have a deficit of ${abs(surplus):.2f} this month.")
