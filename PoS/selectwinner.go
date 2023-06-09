package main

import (
	"errors"
	math "math/rand"
)

func (n PoSNetwork) SelectWinner() (*Node, error) {
	var winnerPool []*Node
	totalStake := 0
	for _, node := range n.Validators {
		if node.Stake > 0 {
			winnerPool = append(winnerPool, node)
			totalStake += node.Stake
		}
	}
	if winnerPool == nil {
		return nil, errors.New("There are no nodes with stake in the network")
	}
	winnerNumber := math.Intn(totalStake)
	tmp := 0
	for _, node := range n.Validators {
		tmp += node.Stake
		if winnerNumber < tmp {
			return node, nil
		}
	}
	return nil, errors.New("A winner should have been picked but wasn't")
}
