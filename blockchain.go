package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// Block represents a single block in the blockchain.
type Block struct {
	Transaction  string // The data or transaction stored in the block.
	Nonce        int    // A random number used in the proof of work process.
	PreviousHash string // The hash of the previous block in the chain.
	CurrentHash  string // The hash of the current block.
}

// NewBlock creates a new block with the given transaction and previous hash.
func NewBlock(transaction string, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		PreviousHash: previousHash,
		Nonce:        0, // Initialize nonce to 0
	}
	block.CurrentHash = CalculateHash(block)
	return block
}

// DisplayBlocks prints information about each block in the blockchain.
func DisplayBlocks(blocks []*Block) {
	for i, block := range blocks {
		fmt.Printf("Block %d:\n", i+1)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.CurrentHash)
		fmt.Println()
	}
}

// ChangeBlock updates a block with a new transaction and resets the nonce.
func ChangeBlock(block *Block, newTransaction string) {
	block.Transaction = newTransaction
	block.Nonce = 0 // Reset nonce
	block.CurrentHash = CalculateHash(block)
}

// VerifyChain checks if the blockchain is valid by verifying previous hashes and proof of work.
func VerifyChain(blocks []*Block) bool {
	for i := 1; i < len(blocks); i++ {
		currentBlock := blocks[i]
		previousBlock := blocks[i-1]

		// Check if the previous hash of the current block matches the current hash of the previous block.
		if currentBlock.PreviousHash != previousBlock.CurrentHash {
			return false
		}

		// Check if the current block satisfies the proof of work requirement.
		if !ValidateProofOfWork(currentBlock) {
			return false
		}
	}
	return true
}

// CalculateHash computes the hash of a block using its transaction, nonce, and previous hash.
func CalculateHash(block *Block) string {
	data := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// ValidateProofOfWork checks if the block's current hash satisfies the proof of work requirement.
func ValidateProofOfWork(block *Block) bool {
	// In this example, we require at least 2 leading zeros in the hash.
	return strings.HasPrefix(block.CurrentHash, "00")
}
