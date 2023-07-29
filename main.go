package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     string `json:"hash"`
	PrevHash string `json:"prev_hash"`
	Data     string `json:"data"`
}

var Blockchain = []Block{Block{Hash: "genesis hash", PrevHash: "", Data: "first data block"}}

func main() {
	// simple blockchain implementation

	// 1. create genesis block
	// 2. create new block
	// 3. add new block to blockchain
	// 4. print blockchain
	// 5. validate blockchain
	// 6. change data in block
	// 7. validate blockchain

	AddingBlock("second data block")
	fmt.Printf("blockchain is valid: %v\n", Validate())
	AddingBlock("third data block")
	fmt.Printf("blockchain is valid: %v\n", Validate())
	AddingBlock("fourth data block")
	fmt.Printf("blockchain is valid: %v\n", Validate())

	fmt.Printf("blockchain: %v\n", Blockchain)
}

func AddingBlock(data string) {
	hashes := sha256.New()
	hashes.Write([]byte(Blockchain[len(Blockchain)-1].Hash + ":%:" + data))
	hash := hashes.Sum(nil)
	Blockchain = append(Blockchain, Block{Hash: string(hash), PrevHash: Blockchain[len(Blockchain)-1].Hash, Data: data})
}

func Validate() bool {
	for i := 1; i < len(Blockchain); i++ {
		if Blockchain[i].PrevHash != Blockchain[i-1].Hash {
			return false
		}
	}
	return true
}
