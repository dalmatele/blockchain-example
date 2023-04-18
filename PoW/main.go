package main

import (
	"fmt"
)

func main() {
	newblockchain := NewBlockChain()
	newblockchain.AddBlock("First transaction")
	newblockchain.AddBlock("Second transaction")
	for i, block := range newblockchain.Blocks {
		fmt.Printf("Block ID: %d \n", i)
		fmt.Printf("Timestamp: %d \n", block.Timestamp+int64(i))
		fmt.Printf("Hash of the block %x\n", block.MyBlockHash)
		fmt.Printf("Hash of the previous Block: %x\n", block.PreviousBlockhash)
		fmt.Printf("Current transaction: %s\n", block.AllData)
		fmt.Printf("============================\n")
	}
}
