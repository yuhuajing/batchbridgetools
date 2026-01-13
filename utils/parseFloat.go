package utils

import (
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

func ParseFloat(decimal float64, amount string) string {
	// 解析输入金额
	f := new(big.Float)
	f.SetString(amount)

	result := new(big.Float).Mul(f, big.NewFloat(decimal))

	// 转换为整数并转为字符串（防止科学计数法）
	intResult := new(big.Int)
	result.Int(intResult) // 直接截取整数部分
	return intResult.String()
}

func ParseApprovalFloat() *big.Int {
	// 解析输入金额
	f := new(big.Float)
	f.SetString("10")

	result := new(big.Float).Mul(f, big.NewFloat(params.Ether))

	// 转换为整数并转为字符串（防止科学计数法）
	intResult := new(big.Int)
	result.Int(intResult) // 直接截取整数部分
	return intResult
}
