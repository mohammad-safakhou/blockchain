package block

import (
	"github.com/mohammad-safakhou/blockchain/utils"
	"math/big"
)

// Block is a struct that represents a block in the blockchain.
type Block struct {
	Header       Header
	Transactions []Transaction
}

// Hash returns the hash of the block.
func (b *Block) Hash() string {
	// Combine the block Header and nonce
	data := b.Header.Nonce + b.Header.Hash + b.Header.PrevHash + b.Header.Data + b.Header.MerkleRoot

	return utils.DoubleHash(data)
}

func NewBlock(header Header, transactions []Transaction) Block {
	return Block{
		Header:       header,
		Transactions: transactions,
	}
}

// Header is a struct that represents a Header of block in the blockchain.
type Header struct {
	Hash          string   `json:"hash"`
	PrevHash      string   `json:"prev_hash"`
	Timestamp     int64    `json:"timestamp"`
	Size          int      `json:"size"`
	Data          string   `json:"data"`
	Nonce         string   `json:"nonce"`
	MerkleRoot    string   `json:"merkle_root"`
	Height        *big.Int `json:"height"`
	Confirmations *big.Int `json:"confirmations"`
	Difficulty    *big.Int `json:"difficulty"`
	Bits          string   `json:"bits"`
	Reward        *big.Int `json:"reward"`
	FeeReward     *big.Int `json:"fee_reward"`
}

func calculateNewDifficulty(previousDifficulty *big.Int, targetTimeSeconds, actualTimeSeconds int64) *big.Int {
	// Convert the target times to big.Int for precision calculations
	targetTimeBigInt := big.NewInt(targetTimeSeconds)

	// Calculate the new difficulty adjustment factor
	diffAdjustmentFactor := new(big.Int).Mul(previousDifficulty, big.NewInt(actualTimeSeconds))
	diffAdjustmentFactor.Div(diffAdjustmentFactor, targetTimeBigInt)

	// Calculate the new difficulty by dividing the previous difficulty by the adjustment factor
	newDifficulty := new(big.Int).Div(previousDifficulty, diffAdjustmentFactor)

	// Ensure the new difficulty is within a valid range (minimum and maximum limits)
	minDifficulty := big.NewInt(1) // Minimum difficulty allowed
	maxDifficulty := new(big.Int).Lsh(minDifficulty, 224)
	if newDifficulty.Cmp(minDifficulty) < 0 {
		newDifficulty.Set(minDifficulty)
	} else if newDifficulty.Cmp(maxDifficulty) > 0 {
		newDifficulty.Set(maxDifficulty)
	}

	return newDifficulty
}
