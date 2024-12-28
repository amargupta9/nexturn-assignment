products = {}

while True:
    product_name = input("enter a product name: ")
    if product_name.lower() == "done":
        break 
    quantity = int(input(f"enter the quantity of {product_name}: "))
    products[product_name] = quantity

    print("\nInventory:")
    for product , quantity in products.items():
        print(f"{product} : {quantity}")






