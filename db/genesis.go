package db

type Genesis struct {
	Hash     string `json:"hash"`
	PrevHash string `json:"prev_hash"`
	Data     string `json:"data"`
	Nonce    string `json:"nonce"`
}
