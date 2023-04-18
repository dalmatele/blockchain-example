package main

type PoSNetwork struct {
	Blockchain     []*Block
	BlockchainHead *Block // the latest block
	Validators     []*Node
}

type Node struct {
	Stake   int
	Address string
}

type Block struct {
	Timestamp     string
	PrevHash      string //hash of previous block
	Hash          string //hash of this block
	ValidatorAddr string
}
