package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		logrus.Error(err)
	}
	// 加载私钥
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		logrus.Error(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logrus.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 获取nonce，每笔交易需要一个单独的nonce
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonceAt, err := client.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		logrus.Error(err)
	}
	// 设置需要转账的eth，单位是wei
	value := big.NewInt(1000000000000000000)
	// 设置燃气上限
	gasLimit := uint64(21000)
	// 获取平均gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logrus.Error(err)
	}
	// 设置接收eth的地址
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	// 设置一个新的事物
	tx := types.NewTransaction(nonceAt, toAddress, value, gasLimit, gasPrice, data)
	// 获取私钥签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logrus.Error(err)
	}
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		logrus.Error(err)
	}
	// 发送事物
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println("tx send: %s", signTx.Hash().Hex())

}
