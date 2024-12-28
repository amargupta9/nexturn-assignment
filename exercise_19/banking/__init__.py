from .account import create_account,get_account_info
from .transaction import deposit , withdraw
from .balance import check_balance

__all__ = ['create_account','get_account_info','deposit','withdraw','check_balance']



# Practical Example with and without __init__.py
# Without __init__.py
# If banking had no __init__.py:

# python
# Copy code
# # main.py
# from banking.account import create_account  # You need to specify the exact module path.
# from banking.transaction import deposit
# With __init__.py
# Since __init__.py exposes these functions:

# python
# Copy code
# # main.py
# from banking import create_account, deposit  # You can import functions directly from `ban


#  What is __init__.py?
# In Python, the presence of an __init__.py file in a directory marks it as a package. Without this file, Python treats the folder as a regular directory and won’t recognize it as a package, which would prevent importing functions or modules from it.


# Why We Create __init__.py in the banking Package
# In your banking package, __init__.py serves several key functions:

# Defining Available Imports: By listing specific functions in __init__.py, you control what users can import directly from the package. Without it, they would need to import functions from each module separately, like from banking.account import create_account.
# Organizing Access to Functions: This file makes it easier for users to access functions directly from the package, so they can write from banking import create_account instead of specifying each module path.