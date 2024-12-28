package main

import "fmt"

func main()  {

	var number int 
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&number)

	factorial := 1

	for i := 1;i<=number ;i++{
		factorial *= i
	}

	fmt.Printf("The factorial of %d is %d \n", number ,factorial)
}