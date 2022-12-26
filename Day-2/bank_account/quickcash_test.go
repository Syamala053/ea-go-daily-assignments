package quickcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCashFromSavingsAccount(t *testing.T) {
	sa := &SavingsAccount{balance: 600}
	cca := &CreditCardAccount{balance: 1000}
	pw := &PaytmWallet{balance: 100}

	qc := QuickCash{
		sa,
		cca,
		pw,
	}

	amt, accType := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, sa.GetIdentifier(), accType)
}

func TestGetCashFromCreditCardAccount(t *testing.T) {

	sa := &SavingsAccount{balance: 100}
	cca := &CreditCardAccount{balance: 1000}
	pw := &PaytmWallet{balance: 100}

	qc := QuickCash{
		sa,
		cca,
		pw,
	}

	amt, accType := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, cca.GetIdentifier(), accType)
}

func TestGetCashFromWallet(t *testing.T) {

	sa := &SavingsAccount{balance: 100}
	cca := &CreditCardAccount{balance: 1000}
	pw := &PaytmWallet{balance: 500}

	qc := QuickCash{
		sa,
		cca,
		pw,
	}

	amt, accType := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, pw.GetIdentifier(), accType)
}
