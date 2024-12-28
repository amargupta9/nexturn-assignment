def compound_interest_calculator(principle1,time1 , rate1):
    amount = principle1 *(1 + rate1/100)**time1
    compound_intreset = amount - principle1
    return compound_intreset

principle = float(input("enter principle amount: "))
rate = float(input("enter rate on principle amount: "))
time = float(input("enter principle time(in years): "))

compound_intreset = compound_interest_calculator(principle,rate,time)
print(f"compound intreset for {principle:.2f} is {compound_intreset:.2f}")