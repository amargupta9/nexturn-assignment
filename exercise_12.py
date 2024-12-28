# Program to print all prime numbers between 1 and 50
for num in range(2, 51):  # Starting from 2, as 1 is not a prime number
    is_prime = True
    for i in range(2, int(num ** 0.5) + 1):
        if num % i == 0:
            is_prime = False
            break
    if not is_prime:
        continue  # Skip to the next number if num is not prime
    print(num)
