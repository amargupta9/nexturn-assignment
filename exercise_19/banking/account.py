accounts = {}

def create_account(account_id,name,initial_balance):
    if account_id in accounts:
        raise ValueError("Account ID already exists.")
    accounts[account_id] = {
        'name' : name,
        'balance' : initial_balance
    }
    return f"Account {account_id} created succesfully."


def get_account_info(account_id):
   account = accounts.get(account_id)
   if not account:
       raise ValueError("Account not found.")
   return account












# In this code, raise ValueError("Account ID already exists.") is used instead of a print statement because raise actually throws an error (specifically a ValueError), which interrupts the normal flow of the program and signals that something went wrong. This makes error handling more effective, for these reasons:

# Control Flow: Raising an error stops the function's execution immediately, preventing further actions. A print statement would just display a message but would allow the program to continue, possibly causing further issues if the account already exists.

# Error Handling: The raise statement allows other parts of the code to catch and handle the error using try and except blocks. This is useful for creating robust applications because you can specify exactly what to do when an error occurs, rather than just logging it.

# Intent and Documentation: Using raise ValueError clearly indicates that this situation (an existing account ID) is a serious issue, not just an informational message. This makes the code easier to understand and maintain.

# Logging and Debugging: Errors raised with raise are logged in error-handling systems with stack traces, which help diagnose where and why an error occurred.

# In short, raise makes your code safer and easier to manage compared to using print for errors.





