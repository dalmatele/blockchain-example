package main

import "errors"

func (n PoSNetwork) ValidateBlockchain() error {
	if len(n.Blockchain) <= 1 {
		return nil
	}
	currBlockIdx := len(n.Blockchain) - 1
	prevBlockIdx := len(n.Blockchain) - 2
	for prevBlockIdx >= 0 {
		currBlock := n.Blockchain[currBlockIdx]
		prevBlock := n.Blockchain[prevBlockIdx]
		if currBlock.PrevHash != prevBlock.Hash {
			return errors.New("Blockchain has incossistent hashes")
		}
		if currBlock.Timestamp <= prevBlock.Timestamp {
			return errors.New("Blockchain has inconsistent timestamps")
		}
		if NewBlockHash(prevBlock) != currBlock.Hash {
			return errors.New("Blockchain has inconsistent hash genration")
		}
		currBlockIdx--
		prevBlockIdx--
	}
	return nil
}

func (n PoSNetwork) ValidateBlockCandidate(newBlock *Block) error {
	if n.BlockchainHead.Hash != newBlock.PrevHash {
		return errors.New("Blockchain HEAD hash is not equal to new block previous hash")
	}
	if n.BlockchainHead.Timestamp >= newBlock.Timestamp {
		return errors.New("Block HEAD timestamp is greater than or equal to new block")
	}
	if NewBlockHash(n.BlockchainHead) != newBlock.Hash {
		return errors.New("New block hash of blockchain HEAD does not equal new block hash")
	}
	return nil
}
