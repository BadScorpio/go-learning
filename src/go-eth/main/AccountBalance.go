package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	// 调用客户端的BalanceAt方法，传递账户地址和可选取号，区号设置为nil将返回最新余额
	account := common.HexToAddress("0x60fFc2acc5205e9f538bF9fbFD84fC0c079375Af")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal("balance===========>", err)
	}
	fmt.Println(balance)

	blockNumber := big.NewInt(0)
	// 设置返回指定区块的余额
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal("balanceAt===========>", err)
	}
	fmt.Println(balanceAt)

	// 将余额转换为ETH中的_wei_
	fblance := new(big.Float)
	fblance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fblance, big.NewFloat(math.Pow10(18)))
	fmt.Println("ethValue================>", ethValue)
	// 获取待处理的余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance)
}
