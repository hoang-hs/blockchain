package main

import (
	"go_blockchain/log"
	"strconv"
)

func init() {
	l, err := log.NewLogger()
	if err != nil {
		panic(err)
	}
	log.RegisterGlobal(l)
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, bl := range bc.Blocks {
		log.Infof("Prev. hash: %x", bl.PrevBlockHash)
		log.Infof("Data: %s", bl.Data)
		log.Infof("Hash: %x", bl.Hash)
		pow := NewProofOfWork(bl)
		log.Infof("validate: %s\n", strconv.FormatBool(pow.Validate()))
	}

}
