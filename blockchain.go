package Asssignment01bca

import (
	"crypto/sha256" // Importing the sha256 library for hashing
	"fmt"           // Importing the fmt library for printing
)

// Defining the block structure for the blockchain
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

// Function to add a new block to the blockchain
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CurrentHash = CalculateHash(block)
	return block
}

// Function to display all blocks of the blockchain
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

// Function to change an already added block of the blockchain
func ChangeBlock(block *Block, newTransaction string) {
	block.Transaction = newTransaction
	block.CurrentHash = CalculateHash(block)
}

// Function to verify the blockchain for any changes
func VerifyChain(blocks []*Block) bool {
	for i := 1; i < len(blocks); i++ {
		currentBlock := blocks[i]
		previousBlock := blocks[i-1]

		if currentBlock.PreviousHash != previousBlock.CurrentHash {
			return false
		}

		if currentBlock.CurrentHash != CalculateHash(currentBlock) {
			return false
		}
	}
	return true
}

// Function to calculate the hash of a block
func CalculateHash(block *Block) string {
	data := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
