package main

import (
	"errors"
	"fmt"

	"example.com/bank/fileops"
)

type Bank struct {
	balance     float64
	balanceFile string
}

func NewBank(balanceFile string) (*Bank, error) {
	balance, err := fileops.ReadValueFromFile(balanceFile)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize bank: %w", err)
	}

	return &Bank{
		balance:     balance,
		balanceFile: balanceFile,
	}, nil
}

func (b *Bank) GetBalance() float64 {
	return b.balance
}

func (b *Bank) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}

	b.balance += amount

	if err := b.saveBalance(); err != nil {
		b.balance -= amount // rollback
		return fmt.Errorf("failed to save balance: %w", err)
	}

	return nil
}

func (b *Bank) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}

	if amount > b.balance {
		return errors.New("insufficient balance")
	}

	b.balance -= amount

	if err := b.saveBalance(); err != nil {
		b.balance += amount // rollback
		return fmt.Errorf("failed to save balance: %w", err)
	}

	return nil
}

func (b *Bank) saveBalance() error {
	return fileops.WriteValueToFile(b.balance, b.balanceFile)
}
