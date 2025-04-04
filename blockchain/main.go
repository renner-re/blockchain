package main

import (
	"fmt"
)

func main(){
	NewBlockchain := NewBlockchain() // init blockchain
	// create 2 blocks and add 2 transactions
	NewBlockchain.AddBlock("First transaction")
	NewBlockchain.AddBlock("Second transaction")
	// print blocks and contents
	for i , block := range NewBlockchain.Blocks { // iterate each block
		fmt.Printf("Block ID: %d, \n", i)
		fmt.Printf("Timestamp: %d, \n", block.Timestamp+int64(i))
		fmt.Printf("hash of block %x\n", block.MyBlockHash)
		fmt.Printf("hash of previous block: %x\n", block.PreviousBlockHash)
		fmt.Printf("all transactions: %x\n", block.AllData)
	}
}
