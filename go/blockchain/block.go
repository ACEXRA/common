package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Data         string
	Hash         string
	PreviousHash string
	Timestamp    int64
}

type Blockchain struct {
	Blocks []*Block
}

func createNewBlock(data string, previousHash string) *Block {
	timestamp:= time.Now().Unix()
	hash := createHash(data,previousHash,timestamp)
	return &Block{data, hash, previousHash, timestamp}
}

func createHash(data string,previousHash string,timestamp int64) string {
	input := fmt.Sprintf("%s%s%d",data,previousHash,timestamp);
	hash := sha256.Sum256([]byte(input));
	return hex.EncodeToString(hash[:])
}

func createNewBlockChain() *Blockchain{
	return &Blockchain{[]*Block{genesisBlock()}}
}

func genesisBlock() *Block{
	return &Block{
		"Genesis block",
		"0",
		"genesisHash",
		time.Now().Unix(),
	}
}

func (blockchain *Blockchain) adddBlock(data string){
	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1];
	newBlock :=createNewBlock(data, prevBlock.Hash);
	blockchain.Blocks = append(blockchain.Blocks, newBlock);
}
