package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/gochain-io/gochain/ethclient"
)

func main() {
	ctx := context.Background()
	url := "http://159.65.231.243:8545"
	numBlocks := 25
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	latestBlockNum, err := client.LatestBlockNumber(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(latestBlockNum)
	bnum64 := latestBlockNum.Int64()
	prevBlockTimestamp := time.Now().Unix()
	for i := bnum64; i >= 1 && i > (bnum64-int64(numBlocks)); i-- {
		block, err := client.BlockByNumber(ctx, big.NewInt(i))
		if err != nil {
			log.Fatal(err)
		}
		txs := block.Transactions()
		btimestamp := block.Time().Int64()
		blockTime := prevBlockTimestamp - btimestamp
		fmt.Println("blocknum:", i, "txcount:", len(txs), "blocktime:", blockTime)
		prevBlockTimestamp = btimestamp
	}
}
