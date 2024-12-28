
balance = 0

def deposits_balance(amount):
    global balance
    if amount > 0:
      balance = balance + amount
      print(f"{amount} deposited sucessfully in your account & your new balance is {balance}")
    else:
        print("Deposit amount must be positive.")


def withdraw_amount(amount):
      global balance
      if amount > balance:
        print("Insufficient funds for this withdrawal.")
      elif amount > 0:
        balance -= amount 
        print(f"{amount} withdraw  from  your account & your new balance is {balance}")
      else:
         print("Withdrawal amount must be positive.")

def current_balance():
    print(f"your current balance is {balance}")


deposits_balance(5000)
current_balance()
withdraw_amount(2000)
withdraw_amount(1000)
current_balance()
