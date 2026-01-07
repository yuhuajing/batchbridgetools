package utils

import (
	"fmt"
	"math/big"
)

func StringToBigInt(value string) (*big.Int, error) {
	num := new(big.Int)

	// 第2个参数表示进制：10 表示十进制
	amountInt, ok := num.SetString(value, 10)

	if !ok {
		return nil, fmt.Errorf("convert string %s to bigint", value)
	}
	return amountInt, nil
}

func HexStringToBigInt(hexStr string) *big.Int {
	// 使用 SetString 解析十六进制字符串，base 设为 0 表示自动判断（支持 0x 前缀）
	n := new(big.Int)
	n.SetString(hexStr, 0)

	return n
}
