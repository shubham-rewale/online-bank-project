package bank

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}
	if a.Balance < amount {
		return errors.New("the amount to withdraw should be less than account balance")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) TransferFunds(amount float64, toAccount *Account) error {
	if err := a.Withdraw(amount); err != nil {
		return err
	}

	if err := toAccount.Deposit(amount); err != nil {
		if err := a.Deposit(amount); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}
