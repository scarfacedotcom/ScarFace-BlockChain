package main

import (
	"fmt"
	"log"

	"github.com/scarface-blockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	// myBlockchainAddress := "my_blockchain_address"

	// blockchain := NewBlockchain(myBlockchainAddress)
	// blockchain.Print()

	// blockchain.AddTransaction("Peter", "Jay", 1.0)
	// blockchain.Mining()
	// blockchain.Print()

	// blockchain.AddTransaction("ScarFace", "Mark", 2.0)
	// blockchain.AddTransaction("Alice", "Bob", 3.0)
	// blockchain.Mining()
	// blockchain.Print()

	// fmt.Printf("my_blockchain_address %.1f\n", blockchain.CalculateTotalAmount("my_blockchain_address"))
	// fmt.Printf("ScarFace %.1f\n", blockchain.CalculateTotalAmount("ScarFace"))
	// fmt.Printf("Mark %.1f\n", blockchain.CalculateTotalAmount("Mark"))

	w := wallet.NewWallet()
	fmt.Println("Private Key", w.PrivateKeyStr())
	fmt.Println("Public Key", w.PublicKeyStr())
	fmt.Println("Wallet Address", w.BlockchainAddress())
}
