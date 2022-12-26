package account

import "fmt"

type Account struct {
	balance float64
}

type InsufficientBalanceError struct {
	accountBalance   float64
	requestedBalance float64
}

func (ibe InsufficientBalanceError) Error() string {
	return fmt.Sprintf("Insufficient balance. Current balance is %.02f. %.02f amount was less.", ibe.accountBalance, ibe.requestedBalance-ibe.accountBalance)
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
}

func (acc *Account) Withdraw(amount float64) (float64, error) {
	if acc.balance < amount {
		return 0, InsufficientBalanceError{acc.balance, amount}
	}
	acc.balance -= amount
	return acc.balance, nil
}
