package main

import (
	"bytes"
	"encoding/gob"
	"go_blockchain/log"
	"time"
)

type Block struct {
	Timestamp     int64 //time when creat block
	Data          []byte
	PrevBlockHash []byte //hash of prev block
	Hash          []byte //hash of this block
	Nonce         int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return result.Bytes()
}
