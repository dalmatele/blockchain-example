package main

type Block struct {
	Timestamp         int64
	PreviousBlockhash []byte
	MyBlockHash       []byte
	AllData           []byte
}

type Blockchain struct {
	Blocks []*Block
}
