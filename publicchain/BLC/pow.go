package BLC

import (
	"math/big"
	"fmt"
	"bytes"
	"math"
	"crypto/sha256"
)

var (
	maxNonce = math.MaxInt64
)

type ProofOfWork struct {
	block *Block
	target *big.int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	
	target := big.NewInt(1)
	fmt.Println("---------------------")
	fmt.Println(target)
	pow := &ProofOfWork{block,target}

	return pow
} 

func (pow *ProofOfWork ) Run() {

}
