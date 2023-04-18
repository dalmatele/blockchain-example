package main

import (
	"fmt"
	"log"
	math "math/rand"
	"time"
)

func main() {
	math.Seed(time.Now().UnixNano())
	genesisTime := time.Now().String()
	pos := &PoSNetwork{
		Blockchain: []*Block{
			{
				Timestamp:     genesisTime,
				PrevHash:      "",
				Hash:          newHash(genesisTime),
				ValidatorAddr: "",
			},
		},
	}
	pos.BlockchainHead = pos.Blockchain[0]
	pos.Validators = pos.NewNode(60)
	pos.Validators = pos.NewNode(40)
	for i := 0; i < 5; i++ {
		winner, err := pos.SelectWinner()
		if err != nil {
			log.Fatal(err)
		}
		winner.Stake += 10
		pos.Blockchain, pos.BlockchainHead, err = pos.GenerateNewBlock(winner)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Round ", i)
		fmt.Println("\tAddress: ", pos.Validators[0].Address, "-Stake:", pos.Validators[0].Stake)
		fmt.Println("\tAddress: ", pos.Validators[1].Address, "-Stake:", pos.Validators[1].Stake)
	}
	pos.PrintBlockchainInfo()
}

func (n PoSNetwork) PrintBlockchainInfo() {
	for i := 0; i < len(n.Blockchain); i++ {
		fmt.Println("Block ", i, "Info:")
		fmt.Println("Timestamp:", n.Blockchain[i].Timestamp)
		fmt.Println("Previous Hash:", n.Blockchain[i].PrevHash)
		fmt.Println("Hash:", n.Blockchain[i].Hash)
		fmt.Println("Validator address:", n.Blockchain[i].ValidatorAddr)
	}
}
