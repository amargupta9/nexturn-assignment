basic_salary = float(input("enter your basic salary: "))
tax_rate = float(input("enter tax rate(as a decimal): "))
bonus = float(input("enter your bonus amount: "))
tax_deduction = basic_salary*tax_rate
net_salary = basic_salary  - tax_deduction + bonus
 
 # Display results
print("\nSalary Details:")
print(f"Basic Salary: {basic_salary}")
print(f"Tax Rate: {tax_rate*100}%")
print(f"your net salary is {net_salary:.2f}") 
print(f"your total deduction is {tax_deduction:.2f}")