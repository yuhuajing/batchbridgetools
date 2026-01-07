package utils

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

func ParsePk(key string) (*ecdsa.PrivateKey, common.Address, error) {
	var fromAddress common.Address
	key = strings.Trim(key, " ")
	// 使用 strings.TrimPrefix 去除可能存在的 "0x" 前缀
	key = strings.TrimPrefix(key, "0x")

	pk, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, fromAddress, err
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fromAddress, fmt.Errorf("私钥解析失败")
	}
	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	return pk, fromAddress, nil
}
