package quickcash

type CreditCardAccount struct {
	balance float64
}

func (cca *CreditCardAccount) WithDraw(amount float64) {
	cca.balance -= amount
}

func (cca *CreditCardAccount) CanWithDraw(amount float64) bool {
	if cca.balance < amount {
		return false
	}
	return true
}

func (cca *CreditCardAccount) GetIdentifier() string {
	return "CREDIT_CARD_ACCOUNT"
}
