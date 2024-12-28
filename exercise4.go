     
package main

import "fmt"


func toFarenheit(celcius float64) float64{
	return (celcius * 9 / 5 ) + 32
}

func toCelsius (farenheit float64) float64{
	return (farenheit - 32) * 5 / 9
} 



func main (){
    var choice int 
	var temperature float64

	fmt.Println("Temperature conversion")
	fmt.Println("1. Convert Celsius to Fahrenheit")
    fmt.Println("2. Convert Fahrenheit to Celsius")
    fmt.Print("Enter your choice (1 or 2): ")
    fmt.Scan(&choice)

	if choice == 1 {
		fmt.Print("Enter temperature in Celsius: ")
		fmt.Scan(&temperature)
		fmt.Printf("%.2f Celsius is %.2f Farenheit\n",temperature ,toFarenheit(temperature))
	}else if choice == 2 {
        fmt.Print("Enter temperature in Fahrenheit: ")
        fmt.Scan(&temperature)
        fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", temperature, toCelsius(temperature))
    } else {
        fmt.Println("Invalid choice. Please enter 1 or 2.")
	}
}