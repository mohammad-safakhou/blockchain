package block

import (
	"encoding/json"
	"github.com/mohammad-safakhou/blockchain/account"
	"github.com/mohammad-safakhou/blockchain/config"
	"github.com/mohammad-safakhou/blockchain/utils"
)

type Transactions []Transaction

func (t *Transactions) MerkleRoot() string {
	// Create a slice of strings to store the transaction hashes
	var transactionHashes []string

	// Iterate over the Transactions and add the transaction hashes to the slice
	for _, transaction := range *t {
		transactionHashes = append(transactionHashes, transaction.Hash())
	}

	// Calculate the Merkle root of the transaction hashes
	return utils.CalculateMerkleRoot(transactionHashes)
}

// Transaction is a struct that represents a transaction in the blockchain.
type Transaction struct {
	Input       account.Address `json:"input"`
	Output      account.Address `json:"output"`
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
	Fee         float64         `json:"fee"`
}

func (t *Transaction) Validate() bool {
	return utils.CheckStringLength(t.Description, config.Cfg.Transaction.MaxDescriptionLength)
}

func (t *Transaction) Hash() string {
	dh, _ := json.Marshal(t)
	return utils.DoubleHash(string(dh))
}
