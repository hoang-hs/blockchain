package main

import (
	"fmt"
	"go_blockchain/block"
)

func main() {
	bc := block.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, bl := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", bl.PrevBlockHash)
		fmt.Printf("Data: %s\n", bl.Data)
		fmt.Printf("Hash: %x\n", bl.Hash)
		fmt.Println()
	}
}
