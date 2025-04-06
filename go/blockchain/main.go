package main

import "fmt"

func main() {
	blockChain := createNewBlockChain()

	blockChain.adddBlock("This is my first block")
	blockChain.adddBlock("This is my second block")

	for _,block := range blockChain.Blocks {
		fmt.Println("Data: ",block.Data)
		fmt.Println("Hash: ",block.Hash)
		fmt.Println("Previous Hash: ",block.PreviousHash)
		fmt.Println("-----------------------------------")
	}

}