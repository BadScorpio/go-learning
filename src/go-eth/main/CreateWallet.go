package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

func main() {
	// 获取一个ecdsa加密的私钥
	privateKey, err := crypto.GenerateKey()
	fmt.Println("privateKey=============>", privateKey)
	if err != nil {
		logrus.Error(err)
	}
	// 私钥转换为字节
	privateBytes := crypto.FromECDSA(privateKey)
	fmt.Println("privateBytes=============>", privateBytes)
	fmt.Println("privateBytesEncode=============>", hexutil.Encode(privateBytes))
	fmt.Println("privateBytesEncode[2:]=============>", hexutil.Encode(privateBytes)[2:])

	// 获取私钥对应的公钥
	publicKey := privateKey.Public()
	// 类型断言，转换成具体类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logrus.Error("cannot assert type:publicKey is not of type *ecdsa.PublicKey")
	}
	// 公钥转换成字节
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicKeyBytes===============>", publicKeyBytes)
	fmt.Println("publicKeyBytesEncode===============>", hexutil.Encode(publicKeyBytes))
	fmt.Println("publicKeyBytes[4:]===============>", hexutil.Encode(publicKeyBytes)[4:])
	// 通过公钥生成地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address==============>", address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("hash[12:]==================>", hexutil.Encode(hash.Sum(nil)[12:]))
}
