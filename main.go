package main

import (
	"fmt"
	"log"
)

const balanceFile string = "balance.txt"

func main() {
	bank, err := NewBank(balanceFile)
	if err != nil {
		log.Fatalf("Failed to initialize bank: %v\n", err)
	}

	fmt.Println("Hello, Welcome!")

	for {
		showOptions()

		choice := getUserChoice()

		switch choice {
		case 1:
			handleCheckBalance(bank)
		case 2:
			handleWithdraw(bank)
		case 3:
			handleDeposit(bank)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}
}

func getUserChoice() int {
	var choice int
	fmt.Print(": ")
	fmt.Scan(&choice)
	return choice
}

func handleCheckBalance(bank *Bank) {
	fmt.Printf("Your current balance: %.2f\n", bank.GetBalance())
}

func handleWithdraw(bank *Bank) {
	var amount float64
	fmt.Print("Enter an amount: ")
	fmt.Scan(&amount)

	if err := bank.Withdraw(amount); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("You have withdrawn %.2f. New balance is %.2f\n", amount, bank.GetBalance())
}

func handleDeposit(bank *Bank) {
	var amount float64
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)

	if err := bank.Deposit(amount); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("You have deposited %.2f. New balance is %.2f\n", amount, bank.GetBalance())
}
