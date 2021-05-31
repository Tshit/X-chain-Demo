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
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	
	target := big.NewInt(1)
	fmt.Println("---------------------")
	target.Lsh(target, uint(256-targetBits))
	fmt.Println(target)
	pow := &ProofOfWork{block,target}
	return pow
} 

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork ) Run() (int, []byte) {
	bytes := 
}