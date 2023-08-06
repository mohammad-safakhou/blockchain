package account

import (
	"crypto"
)

type Account struct {
	Addr      Address          `json:"addr"`
	PublicKey crypto.PublicKey `json:"public_key"`
	Balance   int              `json:"balance"`
}

func New(pub crypto.PublicKey) Account {
	address := GenerateAddress(pub)
	return Account{
		Addr:      address,
		PublicKey: pub,
		Balance:   0,
	}
}

type Address string

func (a *Address) ToString() string {
	return string(*a)
}

func GenerateAddress(pub crypto.PublicKey) Address {
	return Address(pub.(crypto.Hash).String())
}
