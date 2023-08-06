package block

type Block struct {
	header       Header
	transactions []Transaction
}

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
