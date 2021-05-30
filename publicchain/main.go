package main

import (
	"fmt"
	"publicchain/part1/BLC"
	"time"
)

func main() {

	blockchain := BLC.NewBlockchain()

	blockchain.AddBlock("Send 20 BTC To HaoLin From Liyuechun")

	blockchain.AddBlock("Send 10 BTC To SaoLin From Liyuechun")

	blockchain.AddBlock("Send 30 BTC To HaoTian From Liyuechun")

	for _, block := range blockchain.Blocks {

		fmt.Printf("Data：%s \n", string(block.Data))
		fmt.Printf("PrevBlockHash：%x \n", block.PrevBlockHash)
		fmt.Printf("Timestamp：%s \n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash：%x \n", block.Hash)

		fmt.Println()
	}
}
