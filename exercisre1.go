
package main

import "fmt"


func main() {

	var date string
	var description string
	var amount float64

	date = "2024-12-22"
	description = "lunch at a cafe"
	amount = 210.50

	fmt.Printf(" Date: %s\n Description: %s\n Amount : $%.2f\n", date,description, amount)
}
