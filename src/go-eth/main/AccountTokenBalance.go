package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	token "go-learning/src/go-eth/contract/go"
	"math"
	"math/big"
	"strconv"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		logrus.Error(err)
	}
	// 合约地址
	tokenAddress := common.HexToAddress("0x60fFc2acc5205e9f538bF9fbFD84fC0c079375Af")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		logrus.Error(err)
	}
	address := common.HexToAddress("0x649A2B768705244C0983Ba7Dce8B782576494e68")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		logrus.Error(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		logrus.Error(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		logrus.Error(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		logrus.Error(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	atoi, err := strconv.Atoi(decimals)
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(atoi)))

	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
