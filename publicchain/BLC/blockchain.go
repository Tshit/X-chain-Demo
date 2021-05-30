package BLC

type Blockchain struct {
	Blocks []*Block
}

func (blockchain *Blockchain) AddBlock(data string) {
	preBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data, preBlock.PrevBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
