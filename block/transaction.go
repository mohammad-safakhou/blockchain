package block

type Transaction struct {
	Hash        string  `json:"hash"`
	Input       Address `json:"input"`
	Output      Address `json:"output"`
	Description string  `json:"description"`
	Amount      int     `json:"amount"`
}
