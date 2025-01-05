package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Struct to store a question
type Question struct {
	Question string
	Options  [4]string
	Answer   int
}

// Function to take the quiz
func takeQuiz(questions []Question) int {
	score := 0
	reader := bufio.NewReader(os.Stdin)

	for i, q := range questions {
		fmt.Printf("Question %d: %s\n", i+1, q.Question)
		for j, option := range q.Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}

		// Timer setup
		answerCh := make(chan int)
		go func() {
			for {
				fmt.Print("Enter your answer (or type 'exit' to quit): ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if strings.ToLower(input) == "exit" {
					answerCh <- -1
					return
				}

				answer, err := strconv.Atoi(input)
				if err != nil || answer < 1 || answer > 4 {
					fmt.Println("Invalid input. Please enter a number between 1 and 4.")
					continue
				}

				answerCh <- answer
				return
			}
		}()

		select {
		case answer := <-answerCh:
			if answer == -1 {
				fmt.Println("Exiting the quiz...")
				return score
			}

			if answer == q.Answer {
				fmt.Println("Correct!\n")
				score++
			} else {
				fmt.Printf("Wrong! The correct answer was %d.\n\n", q.Answer)
			}

		case <-time.After(15 * time.Second):
			fmt.Println("Time's up for this question!\n")
		}
	}

	return score
}

// Function to classify performance
func classifyPerformance(score, total int) string {
	percentage := (float64(score) / float64(total)) * 100

	switch {
	case percentage >= 80:
		return "Excellent"
	case percentage >= 50:
		return "Good"
	default:
		return "Needs Improvement"
	}
}

func main() {
	// Question bank
	questions := []Question{
		{"What is the capital of India?", [4]string{"Mumbai", "Delhi", "Kolkata", "Chennai"}, 2},
		{"Which programming language is known as Go?", [4]string{"Python", "C++", "Golang", "Ruby"}, 3},
		{"What is the value of Pi (approximately)?", [4]string{"3.14", "2.71", "1.62", "1.41"}, 1},
		{"Who wrote 'Hamlet'?", [4]string{"Charles Dickens", "William Shakespeare", "Mark Twain", "J.K. Rowling"}, 2},
	}

	fmt.Println("Welcome to the Online Examination System!")
	fmt.Println("You have 15 seconds to answer each question.")
	fmt.Println("Type 'exit' anytime to quit the quiz.\n")

	// Take the quiz
	score := takeQuiz(questions)
	total := len(questions)

	// Display results
	fmt.Printf("\nQuiz completed!\n")
	fmt.Printf("Your score: %d/%d\n", score, total)
	fmt.Printf("Performance: %s\n", classifyPerformance(score, total))
}
