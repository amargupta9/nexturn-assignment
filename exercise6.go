package main 

import "fmt"


func main() {
   var targetnumber int
   var noofattempts int 
   var userinput int 

  fmt.Println("Enter target number ")
  fmt.Scanln(&targetnumber)

  fmt.Println("Enter no of attempts")
  fmt.Scanln(&noofattempts)



  for i:=noofattempts ; i>0 ; i--{

	fmt.Println("enter userinput")
    fmt.Scanln(&userinput)
     
	if userinput == targetnumber{
		fmt.Println("correct number , you win game.")
		break;
	}else if (userinput > targetnumber){
		fmt.Println("too high")
	}else if (userinput < targetnumber){
		fmt.Println("too low")
	}

  }   
    fmt.Println("you reach maximum number of attempts . you lost the game. ")
}
//********************************************************************************************************************************************
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // Function to return a hint based on the user's guess
// func getHint(target, guess int) string {
// 	if guess < target {
// 		return "Too low"
// 	} else if guess > target {
// 		return "Too high"
// 	} else {
// 		return "Correct"
// 	}
// }

// func main() {
// 	// Step 1: Declare variables for the target number, number of attempts, and user input
// 	var target, guess, attempts int
// 	maxAttempts := 5

// 	// Seed the random number generator to get different numbers on each run
// 	rand.Seed(time.Now().UnixNano())

// 	// Generate a random target number between 1 and 100
// 	target = rand.Intn(100) + 1

// 	// Step 2: Allow the user to keep guessing
// 	fmt.Println("Welcome to the Number Guessing Game!")
// 	fmt.Printf("Guess the number between 1 and 100. You have %d attempts.\n", maxAttempts)

// 	for attempts = 1; attempts <= maxAttempts; attempts++ {
// 		// Step 5: Prompt the user for their guess
// 		fmt.Print("Enter your guess: ")
// 		fmt.Scan(&guess)

// 		// Step 3: Check the guess using the getHint function
// 		hint := getHint(target, guess)
// 		fmt.Println(hint)

// 		// Step 4: Break the loop if the user guesses correctly
// 		if hint == "Correct" {
// 			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
// 			return
// 		}
// 	}

// 	// If the user didn't guess correctly in the given attempts, they lose
// 	fmt.Printf("Sorry, you've used all %d attempts. The correct number was %d.\n", maxAttempts, target)
// }

//****************************************************************************************************************************************

// package main

// import "fmt"

// // Function to provide hints based on the user's guess
// func getHint(target, guess int) string {
//     if guess == target {
//         return "Correct"
//     } else if guess > target {
//         return "Too high"
//     } else {
//         return "Too low"
//     }
// }

// func main() {
//     var targetnumber int
//     var noofattempts int
//     var userinput int

//     // Step 1: Ask for the target number and number of attempts
//     fmt.Println("Enter target number: ")
//     fmt.Scanln(&targetnumber)

//     fmt.Println("Enter number of attempts: ")
//     fmt.Scanln(&noofattempts)

//     // Step 2: Loop to allow user to guess
//     for i := noofattempts; i > 0; i-- {
//         fmt.Println("Enter your guess: ")
//         fmt.Scanln(&userinput)

//         // Step 3: Get the hint from the getHint function
//         hint := getHint(targetnumber, userinput)
        
//         // Check if the guess is correct, too high, or too low
//         if hint == "Correct" {
//             fmt.Println("Correct number, you win the game!")
//             break // End the loop if the guess is correct
//         } else {
//             fmt.Println(hint)
//         }

//         // Step 4: Let the user know how many attempts are left
//         if i-1 > 0 {
//             fmt.Printf("You have %d attempts left.\n", i-1)
//         }
//     }

//     // Step 5: If the user runs out of attempts without guessing correctly
//     if userinput != targetnumber {
//         fmt.Println("You reached the maximum number of attempts. You lost the game.")
//     }
// }

