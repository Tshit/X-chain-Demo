package main

import (
	"fmt"
	"publicchain/BLC"
	"time"
)

func main() {

	blockchain := BLC.NewBlockchain()

	blockchain.AddBlock("Send 20 BTC To Ben From Admin")

	blockchain.AddBlock("Send 10 BTC To Sam From Admin")

	blockchain.AddBlock("Send 30 BTC To Alice From Admin")

	for _, block := range blockchain.Blocks {

		fmt.Printf("Data：%s \n", string(block.Data))
		fmt.Printf("PrevBlockHash：%x \n", block.PrevBlockHash)
		fmt.Printf("Timestamp：%s \n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash：%x \n", block.Hash)

		fmt.Println()
	}
}
