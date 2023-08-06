package block

// Block is a struct that represents a block in the blockchain.
type Block struct {
	header       Header
	transactions []Transaction
}

// Header is a struct that represents a header of block in the blockchain.
type Header struct {
	Hash          string `json:"hash"`
	PrevHash      string `json:"prev_hash"`
	Data          string `json:"data"`
	Nonce         string `json:"nonce"`
	MerkleRoot    string `json:"merkle_root"`
	Height        int    `json:"height"`
	Confirmations int    `json:"confirmations"`
	Difficulty    int    `json:"difficulty"`
	Reward        int    `json:"reward"`
	FeeReward     int    `json:"fee_reward"`
}
