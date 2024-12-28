from banking.account import accounts

def deposit(account_id, amount):
    if account_id not in accounts:
        raise ValueError("Account not found.")
    if amount <= 0 :
        raise ValueError("Deposit amount must be positive.")
    accounts[account_id]['balance'] += amount
    return f"{amount} deposited to account {account_id}."

def withdraw(account_id, amount):
     if account_id not in accounts:
        raise ValueError("Account not found.")
     if amount <= 0 :
        raise ValueError("withdrawal amount must be positive.")
     if accounts[account_id]['balance'] < amount:
         raise ValueError("insufficient balance.")
     accounts[account_id]['balance'] -= amount
     return f"{amount} withdrawn from account{account_id}."

      