package account

import "fmt"

type Account struct {
	ID          string
	AccountName string
	Balance     float64
}

type IAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
}

type PrintableToConsole interface {
	PrintToConsole()
}

func (acc *Account) Deposit(amount float64) error {
	if amount < 0.0 {
		return fmt.Errorf("could not deposit negative amount")
	}

	acc.Balance += amount
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount > acc.Balance {
		return fmt.Errorf("withdraw amount is larger than account balance")
	}

	acc.Balance -= amount
	return nil
}

func (acc *Account) PrintToConsole() {
	fmt.Printf("ID: %s\n", acc.ID)
	fmt.Printf("Account Name: %s\n", acc.AccountName)
	fmt.Printf("Balance: %f\n", acc.Balance)
}
