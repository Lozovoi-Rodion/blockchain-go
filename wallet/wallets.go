package wallet

const walletFile = "./tmp/wallets.data"

type Wallets struct {
	Wallet map[string]*Wallet
}