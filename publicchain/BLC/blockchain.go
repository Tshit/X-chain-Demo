package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)
const dbFile = "blockchain.db"

const blocksBucket = "blocks"

type Blockchain struct {
	Tip []byte
	DB *bolt.DB

}

func (blockchain *Blockchain) AddBlock(data string) {
	preBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data, preBlock.PrevBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	//get the hash value of thr last block 
	var tip []byte 

	//open or create DB
	//create table if not exists
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b== nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil
			log.Panic(err)
			err = b.Put([]byte("l"),genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			tip = genesisBlock.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	}
	if err != nil {
		log.Panic(err)
	}


	return &Blockchain{tip,db}
}
