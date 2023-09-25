package main

import (
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	prevHash     string
	timestamp    int64
	transactions []string
}

func (b *Block) Print() {
	fmt.Printf("timestamp  %d\n", b.timestamp)
	fmt.Printf("nonce  %d\n", b.nonce)
	fmt.Printf("prevvoius_hash  %s\n", b.prevHash)
	fmt.Printf("transactions  %s\n", b.transactions)

}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreatedBlock(0, "Init hash")
	return bc
}

func (bc *Blockchain) CreatedBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func NewBlock(nonce int, prevHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.prevHash = prevHash
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func main() {
	// b := NewBlock(0, "init Hash")
	// b.Print()
	blockchain := NewBlockchain()
	blockchain.Print()
	blockchain.CreatedBlock(5, "hash 1")
	blockchain.Print()
	blockchain.CreatedBlock(2, "hash 2")
	blockchain.Print()
}
