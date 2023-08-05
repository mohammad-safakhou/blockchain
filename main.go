// package main
//
// import (
//
//	"crypto/sha256"
//	"fmt"
//
// )
//
//	type Block struct {
//		Hash     string `json:"hash"`
//		PrevHash string `json:"prev_hash"`
//		Data     string `json:"data"`
//	}
//
// var Blockchain = []Block{Block{Hash: "genesis hash", PrevHash: "", Data: "first data block"}}
//
//	func main() {
//		// simple blockchain implementation
//
//		// 1. create genesis block
//		// 2. create new block
//		// 3. add new block to blockchain
//		// 4. print blockchain
//		// 5. validate blockchain
//		// 6. change data in block
//		// 7. validate blockchain
//
//		AddingBlock("second data block")
//		fmt.Printf("blockchain is valid: %v\n", Validate())
//		AddingBlock("third data block")
//		fmt.Printf("blockchain is valid: %v\n", Validate())
//		AddingBlock("fourth data block")
//		fmt.Printf("blockchain is valid: %v\n", Validate())
//
//		fmt.Printf("blockchain: %v\n", Blockchain)
//	}
//
//	func AddingBlock(data string) {
//		hashes := sha256.New()
//		hashes.Write([]byte(Blockchain[len(Blockchain)-1].Hash + ":%:" + data))
//		hash := hashes.Sum(nil)
//		Blockchain = append(Blockchain, Block{Hash: string(hash), PrevHash: Blockchain[len(Blockchain)-1].Hash, Data: data})
//	}
//
//	func Validate() bool {
//		for i := 1; i < len(Blockchain); i++ {
//			if Blockchain[i].PrevHash != Blockchain[i-1].Hash {
//				return false
//			}
//		}
//		return true
//	}
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"
)

func calculateBlockHash(header string, nonce uint64) string {
	// Combine the block header and nonce
	data := header + fmt.Sprintf("%08x", nonce)

	// Calculate the double SHA256 hash of the combined data
	hash := sha256.Sum256([]byte(data))
	hash = sha256.Sum256(hash[:])

	// Convert the hash to a hexadecimal string
	blockHash := fmt.Sprintf("%x", hash)

	return blockHash
}

func main() {
	start := time.Now()
	// Example block header (simplified version)
	previousBlockHash := "0000000000000000000000000000000000000000000000000000000000000000"
	merkleRoot := "e0c6beaf317633e5466f142aa8f7edeba138c7f0c79e79b804ff9bcdc8a34a37"
	timestamp := "2023-07-31 12:00:00"
	targetDifficulty := "000000" // The required number of leading zeros in the block hash

	// Concatenate the header fields
	blockHeader := previousBlockHash + merkleRoot + timestamp + targetDifficulty

	const numMiners = 100
	winner := make(chan int)
	var c = make([]int, numMiners)
	for i := 0; i < numMiners; i++ {
		go func(ii int) {
			// Mining loop
			for { // 2
				c[ii]++
				// Generate a random 32-bit integer (nonce)
				nonce, err := rand.Int(rand.Reader, big.NewInt(4294967296))
				if err != nil {
					panic(err)
				}
				blockHash := calculateBlockHash(blockHeader, nonce.Uint64())
				if blockHash[:len(targetDifficulty)] == targetDifficulty {
					// Nonce found! Block hash meets the difficulty requirement.
					fmt.Printf("Block Hash: %s\n", blockHash)
					fmt.Printf("Nonce: %d\n", nonce)
					winner <- ii
					break
				}
			}
		}(i)
	}

	winnerNumber := <-winner
	t := 0
	for i := 0; i < numMiners; i++ {
		fmt.Printf("miner %d: %d attempts\n", i, c[i])
		t += c[i]
	}
	fmt.Printf("miner %d won\n", winnerNumber)
	fmt.Printf("total attempts: %d\n", t)

	//fmt.Printf("nonce found after %d attempts\n", c)

	fmt.Printf("process time took: %dms\n", time.Now().Sub(start).Milliseconds())
}
