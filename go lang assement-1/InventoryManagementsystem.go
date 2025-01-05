package main

import (
	"fmt"
	"strings"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product

// Add Product
func addProduct(id int, name string, price float64, stock int) {
	if stock < 0 {
		fmt.Println("Error: Stock cannot be negative.")
		return
	}

	for _, product := range inventory {
		if product.ID == id {
			fmt.Println("Error: Product ID must be unique.")
			return
		}
	}

	inventory = append(inventory, Product{ID: id, Name: name, Price: price, Stock: stock})
	fmt.Println("Product added successfully.")
}

// Update Stock
func updateStock(id int, stock int) {
	if stock < 0 {
		fmt.Println("Error: Stock cannot be negative.")
		return
	}

	for i := range inventory {
		if inventory[i].ID == id {
			inventory[i].Stock = stock
			fmt.Println("Stock updated successfully.")
			return
		}
	}
	fmt.Println("Error: Product not found.")
}

// Search Product
func searchProduct(query string) {
	for _, product := range inventory {
		if fmt.Sprint(product.ID) == query || strings.EqualFold(product.Name, query) {
			fmt.Printf("Product Found: ID: %d, Name: %s, Price: $%.2f, Stock: %d\n", product.ID, product.Name, product.Price, product.Stock)
			return
		}
	}
	fmt.Println("Error: Product not found.")
}

// Display Inventory
func displayInventory() {
	if len(inventory) == 0 {
		fmt.Println("No products in inventory.")
		return
	}

	fmt.Println("ID   Name               Price     Stock")
	fmt.Println("--------------------------------------")
	for _, product := range inventory {
		fmt.Printf("%-4d %-18s $%-8.2f %-5d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func main() {
	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("Exiting...")
			break
		}

		switch choice {
		case 1:
			var id, stock int
			var name string
			var price float64
			fmt.Print("Enter Product ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Product Name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter Product Price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter Product Stock: ")
			fmt.Scanln(&stock)
			addProduct(id, name, price, stock)
		case 2:
			var id, stock int
			fmt.Print("Enter Product ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter New Stock: ")
			fmt.Scanln(&stock)
			updateStock(id, stock)
		case 3:
			var query string
			fmt.Print("Enter Product ID or Name: ")
			fmt.Scanln(&query)
			searchProduct(query)
		case 4:
			displayInventory()
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
