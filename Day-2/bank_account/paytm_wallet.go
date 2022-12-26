package quickcash

type PaytmWallet struct {
	balance float64
}

func (pw *PaytmWallet) WithDraw(amount float64) {
	pw.balance -= amount
}

func (pw *PaytmWallet) CanWithDraw(amount float64) bool {
	if pw.balance < amount {
		return false
	}
	return true
}

func (pw *PaytmWallet) GetIdentifier() string {
	return "PAYTM_WALLET"
}
