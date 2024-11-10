package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		logrus.Error(err)
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		logrus.Error(err)
	}

	fmt.Println("header.Number===========>", header.Number.String())

	blockNumber := big.NewInt(5671744)

	block, err := client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		logrus.Error(err)
	}

	fmt.Println("block.Number=============>", block.Number().Uint64())                 // 5671744
	fmt.Println("block.Time==============>", block.Time())                             // 1527211625
	fmt.Println("block.Difficulty===============>", block.Difficulty().Uint64())       // 3217000136609065
	fmt.Println("block.Hash=================>", block.Hash().Hex())                    // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println("block.Transactions.len=================>", len(block.Transactions())) // 144

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println("count============>", count)
}
