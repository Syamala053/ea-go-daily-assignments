package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	acc := Account{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestSuccessfulDeposit(t *testing.T) {
	acc := Account{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestSuccessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	acc.Withdraw(200)

	assert.Equal(t, float64(300), acc.GetBalance())
}

func TestInsufficientBalance(t *testing.T) {
	acc := Account{balance: 100}

	_, error := acc.Withdraw(300)

	insufficientBalanceError := error.Error()

	assert.Equal(t, insufficientBalanceError, "Insufficient balance. Current balance is 100.00. 200.00 amount was less.")
}
