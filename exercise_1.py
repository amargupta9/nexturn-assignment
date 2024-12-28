

def create_invoice():

      # Get user's name
    name = input("Enter your Name: ")

   # Get the number of products
    num_products = int(input("Enter the number of product you baught:  "))

# Initialize an empty list to store product names and prices
    products = []

 # Loop through the number of products and get the name and price for each
    for i in range(num_products):
      products_name = input(f"Enter the product name {i + 1}: ")
      products_price = int(input(f"Enter price of {products_name}: $"))
      products.append ((products_name,products_price))

 # Calculate the total cost
   #  total_cost += products_price

    total_cost = sum(price for _, price in products)

    #print the invoice
    print("\n -------------Invoice ------------")
    print(f"Customer Name: {name}")
    print("\nitems purchased:")
    for i, (products_name,products_price) in enumerate(products,1):
       print(f"{i}.{products_name} - ${products_price:2f}") 


    print("\nTotal Cost: ${:.2f}".format(total_cost))
    print("-------------------")


  # Call the function to run the program
create_invoice()  


# Get user input
user_name = input("Enter your name: ")
num_products = int(input("Enter the number of products you bought: "))
products = []
total_cost = 0

# # Collect product details
# for i in range(num_products):
#     product_name = input(f"Enter the name of product ({i + 1}): ")
#     product_price = float(input(f"Enter the price of ({product_name}): $"))
#     products.append((product_name, product_price))
#     total_cost += product_price

# # Print the invoice
# print("\n" + "=" * 30)
# print(f"Invoice for {user_name}")
# print("=" * 30)
# print(f"{'Product': <20}{'Price ($)':>10}")
# print("-" * 30)

# for product in products:
#     print(f"{product[0]: <20}{product[1]:>10.2f}")

# print("-" * 30)
# print(f"{'Total': <20}{total_cost:>10.2f}")
# print("=" * 30)
