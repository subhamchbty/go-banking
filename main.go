package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const balanceFile string = "balance.txt"

func main() {
	fmt.Println("Hello, Welcome!")
	balance, _ := fileops.ReadValueFromFile(balanceFile)

	for {
		showOptions()

		var choice int
		fmt.Print(": ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// check balance
			balance, err := fileops.ReadValueFromFile(balanceFile)

			if err != nil {
				fmt.Printf("Error: %s.\n", err)
				continue
			}

			fmt.Printf("Your current balance: %.2f.\n", balance)
		case 2:
			// withdraw amount
			var withdrawalAmount float64
			fmt.Print("Enter an amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount < 0 {
				fmt.Println("Error: amount must be a positive.")
				continue
			}

			if withdrawalAmount > balance {
				fmt.Println("Error: amount must be less than current balance.")
				continue
			}

			balance -= withdrawalAmount

			err := fileops.WriteValueToFile(balance, balanceFile)

			if err != nil {
				fmt.Printf("Error: %s.\n", err)
				continue
			}

			fmt.Printf("You have withdrawn %.2f. New balance is %.2f.\n", withdrawalAmount, balance)
		case 3:
			// deposit amount
			var depositAmount float64

			fmt.Print("Enter amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount < 0 {
				fmt.Println("Error: amount should be a positive.")
				continue
			}

			balance += depositAmount

			err := fileops.WriteValueToFile(balance, balanceFile)

			if err != nil {
				fmt.Printf("Error: %s.\n", err)
				continue
			}

			fmt.Printf("You have deposited %.2f. New balance is %.2f.\n", depositAmount, balance)
		default:
			// exit
			fmt.Println("Goodbye!")
			return
		}
	}
}
