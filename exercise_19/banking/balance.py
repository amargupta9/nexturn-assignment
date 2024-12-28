from banking.account import accounts

def check_balance(account_id):
    if account_id not in accounts:
        raise ValueError ("Account not found.")
    return accounts[account_id]['balance']