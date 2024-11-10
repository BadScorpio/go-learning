package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		logrus.Error(err)
	}
	// 通过块ID获取到具体的区块，通过区块的Transactions可以查询到交易信息
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		logrus.Error(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())
		fmt.Println(tx.Data())
		fmt.Println(tx.To().Hex())

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			logrus.Error(err)
		}

		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex())
		}
		// 获取事物的收据，其中包含事物的结果，返回值，日志
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			logrus.Error(err)
		}

		fmt.Println(receipt.Status)
	}

	// 通过块哈希和事物索引获取事物
	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		logrus.Error(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			logrus.Error(err)
		}

		fmt.Println(tx.Hash().Hex())
	}

	// 通过具体事物哈希值查询单个事物
	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")

	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(tx.Hash().Hex())
	fmt.Println(isPending)
}
