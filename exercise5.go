
package main 

import "fmt"
func main () {
  var name string 
  var age int 
  var favoriteNumber float64


  fmt.Print("enter your name: ")
  fmt.Scanln(&name)
  
  fmt.Print("enter your age: ")
  fmt.Scanln(&age)

  fmt.Print("enter your favorite number : ")
  fmt.Scanln(&favoriteNumber)


  fmt.Printf("Hello, %s! You are %d years old, and your fovourite number is %.2f.\n",name,age,favoriteNumber)

}