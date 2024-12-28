# main.py
from banking import create_account, get_account_info, deposit, withdraw, check_balance

# Create a new account
print(create_account('123', 'Alice', 500))

# Check account information
print(get_account_info('123'))

# Deposit money
print(deposit('123', 200))

# Check balance after deposit
print(check_balance('123'))

# Withdraw money
print(withdraw('123', 100))

# Check balance after withdrawal
print(check_balance('123'))
