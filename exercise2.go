package main

import "fmt"

func main() {
	var side1 , side2 , side3 int 
	fmt.Println("enter length three sides of triangle: ")
	fmt.Scan(&side1,&side2,&side3)

	if side1 == side2 && side2 == side3 {
     fmt.Println("The triangle is Equilateral (all sides are equal).")
  }   else if side1 == side2 || side1 == side3 || side2 == side3 {
	fmt.Println("The triangle is Isosceles (two sides are equal).")
  } else{
	fmt.Println("The triangle is Scalene (no sides are equal).")
  }
}