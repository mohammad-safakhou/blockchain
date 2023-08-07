package utils

import (
	"crypto/sha256"
	"fmt"
)

func CheckStringLength(s string, length int) bool {
	return len(s) == length
}

func DoubleHash(s string) string {
	// Calculate the double SHA256 hash of the combined data
	hash := sha256.Sum256([]byte(s))
	hash = sha256.Sum256(hash[:])

	// Convert the hash to a hexadecimal string
	return fmt.Sprintf("%x", hash)
}

func CalculateMerkleRoot(hashes []string) string {
	// If the slice of hashes is empty, return an empty string
	if len(hashes) == 0 {
		return ""
	}

	// If the slice of hashes contains only one hash, return that hash
	if len(hashes) == 1 {
		return hashes[0]
	}

	// If the slice of hashes contains an odd number of hashes, duplicate the last hash
	if len(hashes)%2 == 1 {
		hashes = append(hashes, hashes[len(hashes)-1])
	}

	// Create a slice of strings to store the new hashes
	var newHashes []string

	// Iterate over the hashes and add the combined hashes to the slice
	for i := 0; i < len(hashes); i += 2 {
		newHashes = append(newHashes, DoubleHash(hashes[i]+hashes[i+1]))
	}

	// Calculate the Merkle root of the new hashes
	return CalculateMerkleRoot(newHashes)
}
