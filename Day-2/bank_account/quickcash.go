package quickcash

type QuickCash struct {
	SavingsAccount    Withdrawable
	CreditCardAccount Withdrawable
	Wallet            Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	if qc.Wallet.CanWithDraw((amount)) {
		qc.Wallet.WithDraw(amount)
		return amount, qc.Wallet.GetIdentifier()
	}
	if qc.SavingsAccount.CanWithDraw(amount) {
		qc.SavingsAccount.WithDraw(amount)
		return amount, qc.SavingsAccount.GetIdentifier()
	}
	qc.CreditCardAccount.WithDraw(amount)
	return amount, qc.CreditCardAccount.GetIdentifier()
}
