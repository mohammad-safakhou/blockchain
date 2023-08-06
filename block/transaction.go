package block

import (
	"github.com/mohammad-safakhou/blockchain/account"
	"github.com/mohammad-safakhou/blockchain/config"
	"github.com/mohammad-safakhou/blockchain/utils"
)

// Transaction is a struct that represents a transaction in the blockchain.
type Transaction struct {
	Hash        string          `json:"hash"`
	Input       account.Address `json:"input"`
	Output      account.Address `json:"output"`
	Description string          `json:"description"`
	Amount      int             `json:"amount"`
	Fee         int             `json:"fee"`
}

func (t *Transaction) Validate() bool {
	return utils.CheckStringLength(t.Description, config.Cfg.Transaction.MaxDescriptionLength)
}
