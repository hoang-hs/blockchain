package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

const (
	baseFormatInt = 10
)

type Block struct {
	Timestamp     int64 //time when creat block
	Data          []byte
	PrevBlockHash []byte //hash of prev block
	Hash          []byte //hash of this block
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	block.SetHash()
	return block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, baseFormatInt))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
