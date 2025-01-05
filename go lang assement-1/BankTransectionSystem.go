package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID              int
	Name            string
	Balance         float64
	TransactionHist []string
}

var accounts []Account

// Find account by ID
func findAccountByID(id int) (*Account, error) {
	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

// Deposit money
func deposit(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	account, err := findAccountByID(id)
	if err != nil {
		return err
	}

	account.Balance += amount
	account.TransactionHist = append(account.TransactionHist, fmt.Sprintf("Deposited: $%.2f", amount))
	return nil
}

// Withdraw money
func withdraw(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}

	account, err := findAccountByID(id)
	if err != nil {
		return err
	}

	if account.Balance < amount {
		return errors.New("insufficient balance")
	}

	account.Balance -= amount
	account.TransactionHist = append(account.TransactionHist, fmt.Sprintf("Withdrew: $%.2f", amount))
	return nil
}

// View transaction history
func viewTransactionHistory(id int) ([]string, error) {
	account, err := findAccountByID(id)
	if err != nil {
		return nil, err
	}
	return account.TransactionHist, nil
}

func main() {
	const (
		DepositOption       = 1
		WithdrawOption      = 2
		ViewBalanceOption   = 3
		ViewHistoryOption   = 4
		ExitOption          = 5
	)

	accounts = append(accounts, Account{ID: 1, Name: "John Doe", Balance: 1000.0})
	accounts = append(accounts, Account{ID: 2, Name: "Jane Smith", Balance: 500.0})

	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice, id int
		fmt.Scanln(&choice)

		if choice == ExitOption {
			fmt.Println("Exiting...")
			break
		}

		fmt.Print("Enter Account ID: ")
		fmt.Scanln(&id)

		switch choice {
		case DepositOption:
			fmt.Print("Enter deposit amount: ")
			var amount float64
			fmt.Scanln(&amount)
			if err := deposit(id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful")
			}
		case WithdrawOption:
			fmt.Print("Enter withdrawal amount: ")
			var amount float64
			fmt.Scanln(&amount)
			if err := withdraw(id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful")
			}
		case ViewBalanceOption:
			account, err := findAccountByID(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Balance: $%.2f\n", account.Balance)
			}
		case ViewHistoryOption:
			history, err := viewTransactionHistory(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Transaction History:")
				for _, record := range history {
					fmt.Println(record)
				}
			}
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
